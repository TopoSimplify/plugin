package deform

import (
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/knn"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
)

func optimizeNeighbours(hull *node.Node, neighbours []*node.Node) []*node.Node {
	if hull.Instance.State().IsDirty() {
		return neighbours
	}
	var n = len(neighbours)

	for _, o := range neighbours[:n] {
		if !isSame(hull.Instance, o.Instance) {
			neighbours = append(neighbours, o)
		}
	}
	return neighbours[n:]
}

//find context_geom deformable hulls
func SelectFeatureClass(options *opts.Opts, hullDB *hdb.Hdb, hull *node.Node) []*node.Node {
	var n int
	var h *node.Node
	var inters, contig bool
	var dict = make(map[[2]int]*node.Node, 0)
	var ctxHulls = knn.NodeNeighbours(hullDB, hull, knn.EpsilonDist)

	ctxHulls = optimizeNeighbours(hull, ctxHulls)

	// for each item in the context_geom list
	for i := range ctxHulls {
		n = 0
		h = ctxHulls[i]

		var sameFeature = isSame(hull.Instance, h.Instance)
		// find which item to deform against current hull
		if sameFeature { // check for contiguity
			inters, contig, n = node.IsContiguous(hull, h)
		} else {
			// contiguity is by default false for different features
			contig = false
			inters = hull.Geom.Intersects(h.Geom)
			if inters {
				var pts = hull.Geom.Intersection(h.Geom)
				inters = len(pts) > 0
				n = len(pts)
			}
		}

		if !inters { // disjoint : nothing to do, continue
			continue
		}

		var sa, sb *node.Node
		if contig && n > 1 { // contiguity with overlap greater than a vertex
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
