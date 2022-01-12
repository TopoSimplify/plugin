package relate

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/homotopy"
)

//Homotopy Relate
func Homotopy(coordinates geom.Coords, contexts *ctx.ContextGeometries) bool {
	return homotopy.Homotopy(coordinates, contexts)
}
