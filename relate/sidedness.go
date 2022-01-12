package relate

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/homotopy"
	"github.com/intdxdt/geom"
)

//Homotopy Relate
func Homotopy(coordinates geom.Coords, contexts *ctx.ContextGeometries) bool {
	return homotopy.Homotopy(coordinates, contexts)
}
