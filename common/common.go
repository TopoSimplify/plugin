package common

import (
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"sort"
)

func SortInts(iter []int) []int {
	sort.Ints(iter)
	return iter
}

//Geometry - hull geom
func Geometry(coordinates geom.Coords) geom.Geometry {
	var g geom.Geometry
	if coordinates.Len() > 2 {
		g = geom.NewPolygon(coordinates)
	} else if coordinates.Len() == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		g = coordinates.Pt(0)
	}
	return g
}

func LinearCoords(wkt string) geom.Coords {
	return geom.NewLineStringFromWKT(wkt).Coordinates
}

func CreateHulls(id *iter.Igen, indices [][]int, coords geom.Coords, instance lnr.Linegen) []node.Node {
	var poly = geom.NewLineString(coords)
	var hulls []node.Node
	for _, o := range indices {
		hulls = append(hulls, nodeFromPolyline(
			id, poly, rng.Range(o[0], o[1]), Geometry, instance,
		))
	}
	return hulls
}

//New Node
func nodeFromPolyline(id *iter.Igen, polyline *geom.LineString, rng rng.Rng,
	geomFn func(geom.Coords) geom.Geometry, instance lnr.Linegen) node.Node {
	return node.CreateNode(id, SubCoordinates(polyline, rng), rng, geomFn, instance)
}

func SubCoordinates(ln *geom.LineString, rng rng.Rng) geom.Coords {
	var coords = ln.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}
