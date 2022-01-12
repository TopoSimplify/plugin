package lnr

import (
	"github.com/TopoSimplify/plugin/pln"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"testing"
	"time"
)

func newPolyline(wkt string) pln.Polyline {
	return pln.CreatePolyline(geom.NewLineStringFromWKT(wkt).Coordinates)
}

func TestToSelfIntersects(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("planar self intersection", func() {
		g.It("should test constrain to self intersects", func() {
			g.Timeout(1 * time.Hour)
			var ln = newPolyline("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			var inters = SelfIntersection(ln, true, true)
			g.Assert(inters.Len()).Equal(2)

			ln = newPolyline("LINESTRING ( 1000 600, 1100 600, 1100 500, 1000 500, 1000 400, 1100 400, 1100 500, 1200 500, 1200 400, 1300 400, 1300 500, 1200 500, 1200 600, 1100 600 )")
			inters = SelfIntersection(ln, true, true)
			g.Assert(inters.Len()).Equal(3)

			ln = newPolyline("LINESTRING ( 1100 100, 1300 300, 1400 200, 1400 100, 900 100, 900 0, 1100 0, 1100 100, 1000 300, 900 200, 1100 100, 1300 0, 1200 -100, 1100 -100, 1100 -200, 1300 -200, 1300 0 )")
			inters = SelfIntersection(ln, true, true)
			g.Assert(inters.Len()).Equal(7)

			ln = newPolyline("LINESTRING ( 1514.8265008716717 -573.3122984309911, 1739.3156848963422 -169.23176718658382, 1837.1395275571917 -226.77794673397926, 1808.7887862547589 -332.42242749048233, 1900.578967859 -400.895763554, 1767.3316987768692 -426.84574468170746, 1679.1814914770653 -277.4733153412824, 1591.5634006896519 -435.1858787586265, 1626.1503907141578 -525.5590212885274, 1767.3316987768692 -426.84574468170746, 1778.380225764687 -582.4209533259952 )")
			inters = SelfIntersection(ln, true, true)
			g.Assert(inters.Len()).Equal(3)
		})
	})
}
