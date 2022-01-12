package pln

import (
	"github.com/TopoSimplify/rng"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Pln", func() {
		g.It("should test Pln", func() {
			g.Timeout(1 * time.Minute)
			var wkt = "LINESTRING ( 470 480, 470 450, 490 430, 520 420, 540 440, 560 430, 580 420, 590 410, 630 400, 630 430, 640 460, 630 490, 630 520, 640 540, 660 560, 690 580, 700 600, 730 600, 750 570, 780 560, 790 550, 800 520, 830 500, 840 480, 850 460, 900 440, 920 440, 950 480, 990 480, 1000 520, 1000 570, 990 600, 1010 620, 1060 600 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates
			var poly = CreatePolyline(coords)
			var sub = poly.SubPolyline(rng.Range(3, 13))
			var n = len(coords.Points())
			g.Assert(poly.Len()).Equal(n)
			g.Assert(poly.Range()).Equal(rng.Range(0, n-1))
			g.Assert(sub.Len()).Equal(13 - 3 + 1)
		})
	})
}
