package constrain

import (
	"github.com/TopoSimplify/plugin/deform"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
)

//Constrain for self-intersection as a result of simplification
//returns boolean : is hull collapsible
func ByFeatureClassIntersection(
	options *opts.Opts,
	hull *node.Node,
	db *hdb.Hdb,
	selections *[]*node.Node,
) bool {
	var bln = true
	var hulls = deform.SelectFeatureClass(options, db, hull)
	for _, h := range hulls {
		if bln && (h == hull) {
			bln = false
		}
		*selections = append(*selections, h)
	}
	return bln
}
