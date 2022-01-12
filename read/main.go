package main

import (
	"encoding/json"
	"fmt"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/geom"
	geojson "github.com/paulmach/go.geojson"
	"strings"
)

func main() {
	var polyline = readInputPolylines("input.json")
	var constraints = readInputConstraints("constraints.json")
	fmt.Println(polyline)
	fmt.Println(constraints)
}

func readInputPolylines(inputJsonFile string) []geometry.Polyline {
	return parseInputLinearFeatures(readJsonFile(inputJsonFile))
}

func readInputConstraints(inputJsonFile string) geometry.GeoJSONGeometries {
	return parseConstraintFeatures(readJsonFile(inputJsonFile))
}

type JSONLineString struct {
	Id          string
	Coordinates [][]float64
	Meta        string
}

type JSONPoint struct {
	Id          string
	Coordinates []float64
	Meta        string
}

type JSONPolygon struct {
	Id          string
	Coordinates [][][]float64
	Meta        string
}

func parseInputLinearFeatures(inputs []string) []geometry.Polyline {
	var plns = make([]geometry.Polyline, 0, len(inputs))
	for idx, fjson := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)
		var objs = getLineStringObjects(idx, feat)
		for _, o := range objs {
			var pln = createPolyline(o)
			plns = append(plns, pln)
		}
	}
	return plns
}

func parseConstraintFeatures(inputs []string) geometry.GeoJSONGeometries {
	var pts = make([]geometry.Point, 0, len(inputs))
	var plns = make([]geometry.Polyline, 0, len(inputs))
	var polys = make([]geometry.Polygon, 0, len(inputs))

	for idx, fjson := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)

		var ptObjs = getPointObjects(idx, feat)
		for _, o := range ptObjs {
			pts = append(pts, createPoint(o))
		}

		var lnObjs = getLineStringObjects(idx, feat)
		for _, o := range lnObjs {
			plns = append(plns, createPolyline(o))
		}

		var plyObjs = getPolygonObjects(idx, feat)
		for _, o := range plyObjs {
			polys = append(polys, createPolygon(o))
		}
	}

	return geometry.GeoJSONGeometries{
		Points:      pts,
		LineStrings: plns,
		Polygons:    polys,
	}
}

func createPoint(jsonLine JSONPoint) geometry.Point {
	return geometry.CreatePoint(jsonLine.Id, jsonLine.Coordinates, jsonLine.Meta)
}

func createPolyline(jsonLine JSONLineString) geometry.Polyline {
	var coords = geom.AsCoordinates(jsonLine.Coordinates)
	return geometry.CreatePolyline(jsonLine.Id, coords, jsonLine.Meta)
}

func createPolygon(jsonLine JSONPolygon) geometry.Polygon {
	var coords = make([]geom.Coords, 0, len(jsonLine.Coordinates))
	for _, array := range jsonLine.Coordinates {
		coords = append(coords, geom.AsCoordinates(array))
	}
	return geometry.CreatePolygon(jsonLine.Id, coords, jsonLine.Meta)
}

func getPointObjects(index int, feat *geojson.Feature) []JSONPoint {
	var objs = make([]JSONPoint, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsPoint() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONPoint{id, feat.Geometry.Point, string(meta)})
	} else if feat.Geometry.IsMultiPoint() {
		objs = make([]JSONPoint, 0, len(feat.Geometry.MultiPoint))
		for pos, coords := range feat.Geometry.MultiPoint {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONPoint{id, coords, string(meta)})
		}
	}
	return objs
}

func getLineStringObjects(index int, feat *geojson.Feature) []JSONLineString {
	var objs = make([]JSONLineString, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsLineString() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONLineString{id, feat.Geometry.LineString, string(meta)})
	} else if feat.Geometry.IsMultiLineString() {
		objs = make([]JSONLineString, 0, len(feat.Geometry.MultiLineString))
		for pos, coords := range feat.Geometry.MultiLineString {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONLineString{id, coords, string(meta)})
		}
	}
	return objs
}

func getPolygonObjects(index int, feat *geojson.Feature) []JSONPolygon {
	var objs = make([]JSONPolygon, 0, 1)
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	if feat.Geometry.IsPolygon() {
		var id = composeId(index, getFId(feat.Properties), 0)
		objs = append(objs, JSONPolygon{id, feat.Geometry.Polygon, string(meta)})
	} else if feat.Geometry.IsMultiPolygon() {
		objs = make([]JSONPolygon, 0, len(feat.Geometry.MultiPolygon))
		for pos, coords := range feat.Geometry.MultiPolygon {
			var id = composeId(index, getFId(feat.Properties), pos)
			objs = append(objs, JSONPolygon{id, coords, string(meta)})
		}
	}
	return objs
}

func getFId(properties map[string]interface{}) string {
	var id = properties["id"]
	if id == nil {
		return "?"
	}
	return fmt.Sprintf("%v", id)
}

func composeId(index int, fid string, pos int) string {
	return fmt.Sprintf("idx:%v-fid:%v-pos:%v", index, fid, pos)
}

func readJsonFile(file string) []string {
	var data, err = fileutil.ReadAllOfFile(file)
	checkError(err)
	var tokens = strings.Split(strings.TrimSpace(data), "\n")
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	return tokens
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
