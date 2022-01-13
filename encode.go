package main

import (
	"fmt"
	"github.com/TopoSimplify/plugin/geometry"
	"regexp"
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
	fmt.Println(dictionary)
	return []string{}

}
