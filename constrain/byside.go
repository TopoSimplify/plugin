package constrain

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/relate"
)

func BySideRelation(hull *node.Node, cgs *ctx.ContextGeometries) bool {
	return relate.IsDirRelateValid(hull, cgs)
}
