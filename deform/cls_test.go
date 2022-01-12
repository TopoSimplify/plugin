package deform

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/franela/goblin"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

func TestSelectFeatureClass(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen()
	g.Describe("constrain by select fc", func() {
		g.It("should test fc selection", func() {
			g.Timeout(1 * time.Hour)
			var options = &opts.Opts{MinDist: 10}
			var coords = common.LinearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			var inst = dp.New(id.Next(), coords, options, offset.MaxOffset)
			var hulls = common.CreateHulls(id,
				[][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, coords.Len() - 1}}, coords, inst)

			var db = hdb.NewHdb().Load(hulls)

			coords = common.LinearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			hulls = common.CreateHulls(id,
				[][]int{{0, coords.Len() - 1}}, coords, inst)

			db.Load(hulls)

			var q1 = hulls[0]
			coords = common.LinearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst)

			db.Load(hulls)

			var q2 = hulls[0]

			var selections = SelectFeatureClass(options, db, &q1)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, &q2)
			g.Assert(len(selections)).Equal(0)

		})
		g.It("should test fc selection different features", func() {
			g.Timeout(1 * time.Hour)

			var options = &opts.Opts{MinDist: 10}
			var coords = common.LinearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")

			var inst0 = dp.New(id.Next(), coords, options, offset.MaxOffset)
			var hulls = common.CreateHulls(id,
				[][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, coords.Len() - 1}}, coords, inst0)
			//DebugPrintNodes(hulls)
			var q0 = &hulls[2]

			var db = hdb.NewHdb().Load(hulls)

			coords = common.LinearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			var inst1 = dp.New(id.Next(), coords, options, offset.MaxOffset)
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst1)

			db.Load(hulls)

			var q1 = &hulls[0]
			coords = common.LinearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst1)

			db.Load(hulls)

			var q2 = &hulls[0]

			coords = common.LinearCoords("LINESTRING ( 750.5719204078739 667.8504262852285, 731.1163192182406 669.4717263843646, 730.3819045734933 682.6968257108445, 734.5615819289048 700, 740.8441198130572 706.1536411273189, 756.0438082424582 709.5989038379831, 752.801208044186 700, 757.5947734471756 691.9692691038592 )")
			hulls = common.CreateHulls(id, [][]int{{0, coords.Len() - 1}}, coords, inst1)

			db.Load(hulls)

			var q3 = &hulls[0]

			var selections = SelectFeatureClass(options, db, q0)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, q1)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, q2)
			g.Assert(len(selections)).Equal(0)

			selections = SelectFeatureClass(options, db, q3)
			g.Assert(len(selections)).Equal(0)

			//fmt.Println(q0)
			//fmt.Println(q1)
			//fmt.Println(q2)
		})

		//g.It("should test fc selection different features", func() {
		//	g.Timeout(1 * time.Hour)
		//
		//	var options = &opts.Opts{MinDist: 10}
		//	var coords = common.LinearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
		//	var hulls = common.CreateHulls([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, coords.Len() - 1}}, coords)
		//	//DebugPrintNodes(hulls)
		//	var q0 = hulls[2]
		//
		//	var inst0 = dp.New(coords, options, offset.MaxOffset)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst0
		//	}
		//
		//	var db = hullsDB(hulls)
		//
		//	coords = common.LinearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
		//	hulls = common.CreateHulls([][]int{{0, coords.Len() - 1}}, coords)
		//
		//	var inst1 = dp.New(coords, options, offset.MaxOffset)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst1
		//		db.Insert(h)
		//	}
		//
		//	var q1 = hulls[0]
		//	coords = common.LinearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
		//	hulls = common.CreateHulls([][]int{{0, coords.Len() - 1}}, coords)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst1
		//		db.Insert(h)
		//	}
		//	var q2 = hulls[0]
		//
		//	var selections = SelectFeatureClass(options, db, q0)
		//	g.Assert(len(selections)).Equal(1)
		//
		//	selections = SelectFeatureClass(options, db, q1)
		//	g.Assert(len(selections)).Equal(1)
		//
		//	selections = SelectFeatureClass(options, db, q2)
		//	g.Assert(len(selections)).Equal(0)
		//
		//	//fmt.Println(q0)
		//	//fmt.Println(q1)
		//	//fmt.Println(q2)
		//
		//})
	})
}
