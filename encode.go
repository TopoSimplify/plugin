package main

import (
	"encoding/json"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"regexp"
	"sort"
	"strconv"
)

var numRegx = regexp.MustCompile("[0-9]+")

type PolyGeom struct {
	index int
	pos   int
	geom  *geometry.Polyline
}

func indexPos(id string) (int, int) {
	var vals = numRegx.FindAllString(id, -1)
	var indx, _ = strconv.Atoi(vals[0])
	var pos, _ = strconv.Atoi(vals[1])
	return indx, pos
}

func groupById(plns []*geometry.Polyline) map[int][]PolyGeom {
	var dictionary = make(map[int][]PolyGeom, 0)
	for _, pln := range plns {
		var indx, pos = indexPos(pln.Id)
		dictionary[indx] = append(dictionary[indx], PolyGeom{indx, pos, pln})
	}
	return dictionary
}

func groupPlnsAsGeoJSONS(plns []*geometry.Polyline) []string {
	var dictionary = groupById(plns)
	var indxs = indices(dictionary)
	var jsons = make([]string, 0, len(plns))
	for idx := range indxs {
		var plns = dictionary[idx]
		sort.SliceStable(plns, func(i, j int) bool {
			return plns[i].pos < plns[j].pos
		})
		var s = createGeoJSONString(plns)
		jsons = append(jsons, s)
	}
	return jsons
}

func isCoords3D(coords []geom.Point) bool {
	var bln = false
	for _, pt := range coords {
		bln = !math.FloatEqual(pt[2], 0)
		if bln {
			break
		}
	}
	return bln
}

func extractCoords(coords []geom.Point) [][]float64 {
	var is3D = isCoords3D(coords)
	var coordinates = make([][]float64, 0, len(coords))
	for _, pt := range coords {
		var c = append([]float64{}, pt[:]...)
		if !is3D {
			c = c[:2]
		}
		coordinates = append(coordinates, c)
	}
	return coordinates
}

func getMeta(g PolyGeom) map[string]interface{} {
	var meta = g.geom.Meta
	var properties map[string]interface{}
	if meta == "null" {
		properties = make(map[string]interface{}, 0)
	} else {
		var err = json.Unmarshal([]byte(meta), &properties)
		if err != nil {
			panic(err)
		}
	}
	return properties
}

func createGeoJSONString(featLns []PolyGeom) string {
	if len(featLns) == 1 {
		var gpoly = featLns[0]
		var simple = gpoly.geom.Simple.Points()
		var coords = extractCoords(simple)
		var meta = getMeta(gpoly)

		var feature = geometry.LineStringFeature{
			Type:       "Feature",
			Properties: meta,
			Geometry: &geometry.LinearGeometry{
				Type:        "LineString",
				Coordinates: coords,
			},
		}
		var dat, _ = json.Marshal(feature)
		return string(dat)
	}
	var meta = getMeta(featLns[0])
	var coordinates = make([][][]float64, 0, len(featLns))
	for _, gpoly := range featLns {
		var simple = gpoly.geom.Simple.Points()
		var coords = extractCoords(simple)
		coordinates = append(coordinates, coords)
	}

	var feature = geometry.LineStringFeature{
		Type:       "Feature",
		Properties: meta,
		Geometry: &geometry.MultiLinearGeometry{
			Type:        "MultiLineString",
			Coordinates: coordinates,
		},
	}
	var dat, _ = json.Marshal(feature)
	return string(dat)

}

func indices(dict map[int][]PolyGeom) []int {
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
