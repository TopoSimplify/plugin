package constrain

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/franela/goblin"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

func TestByFeatureClassIntersection(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen()
	g.Describe("constrain by mindist relation", func() {
		g.It("should test constrain by context geometry", func() {
			g.Timeout(1 * time.Hour)

			var options = &opts.Opts{MinDist: 10}
			var coords = common.LinearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			var inst = dp.New(id.Next(), coords, options, offset.MaxOffset)
			var hulls = common.CreateHulls(id,
				[][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, coords.Len() - 1}}, coords, inst)

			var db = hdb.NewHdb().Load(hulls)
			var sels = []*node.Node{}

			coords = common.LinearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst)
			db.Load(hulls)

			var q1 = hulls[0]
			coords = common.LinearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst)

			db.Load(hulls)

			var q2 = hulls[0]

			g.Assert(ByFeatureClassIntersection(options, &q1, db, &sels)).IsFalse()
			g.Assert(ByFeatureClassIntersection(options, &q2, db, &sels)).IsTrue()

		})
	})
}
