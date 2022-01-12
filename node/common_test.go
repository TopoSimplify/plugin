package node

import (
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

var idgen = iter.NewIgen(0)

// hull geom
func hullGeom(coords geom.Coords) geom.Geometry {
	var g geom.Geometry

	if coords.Len() > 2 {
		g = geom.NewPolygon(coords)
	} else if coords.Len() == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords.Pt(0)
	}
	return g
}

func linearCoords(wkt string) geom.Coords {
	return geom.NewLineStringFromWKT(wkt).Coordinates
}

func createHulls(indxs [][]int, coords geom.Coords) []Node {
	var poly = pln.CreatePolyline(coords)
	var hulls = make([]Node, 0, len(indxs))
	for _, o := range indxs {
		r := rng.Range(o[0], o[1])
		hulls = append(hulls, CreateNode(idgen, poly.SubCoordinates(r), r, hullGeom, nil))
	}
	return hulls
}

// CreateNode Node
func nodeFromPolyline(polyline *pln.Polyline, rng rng.Rng, geomFn func(geom.Coords) geom.Geometry) Node {
	return CreateNode(idgen, polyline.SubCoordinates(rng), rng, geomFn, nil)
}
