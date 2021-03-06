package split

import (
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

//split hull at vertex with
//maximum_offset offset -- k
func AtScoreSelection(id *iter.Igen, hull *node.Node, scoreFn lnr.ScoreFn, gfn func(geom.Coords) geom.Geometry) (node.Node, node.Node) {
	var coordinates = hull.Coordinates()
	var rg = hull.Range
	var i, j = rg.I, rg.J
	var k, _ = scoreFn(coordinates)
	var idx = rg.I + k
	// ---------------------------------------------------------------
	// i..[ha]..k..[hb]..j
	var ha = node.CreateNode(id, coordinates.Slice(0, k+1), rng.Range(i, idx), gfn, hull.Instance)
	var hb = node.CreateNode(id, coordinates.Slice(k, coordinates.Len()), rng.Range(idx, j), gfn, hull.Instance)
	// ---------------------------------------------------------------
	return ha, hb
}

//split hull at indices (index, index, ...)
func AtIndex(id *iter.Igen, hull *node.Node, indices []int, gfn func(geom.Coords) geom.Geometry) []node.Node {
	//formatter:off
	var coordinates = hull.Coordinates()
	var ranges = hull.Range.Split(indices)
	var subHulls = make([]node.Node, 0, len(ranges))
	var I = hull.Range.I
	for _, r := range ranges {
		subHulls = append(subHulls, node.CreateNode(id, coordinates.Slice(r.I-I, r.J-I+1), r, gfn, hull.Instance))
	}
	return subHulls
}
