package decompose

import (
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

func TestDecompose(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen(0)
	g.Describe("hull decomposition", func() {
		g.It("should test decomposition of a line", func() {

			g.Timeout(1 * time.Hour)
			var options = &opts.Opts{
				Threshold:              50.0,
				MinDist:                20.0,
				NonPlanarDisplacement:  30.0,
				PlanarSelf:             true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}

			var sqrRelation = func(val float64) bool {
				var t = options.Threshold
				return val <= t*t
			}
			// self.relates = relations(self)
			var wkt = "LINESTRING ( 470 480, 470 450, 490 430, 520 420, 540 440, 560 430, 580 420, 590 410, 630 400, 630 430, 640 460, 630 490, 630 520, 640 540, 660 560, 690 580, 700 600, 730 600, 750 570, 780 560, 790 550, 800 520, 830 500, 840 480, 850 460, 900 440, 920 440, 950 480, 990 480, 1000 520, 1000 570, 990 600, 1010 620, 1060 600 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates
			var poly = &geometry.Polyline{}
			var inst = &dpTest{
				Pln:     geometry.CreatePolyline("0", coords, ""),
				Opts:    options,
				ScoreFn: offset.SquareMaxOffset,
			}
			var decomp = offset.EpsilonDecomposition{
				ScoreFn:  inst.ScoreFn,
				Relation: sqrRelation,
			}
			var hulls = DouglasPeucker(id, poly, decomp, hullGeom, inst)
			g.Assert(len(hulls)).Equal(0)

			inst.Opts.Threshold = 120
			hulls = DouglasPeucker(id, inst.Polyline(), decomp, hullGeom, inst)
			g.Assert(len(hulls)).Equal(4)

			inst.Opts.Threshold = 150
			hulls = DouglasPeucker(id, inst.Polyline(), decomp, hullGeom, inst)

			g.Assert(len(hulls)).Equal(1)
			g.Assert(hulls[0].Range.AsSlice()).Equal([]int{0, coords.Len() - 1})
		})
	})
}
