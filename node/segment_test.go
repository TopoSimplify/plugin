package node

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"testing"
	"time"
)

func TestHullSeg(t *testing.T) {
	var g = goblin.Goblin(t)

	var createHulls = func(ranges [][]int, coords geom.Coords) []Node {
		var polyline = pln.CreatePolyline(coords)
		var hulls = make([]Node, 0)
		for _, r := range ranges {
			var i, j = r[0], r[len(r)-1]
			h := nodeFromPolyline(&polyline, rng.Range(i, j), hullGeom)
			hulls = append(hulls, h)
		}
		return hulls
	}

	g.Describe("node decomposition", func() {
		g.It("should test decomposition of a line as nodes", func() {
			g.Timeout(1 * time.Hour)
			var wkt = "LINESTRING ( 670 550, 680 580, 750 590, 760 630, 830 640, 870 630, 890 610, 920 580, 910 540, 890 500, 900 460, 870 420, 860 390, 810 360, 770 400, 760 420, 800 440, 810 470, 850 500, 820 560, 780 570, 760 530, 720 530, 707.3112236920351 500.3928552814154, 650 450 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates
			var ranges = [][]int{{0, 12}, {12, 18}, {18, coords.Len() - 1}}
			var hulls = createHulls(ranges, coords)

			for i, r := range ranges {
				var s = hulls[i].Segment()
				var a, b = coords.Pt(r[0])[:2], coords.Pt(r[1])[:2]
				g.Assert(r).Equal(s.Coords.Idxs)
				g.Assert(s.A()[:2]).Equal(a)
				g.Assert(s.B()[:2]).Equal(b)

				g.Assert([]geom.Point{*s.A(), *s.B()}).Equal(hulls[i].SegmentAsPolyline().Coordinates.Points())
				var cs = hulls[i].Polyline.Coordinates.Points()
				cs = append(cs, cs[0])
				g.Assert(hulls[i].Coordinates()).Equal(hulls[i].Polyline.Coordinates)
				g.Assert(hulls[i].String()).Equal(hulls[i].Geom.WKT())
				g.Assert(hulls[i].Geom).Equal(hulls[i].Geom)
				g.Assert(hulls[i].Geom.BBox()).Equal(hulls[i].BBox())
			}
		})
	})
}
