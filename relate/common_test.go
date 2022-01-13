package relate

import (
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/intdxdt/geom"
)

func polyln(wkt string) *geometry.Polyline {
	return geometry.CreatePolyline("0", geom.NewLineStringFromWKT(wkt).Coordinates, "")
}
