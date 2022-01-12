package constrain

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/relate"
)

func ByMinDistRelation(options *opts.Opts, hull *node.Node, cg *ctx.ContextGeometries) bool {
	return relate.IsDistRelateValid(options, hull, cg)
}
