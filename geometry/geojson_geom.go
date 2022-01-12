package geometry

import (
	"encoding/json"
	geojson "github.com/paulmach/go.geojson"
)

func pointsFromFeature(index int, feat *geojson.Feature) []JSONPoint {
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	return geoJsonPoints(index, feat.Geometry, string(meta))
}

func geoJsonPoints(index int, g *geojson.Geometry, meta string) []JSONPoint {
	var objs = make([]JSONPoint, 0, 1)
	if g.IsPoint() {
		var id = composeId(index, 0)
		objs = append(objs, JSONPoint{id, g.Point, meta})
	} else if g.IsMultiPoint() {
		objs = make([]JSONPoint, 0, len(g.MultiPoint))
		for pos, coords := range g.MultiPoint {
			var id = composeId(index, pos)
			objs = append(objs, JSONPoint{id, coords, string(meta)})
		}
	}
	return objs
}

func lineStringFromFeature(index int, feat *geojson.Feature) []JSONLineString {
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	return geoJsonLineString(index, feat.Geometry, string(meta))
}

func geoJsonLineString(index int, g *geojson.Geometry, meta string) []JSONLineString {
	var objs = make([]JSONLineString, 0, 1)
	if g.IsLineString() {
		objs = append(objs, JSONLineString{composeId(index, 0), g.LineString, meta})
	} else if g.IsMultiLineString() {
		objs = make([]JSONLineString, 0, len(g.MultiLineString))
		for pos, coords := range g.MultiLineString {
			objs = append(objs, JSONLineString{composeId(index, pos), coords, string(meta)})
		}
	}
	return objs
}

func polygonFromFeature(index int, feat *geojson.Feature) []JSONPolygon {
	var meta, err = json.Marshal(feat.Properties)
	checkError(err)
	return geoJsonPolygons(index, feat.Geometry, string(meta))

}

func geoJsonPolygons(index int, g *geojson.Geometry, meta string) []JSONPolygon {
	var objs = make([]JSONPolygon, 0, 1)
	if g.IsPolygon() {
		objs = append(objs, JSONPolygon{composeId(index, 0), g.Polygon, meta})
	} else if g.IsMultiPolygon() {
		objs = make([]JSONPolygon, 0, len(g.MultiPolygon))
		for pos, coords := range g.MultiPolygon {
			objs = append(objs, JSONPolygon{composeId(index, pos), coords, meta})
		}
	}
	return objs
}
