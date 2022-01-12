package deform

import (
	"github.com/TopoSimplify/plugin/node"
)

//select contiguous candidates
func contiguousCandidates(a, b *node.Node) (*node.Node, *node.Node) {
	//var selection = make([]*node.Node, 0)
	// compute sidedness relation between contiguous hulls to avoid hull flip

	// all hulls that are simple should be collapsible
	// if not collapsible -- add to selection for deformation
	// to reach colapsibility
	var sa, sb *node.Node
	//& the present should not affect the future
	if !a.Collapsible(b) {
		sa = a
	}

	// future should not affect the present
	if !b.Collapsible(a) {
		sb = b
	}
	return sa, sb
}
