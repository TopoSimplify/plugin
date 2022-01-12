package decompose

import (
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

//Douglas-Peucker decomposition at a given threshold
func DouglasPeucker(
	id *iter.Igen,
	pln pln.Polyline,
	decomp offset.EpsilonDecomposition,
	geomFn func(geom.Coords) geom.Geometry,
	instance lnr.Linegen,
) []node.Node {
	var k, n int
	var val float64
	var coordinates geom.Coords
	var hque []node.Node

	if pln.LineString == nil {
		return hque
	}

	var r = pln.Range()
	var stack = make([]rng.Rng, 0, (r.J-r.I)+1)
	stack = append(stack, r)

	for !(len(stack) == 0) {
		n = len(stack) - 1
		r = stack[n]
		stack = stack[:n]

		coordinates = pln.SubCoordinates(r)
		k, val = decomp.ScoreFn(coordinates)
		k = r.I + k //offset

		if decomp.Relation(val) {
			hque = append(hque, node.CreateNode(id, coordinates, r, geomFn, instance))
		} else {
			stack = append(stack, rng.Range(k, r.J)) // right
			stack = append(stack, rng.Range(r.I, k)) // left
		}
	}
	return hque
}
