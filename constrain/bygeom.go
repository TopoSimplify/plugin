package constrain

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/relate"
)

func ByGeometricRelation(hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsGeomRelateValid(hull, cg)
}
