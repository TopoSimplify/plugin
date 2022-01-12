package ctx

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("context neighbours", func() {
		g.It("should test context neighbours", func() {
			var lnGeom = geom.NewLineString(geom.Coordinates([]geom.Point{{0, 0}, {5, 5}}))
			var ctxGeom = geom.Pt(2.5, 2.5)
			var ctxG = New(ctxGeom, 0, -1)

			inters := ctxG.Intersection(lnGeom)
			g.Assert(len(inters) == 1).IsTrue()
			g.Assert(inters[0].Equals2D(&ctxGeom)).IsTrue()
			g.Assert(ctxG.BBox().IsPoint()).IsTrue()

			var box = ctxG.Bounds()
			g.Assert(box.IsPoint()).IsTrue()
			g.Assert(ctxG.String() == ctxGeom.WKT()).IsTrue()

			g.Assert(ctxG.IsSelf()).IsTrue()
			g.Assert(ctxG.AsSelf().IsSelf()).IsTrue()
			g.Assert(ctxG.AsPlanarVertex().IsPlanarVertex())
			g.Assert(ctxG.AsNonPlanarVertex().IsNonPlanarVertex())
			g.Assert(ctxG.AsPlanarSegment().IsPlanarSegment())
			g.Assert(ctxG.AsLinearSimple().IsLinearSimple())
			g.Assert(ctxG.AsNonPlanarVertex().IsNonPlanarVertex())
			g.Assert(ctxG.AsContextNeighbour().IsContextNeighbour())
		})

		g.It("should test context neighbours", func() {
			g.Timeout(1 * time.Hour)
			coords := linearCoords("LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )")
			var rngs = [][]int{{12, coords.Len() - 1}, {8, 12}, {0, 8}}
			var ctxs = NewContexts()
			g.Assert(ctxs.Len()).Equal(0)

			var gs = ctxGeoms(rngs, coords)
			ctxs.Push(gs[0])
			ctxs.Extend(gs[1:])
			g.Assert(ctxs.Len()).Equal(3)

			var list = ctxs.DataView()
			g.Assert(list[0].Geometry()).Equal(list[0].Geom)
			g.Assert([]int{list[0].I, list[0].J}).Equal(rngs[0])
			g.Assert([]int{list[1].I, list[1].J}).Equal(rngs[1])
			g.Assert([]int{list[2].I, list[2].J}).Equal(rngs[2])
			ctxs.Sort()
			g.Assert([]int{list[0].I, list[0].J}).Equal(rngs[2])
			g.Assert([]int{list[1].I, list[1].J}).Equal(rngs[1])
			g.Assert([]int{list[2].I, list[2].J}).Equal(rngs[0])

			var objects []interface{}
			for _, h := range ctxs.list {
				objects = append(objects, h)
			}
			var contexts = NewContextsFromObjects(objects)
			g.Assert(contexts.Len()).Equal(len(objects))
			var ctxts = contexts.DataView()[0].AsContextGeometries(
				contexts.DataView()[1:]...
			)
			g.Assert(ctxts.Len()).Equal(contexts.Len())

			ctxts.Extend(contexts.DataView()[1:])
			g.Assert(ctxts.Len() == contexts.Len()).IsFalse()
			ctxts.SetData(contexts.DataView())
			g.Assert(ctxts.Len() == contexts.Len()).IsTrue()
		})
	})
}
