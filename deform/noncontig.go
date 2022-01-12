package deform

import (
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
)

//select non-contiguous candidates
func nonContiguousCandidates(options *opts.Opts, a, b *node.Node) (*node.Node, *node.Node) {
	var aseg = a.Segment()
	var bseg = b.Segment()

	var aln = a.Polyline.Geometry()
	var bln = b.Polyline.Geometry()

	var asegIntersBseg = aseg.Intersects(bseg)
	var asegIntersBln = aseg.Intersects(bln)
	var bsegIntersAln = bseg.Intersects(aln)
	var alnIntersBln = aln.Intersects(bln)
	var sa, sb *node.Node

	if asegIntersBseg && asegIntersBln && (!alnIntersBln) {
		sa = a
	} else if asegIntersBseg && bsegIntersAln && (!alnIntersBln) {
		sb = b
	} else if alnIntersBln {
		// find out whether is a shared vertex or overlap
		// is aseg inter bset  --- dist --- aln inter bln > relax dist
		var ptLns = aln.Intersection(bln)
		var atSeg = aseg.Intersection(bseg)

		// if segs are disjoint but lines intersect, deform a&b
		if len(atSeg) == 0 && len(ptLns) > 0 {
			sa, sb = a, b
		} else {
		outer:
			for i := range ptLns {
				for j := range atSeg {
					delta := ptLns[i].Distance(atSeg[j])
					if delta > options.RelaxDist {
						sa, sb = a, b
						break outer
					}
				}
			}
		}
	}

	return sa, sb
}
