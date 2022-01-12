package constrain

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/intdxdt/geom"
)

func ctxGeom(wkt string) *ctx.ContextGeometry {
	return ctx.New(geom.ReadGeometry(wkt), 0, -1)
}
