package geometry

import (
	"encoding/json"
	"fmt"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/geom"
	geojson "github.com/paulmach/go.geojson"
	"strings"
)

const (
	TypePoint             string = "Point"
	TypeLineString        string = "LineString"
	TypePolygon           string = "Polygon"
	TypeMultiPoint        string = "MultiPoint"
	TypeMultiLineString   string = "MultiLineString"
	TypeMultiPolygon      string = "MultiPolygon"
	TypeFeature           string = "Feature"
	TypeFeatureCollection string = "FeatureCollection"
)

type IGeometry interface {
	Geometry() geom.Geometry
}

type JSONLineString struct {
	Id          string
	Coordinates [][]float64
	Meta        string
}

type JSONPolygon struct {
	Id          string
	Coordinates [][][]float64
	Meta        string
}

func ReadInputPolylines(inputJsonFile string) []*Polyline {
	var tokens = readJsonFile(inputJsonFile)
	return parseInputLinearFeatures(tokens)
}

func ReadInputConstraints(inputJsonFile string) []IGeometry {
	var tokens = readJsonFile(inputJsonFile)
	return parseConstraintFeatures(tokens)
}

func parseInputLinearFeatures(inputs []string) []*Polyline {
	var plns = make([]*Polyline, 0, len(inputs))

	for idx, fjson := range inputs {
		var jtype JSONType
		var err = json.Unmarshal([]byte(fjson), &jtype)
		checkError(err)

		if jtype.isGeometryType() {
			g, err := geojson.UnmarshalGeometry([]byte(fjson))
			checkError(err)
			var feat = &geojson.Feature{Geometry: g, Type: string(g.Type)}
			plns = append(plns, createPolylineGeoms(idx, feat)...)
		} else if jtype.isFeatureType() {
			feat, err := geojson.UnmarshalFeature([]byte(fjson))
			checkError(err)
			plns = append(plns, createPolylineGeoms(idx, feat)...)
		} else if jtype.isFeatureCollectionType() {
			feats, err := geojson.UnmarshalFeatureCollection([]byte(fjson))
			checkError(err)
			for _, feat := range feats.Features {
				plns = append(plns, createPolylineGeoms(idx, feat)...)
			}
		}
	}
	return plns
}

func parseConstraintFeatures(inputs []string) []IGeometry {
	var geometries = make([]IGeometry, 0, len(inputs))

	for idx, fjson := range inputs {
		var jtype JSONType
		var err = json.Unmarshal([]byte(fjson), &jtype)
		checkError(err)

		if jtype.isGeometryType() {
			g, err := geojson.UnmarshalGeometry([]byte(fjson))
			checkError(err)
			var feat = &geojson.Feature{Geometry: g, Type: string(g.Type)}
			geometries = append(geometries, createIGeom(idx, feat)...)
		} else if jtype.isFeatureType() {
			feat, err := geojson.UnmarshalFeature([]byte(fjson))
			checkError(err)
			geometries = append(geometries, createIGeom(idx, feat)...)
		} else if jtype.isFeatureCollectionType() {
			feats, err := geojson.UnmarshalFeatureCollection([]byte(fjson))
			checkError(err)
			for _, feat := range feats.Features {
				geometries = append(geometries, createIGeom(idx, feat)...)
			}
		}
	}

	return geometries
}

func createPolylineGeoms(idx int, feat *geojson.Feature) []*Polyline {
	var geometries = make([]*Polyline, 0)
	if feat.Geometry.IsLineString() || feat.Geometry.IsMultiLineString() {
		var objs = lineStringFromFeature(idx, feat)
		for _, o := range objs {
			geometries = append(geometries, createPolyline(o))
		}
	}
	return geometries
}

func createIGeom(idx int, feat *geojson.Feature) []IGeometry {
	var geometries = make([]IGeometry, 0)
	if feat.Geometry.IsPoint() || feat.Geometry.IsMultiPoint() {
		var objs = pointsFromFeature(idx, feat)
		for _, o := range objs {
			geometries = append(geometries, createPoint(o))
		}
	} else if feat.Geometry.IsLineString() || feat.Geometry.IsMultiLineString() {
		var objs = lineStringFromFeature(idx, feat)
		for _, o := range objs {
			geometries = append(geometries, createPolyline(o))
		}
	} else if feat.Geometry.IsPolygon() || feat.Geometry.IsMultiPolygon() {
		var objs = polygonFromFeature(idx, feat)
		for _, o := range objs {
			geometries = append(geometries, createPolygon(o))
		}
	}
	return geometries
}

func createPoint(jsonLine JSONPoint) Point {
	return CreatePoint(jsonLine.Id, jsonLine.Coordinates, jsonLine.Meta)
}

func createPolyline(jsonLine JSONLineString) *Polyline {
	var coords = geom.AsCoordinates(jsonLine.Coordinates)
	return CreatePolyline(jsonLine.Id, coords, jsonLine.Meta)
}

func createPolygon(jsonLine JSONPolygon) Polygon {
	var coords = make([]geom.Coords, 0, len(jsonLine.Coordinates))
	for _, array := range jsonLine.Coordinates {
		coords = append(coords, geom.AsCoordinates(array))
	}
	return CreatePolygon(jsonLine.Id, coords, jsonLine.Meta)
}

func composeId(index int, pos int) string {
	return fmt.Sprintf("idx:%v-pos:%v", index, pos)
}

func readJsonFile(file string) []string {
	var data, err = fileutil.ReadAllOfFile(file)
	checkError(err)
	var tokens = make([]string, 0)
	var toks = strings.Split(strings.TrimSpace(data), "\n")
	for _, tok := range toks {
		var v = strings.TrimSpace(tok)
		if v != "" {
			tokens = append(tokens, v)
		}
	}
	return tokens
}
