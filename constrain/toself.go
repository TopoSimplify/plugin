package constrain

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/intdxdt/iter"
	"sort"
)

//Constrain for planar self-intersection
func ToSelfIntersects(id *iter.Igen,
	nodes []node.Node, polyline *geometry.Polyline,
	options *opts.Opts, constVerts []int,
) ([]node.Node, bool, []int) {
	var atVertexSet = make(map[int]bool)
	if !options.PlanarSelf {
		return nodes, true, []int{}
	}

	var hulldb = hdb.NewHdb().Load(nodes)
	var planar, nonPlanar = options.PlanarSelf, options.NonPlanarSelf
	var selfInters = lnr.SelfIntersection(polyline, planar, nonPlanar)

	for _, inter := range selfInters.DataView() {
		var indices = inter.Meta.Planar
		if inter.IsNonPlanarVertex() {
			indices = inter.Meta.NonPlanar
		}
		for _, v := range indices {
			atVertexSet[v] = true
		}
	}

	//update  const vertices if any
	//add const vertices as self inters
	for _, i := range constVerts {
		if atVertexSet[i] { //exclude already self intersects
			continue
		}
		atVertexSet[i] = true
		var pt = polyline.Pt(i)
		var cg = ctx.New(pt, 0, -1).AsPlanarVertex()

		cg.Meta.Planar = append(cg.Meta.Planar, i)
		selfInters.Push(cg)
	}

	splitAtSelfIntersects(id, hulldb, selfInters)

	nodes = make([]node.Node, 0, len(nodes))
	for _, n := range hulldb.All() {
		nodes = append(nodes, *n)
	}

	sort.Sort(node.Nodes(nodes))

	var indices = make([]int, 0, len(atVertexSet))
	for v := range atVertexSet {
		indices = append(indices, v)
	}
	sort.Ints(indices)
	return nodes, true, indices
}
