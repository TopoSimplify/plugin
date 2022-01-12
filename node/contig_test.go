package node

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestIsContiguous(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("test contig nodes", func() {
		g.It("should test contiguous", func() {
			g.Timeout(1 * time.Hour)
			// checks if score is valid at threshold of constrained dp
			var coords = linearCoords("LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )")
			var hulls = createHulls([][]int{{0, 2}, {2, 6}, {6, 8}, {8, 10}, {10, 12}, {12, coords.Len() - 1}}, coords)

			inters, contig, n := IsContiguous(&hulls[0], &hulls[1])
			g.Assert(inters).IsTrue()
			g.Assert(contig).IsTrue()
			g.Assert(n).Equal(1)

			var _nodes = hulls[4:len(hulls):len(hulls)]
			var nodes []*Node
			for i := range _nodes {
				nodes = append(nodes, &_nodes[i])
			}
			nodes = append(nodes, &hulls[0], &hulls[1], &hulls[2], &hulls[3])

			var prv, nxt = Neighbours(&hulls[0], nodes)
			g.Assert(prv == nil).IsTrue()
			g.Assert(nxt != nil).IsTrue()
			g.Assert(nxt.Range.AsArray()).Equal([2]int{2, 6})

			prv, nxt = Neighbours(&hulls[1], nodes)
			g.Assert(prv.Range.AsArray()).Equal([2]int{0, 2})
			g.Assert(nxt.Range.AsArray()).Equal([2]int{6, 8})
		})
	})
}
