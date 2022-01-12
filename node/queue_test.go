package node

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("test nodequeue", func() {
		g.It("should test node queue", func() {
			g.Timeout(1 * time.Hour)
			// checks if score is valid at threshold of constrained dp
			var coords = linearCoords("LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )")
			var hulls = createHulls([][]int{{0, 2}, {2, 6}, {6, 8}, {8, 10}, {10, 12}, {12, coords.Len() - 1}}, coords)
			var queue = NewQueue()
			queue.AppendLeft(hulls[1])
			queue.AppendLeft(hulls[0])
			queue.Append(hulls[2])
			queue.Append(hulls[3])
			queue.Append(hulls[4])
			queue.Append(hulls[5])
			g.Assert(queue.Size()).Equal(len(hulls))
			g.Assert(queue.First().Range.AsArray()).Equal([2]int{0, 2})
			g.Assert(queue.Last().Range.AsArray()).Equal([2]int{12, coords.Len() - 1})
			g.Assert(queue.PopLeft().Range.AsArray()).Equal([2]int{0, 2})

			g.Assert(queue.First().Range.AsArray()).Equal([2]int{2, 6})
			g.Assert(queue.PopLeft().Range.AsArray()).Equal([2]int{2, 6})
			g.Assert(queue.Pop().Range.AsArray()).Equal([2]int{12, coords.Len() - 1})
			g.Assert(queue.Last().Range.AsArray()).Equal([2]int{10, 12})
			g.Assert(queue.Clear().Size()).Equal(0)
		})
	})
}
