package relate

import (
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/node"
)

//direction relate
func IsDirRelateValid(hull *node.Node, ctx *ctx.ContextGeometries) bool {
	return Homotopy(hull.Coordinates(), ctx)
}
