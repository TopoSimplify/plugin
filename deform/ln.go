package deform

import (
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/knn"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
)

//find context deformation list
func Select(options *opts.Opts, hullDB *hdb.Hdb, hull *node.Node) []*node.Node {
	var n int
	var h *node.Node
	var inters, contig bool
	var dict = make(map[[2]int]*node.Node, 0)
	var ctxHulls = knn.NodeNeighbours(hullDB, hull, knn.EpsilonDist)

	// for each item in the context list
	for i := range ctxHulls {
		// find which item to deform against current hull
		h = ctxHulls[i]

		inters, contig, n = node.IsContiguous(hull, h)

		if !inters {
			continue
		}

		var sa, sb *node.Node
		if contig && n > 1 { //contiguity with overlap greater than a vertex
			sa, sb = contiguousCandidates(hull, h)
		} else if !contig {
			sa, sb = nonContiguousCandidates(options, hull, h)
		}

		// add candidate deformation hulls to selection list
		if sa != nil {
			dict[sa.Range.AsArray()] = sa
		}
		if sb != nil {
			dict[sb.Range.AsArray()] = sb
		}
	}

	var items = make([]*node.Node, 0, len(dict))
	for _, v := range dict {
		items = append(items, v)
	}
	return items
}
