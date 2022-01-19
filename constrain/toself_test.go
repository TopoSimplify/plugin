package constrain

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/franela/goblin"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

func TestToSelfIntersects(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen()
	g.Describe("constrain", func() {
		g.It("should test constrain to self intersects - 1", func() {
			g.Timeout(1 * time.Hour)
			var coords = common.LinearCoords("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			//var cong = geom.NewPolygonFromWKT("POLYGON (( 780 560, 780 580, 800 580, 800 560, 780 560 ))")
			var polyline = geometry.CreatePolyline("0", coords, "")
			options := &opts.Opts{
				Threshold:              1.0,
				MinDist:                1.0,
				NonPlanarDisplacement:  1.0,
				PlanarSelf:             true,
				NonPlanarSelf:          true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = common.CreateHulls(id,
				[][]int{{0, 5}, {5, 9}, {9, 11}}, coords, nil)

			g.Assert(len(nodes)).Equal(3)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{}

			options.PlanarSelf = false
			options.NonPlanarSelf = false
			var que, bln, set = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(3)
			g.Assert(len(set)).Equal(0)

			constVerts = []int{10}
			options.PlanarSelf = true
			options.NonPlanarSelf = true
			que, bln, set = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)

			g.Assert(bln).IsTrue()
			g.Assert(set).Equal([]int{0, 1, 3, 7, 9, 10})
		})

		g.It("should test constrain to self intersects - 2", func() {
			g.Timeout(1 * time.Hour)
			var coords = common.LinearCoords("LINESTRING ( 780 480, 750 470, 760 500, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			//var cong = geom.NewPolygonFromWKT("POLYGON (( 780 560, 780 580, 800 580, 800 560, 780 560 ))")
			var polyline = geometry.CreatePolyline("0", coords, "")
			var options = &opts.Opts{
				Threshold:              1.0,
				MinDist:                1.0,
				NonPlanarDisplacement:  1.0,
				PlanarSelf:             true,
				NonPlanarSelf:          true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = common.CreateHulls(id,
				[][]int{{0, 5}, {5, 9}, {9, 11}}, coords, nil)

			g.Assert(len(nodes)).Equal(3)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{10}

			options.PlanarSelf = false
			options.NonPlanarSelf = false
			var que, bln, set = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(3)
			g.Assert(len(set)).Equal(0)

			constVerts = []int{10}
			options.PlanarSelf = true
			options.NonPlanarSelf = true
			que, bln, set = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)

			g.Assert(bln).IsTrue()
			nodes = []node.Node{}

			g.Assert(len(que)).Equal(6)
			g.Assert(set).Equal([]int{3, 7, 10})
		})

		g.It("should test constrain to self intersects - 3", func() {
			g.Timeout(1 * time.Hour)
			var coords = common.LinearCoords("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			//var cong = geom.NewPolygonFromWKT("POLYGON (( 780 560, 780 580, 800 580, 800 560, 780 560 ))")
			//pln = geometry.CreatePolyline("0", coords, "")
			var polyline = geometry.CreatePolyline("0", coords, "")
			var options = &opts.Opts{
				Threshold:              300.0,
				MinDist:                300.0,
				NonPlanarDisplacement:  300.0,
				PlanarSelf:             true,
				NonPlanarSelf:          true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = common.CreateHulls(id,
				[][]int{{0, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 9}, {9, 11}}, coords, nil)
			g.Assert(len(nodes)).Equal(7)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{10}

			options.PlanarSelf = false
			options.NonPlanarSelf = false
			var que, bln, _ = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(len(nodes))

			options.PlanarSelf = true
			options.NonPlanarSelf = true
			que, bln, _ = ToSelfIntersects(id,
				queue, polyline, options, constVerts,
			)
			g.Assert(len(que)).Equal(9)
		})

		g.It("should test constrain to self intersects - 4", func() {
			g.Timeout(1 * time.Hour)
			var coords = common.LinearCoords("LINESTRING ( 300 0, 300 400, 600 600, 600 1000, 900 1000, 900 700, 1300 700, 1400 400, 1600 200, 1300 0, 800 100, 300 0 )")
			var polyline = geometry.CreatePolyline("0", coords, "")
			var options = &opts.Opts{
				Threshold:              300.0,
				MinDist:                300.0,
				NonPlanarDisplacement:  300.0,
				PlanarSelf:             true,
				NonPlanarSelf:          true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var nodes = common.CreateHulls(id, [][]int{{0, 11}}, coords, nil)
			g.Assert(len(nodes)).Equal(1)
			var queue = nodes[:len(nodes):len(nodes)]
			var constVerts = []int{0, 11}

			options.PlanarSelf = true
			options.NonPlanarSelf = true
			var que, bln, _ = ToSelfIntersects(id, queue, polyline, options, constVerts)
			g.Assert(bln).IsTrue()
			g.Assert(len(que)).Equal(1)
		})
	})
}
