package ctx

import (
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
)

func linearCoords(wkt string) geom.Coords {
	return geom.NewLineStringFromWKT(wkt).Coordinates
}

//create ctx geometries
func ctxGeoms(indxs [][]int, coords geom.Coords) []*ContextGeometry {
	var hulls []*ContextGeometry
	for _, o := range indxs {
		var r = rng.Range(o[0], o[1])
		var g = hullGeom(coords.Slice(r.I, r.J+1))
		var cg = New(g, r.I, r.J)
		hulls = append(hulls, cg)
	}
	return hulls
}

//hull geom
func hullGeom(coords geom.Coords) geom.Geometry {
	var g geom.Geometry
	var n = coords.Len()
	if n > 2 {
		g = geom.NewPolygon(coords)
	} else if n == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords.Pt(0)
	}
	return g
}
