package constdp

import (
	"github.com/TopoSimplify/plugin/constrain"
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/knn"
	"github.com/TopoSimplify/plugin/node"
)

func (self *ConstDP) ValidateMerge(hull *node.Node, hulldb *hdb.Hdb) bool {
	var bln = true
	var sideEffects []*node.Node
	// self intersection constraint
	if self.Opts.AvoidNewSelfIntersects {
		bln = constrain.BySelfIntersection(
			self.Opts, hull, hulldb, &sideEffects,
		)
	}

	if len(sideEffects) != 0 || !bln {
		return false
	}

	// context geometry constraint
	bln = self.ValidateContextRelation(hull, &sideEffects)
	return bln && (len(sideEffects) == 0)
}

//ValidateContextRelation - Constrain for context neighbours
// finds the collapsibility of hull with respect to context hull neighbours
// if hull is deformable, its added to selections
func (self *ConstDP) ValidateContextRelation(hull *node.Node, selections *[]*node.Node) bool {
	var bln = true

	if !(self.Opts.GeomRelation || self.Opts.DistRelation || self.Opts.DirRelation) {
		return bln
	}

	// find context neighbours - if valid
	var boxObjs = knn.ContextNeighbours(self.ContextDB, hull.Geom, self.Opts.MinDist)

	var neighbours = make([]*ctx.ContextGeometry, len(boxObjs))
	for i, o := range boxObjs {
		neighbours[i] = o.Geom.(*ctx.ContextGeometry)
	}

	var ctxtgeoms = (&ctx.ContextGeometries{}).SetData(neighbours)

	if bln && self.Opts.GeomRelation {
		bln = constrain.ByGeometricRelation(hull, ctxtgeoms)
	}

	if bln && self.Opts.DistRelation {
		bln = constrain.ByMinDistRelation(self.Options(), hull, ctxtgeoms)
	}

	if bln && self.Opts.DirRelation {
		bln = constrain.BySideRelation(hull, ctxtgeoms)
	}

	if !bln {
		*selections = append(*selections, hull)
	}

	return bln
}
