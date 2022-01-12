package node

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/rtree"
	"sort"
	"testing"
)

func popObj(self *[]*rtree.BoxObject) []*rtree.BoxObject {
	var nodes = *self
	if len(nodes) != 0 {
		n := len(nodes) - 1
		nodes[n] = nil
		*self = nodes[:n]
	}
	return nodes
}

func TestNodes(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("test contig nodes", func() {
		g.It("should test nodes", func() {
			// checks if score is valid at threshold of constrained dp
			var coords = linearCoords("LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )")
			var hulls = createHulls([][]int{{8, 10}, {10, 12}, {0, 2}, {2, 6}, {6, 8}, {12, coords.Len() - 1}}, coords)
			var ns  = make([]Node, 0)
			for _, n := range hulls {
				ns = append(ns, n)
			}
			sort.Sort(Nodes(ns))
			g.Assert(ns[0].Range.AsArray()).Equal([2]int{0, 2})
			sort.Sort(sort.Reverse(Nodes(ns)))
			g.Assert(ns[0].Range.AsArray()).Equal([2]int{12, coords.Len() - 1})
			var n = ns[(len(hulls) - 1)]
			g.Assert(ns).Equal(Pop(&ns))
			g.Assert(n.Range.AsArray()).Equal([2]int{0, 2})
			g.Assert(len(ns)).Equal(len(hulls) - 1)

			var ptset = Nodes(ns).AsPointSet()
			g.Assert(ptset.Values()).Equal([]interface{}{2, 6, 8, 10, 12, 15})
			Clear(&ns)

			g.Assert(len(ns)).Equal(0)

		})
	})
}
