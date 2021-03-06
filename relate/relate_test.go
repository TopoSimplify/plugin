package relate

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

func TestRelate(t *testing.T) {
	var g = goblin.Goblin(t)
	var idgen = iter.NewIgen()
	g.Describe("test relate", func() {
		g.It("should test relate", func() {
			g.Timeout(1 * time.Hour)
			options := &opts.Opts{
				Threshold:              50.0,
				MinDist:                20.0,
				NonPlanarDisplacement:  30.0,
				PlanarSelf:             true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}
			var wkt = "LINESTRING ( 670 550, 680 580, 750 590, 760 630, 830 640, 870 630, 890 610, 920 580, 910 540, 890 500, 900 460, 870 420, 860 390, 810 360, 770 400, 760 420, 800 440, 810 470, 850 500, 820 560, 780 570, 760 530, 720 530, 707.3112236920351 500.3928552814154, 650 450 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates
			var insDP = &dp.DouglasPeucker{Polyline: geometry.CreatePolyline("0", coords, ""), Opts: options}
			var ranges = [][]int{{0, 12}, {12, 18}, {18, coords.Len() - 1}}
			var inst = dp.New(idgen.Next(), geometry.CreatePolyline("0", coords, ""), options, offset.MaxOffset)

			var hulls = common.CreateHulls(idgen, ranges, coords, inst)
			var neib = geom.NewPolygonFromWKT("POLYGON ((674.7409300316725 422.8229196659948, 674.7409300316725 446.72732507918226, 691.3886409444281 446.72732507918226, 691.3886409444281 422.8229196659948, 674.7409300316725 422.8229196659948))")
			const_geom := ctx.New(neib, 0, -1).AsContextNeighbour().AsContextGeometries()
			for i := range hulls {
				g.Assert(IsGeomRelateValid(&hulls[i], const_geom)).IsTrue()
				g.Assert(IsDirRelateValid(&hulls[i], const_geom)).IsTrue()
				g.Assert(IsDistRelateValid(insDP.Options(), &hulls[i], const_geom)).IsTrue()
			}

			neib = geom.NewPolygonFromWKT("POLYGON ((800 614.9282601093252, 800 640, 816.138388266816 640, 816.138388266816 614.9282601093252, 800 614.9282601093252))")
			const_geom = ctx.New(neib, 0, -1).AsContextNeighbour().AsContextGeometries()
			g.Assert(IsGeomRelateValid(&hulls[0], const_geom)).IsFalse()
			g.Assert(IsGeomRelateValid(&hulls[1], const_geom)).IsTrue()
			g.Assert(IsGeomRelateValid(&hulls[2], const_geom)).IsTrue()

			neib = geom.NewPolygonFromWKT("POLYGON ((749.9625484910762 464.581584548546, 749.9625484910762 486.30832777325406, 762.1390749137147 486.30832777325406, 762.1390749137147 464.581584548546, 749.9625484910762 464.581584548546))")
			const_geom = ctx.New(neib, 0, -1).AsContextNeighbour().AsContextGeometries()
			g.Assert(IsGeomRelateValid(&hulls[0], const_geom)).IsFalse()
			g.Assert(IsGeomRelateValid(&hulls[1], const_geom)).IsTrue()
			g.Assert(IsGeomRelateValid(&hulls[2], const_geom)).IsFalse()

			g.Assert(IsDistRelateValid(insDP.Options(), &hulls[0], const_geom)).IsFalse()
			g.Assert(IsDistRelateValid(insDP.Options(), &hulls[1], const_geom)).IsTrue()
			g.Assert(IsDistRelateValid(insDP.Options(), &hulls[2], const_geom)).IsFalse()
		})
	})
}
