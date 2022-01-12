package relate

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/node"
)

//IsDirRelateValid - direction relate
func IsDirRelateValid(hull *node.Node, ctx *ctx.ContextGeometries) bool {
	return Homotopy(hull.Coordinates(), ctx)
}
