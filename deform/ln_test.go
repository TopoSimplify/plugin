package deform

import (
	"fmt"
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/TopoSimplify/plugin/state"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"testing"
	"time"
)

type testDP struct {
	id int
}

func (tdp *testDP) Id() int {
	return tdp.id
}

func (tdp *testDP) State() *state.State {
	var s state.State
	return &s
}

func (tdp *testDP) Options() *opts.Opts {
	return nil
}
func (tdp *testDP) Simple() []int {
	return nil
}

func TestDeform(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen()
	var wktDat = []struct {
		ranges  [][]int
		q       int
		expects []int
		wkt     string
	}{
		{[][]int{{0, 12}, {12, 18}, {18, -1}}, 0, []int{0},
			"LINESTRING ( 670 550, 680 580, 750 590, 760 630, 830 640, 870 630, 890 610, 920 580, 910 540, 890 500, 900 460, 870 420, 860 390, 810 360, 770 400, 760 420, 800 440, 810 470, 850 500, 820 560, 780 570, 760 530, 720 530, 707.3112236920351 500.3928552814154, 650 450 )"},
		{[][]int{{0, 21}, {21, -1}}, 0, []int{1},
			"LINESTRING ( 810 540, 790 570, 800 580, 820 580, 860 570, 880 600, 870 610, 850 610, 800 610, 810 650, 890 640, 900 640, 920 600, 930 580, 930 540, 920 500, 880 490, 860 520, 810 510, 750 520, 780 460, 730 410, 830 440, 890 410, 940 450, 970 500, 1040 510, 1050 570, 1080 620, 1040 660, 1020 720, 950 720, 840 680, 760 690, 690 720, 710 640, 630 620 )"},
		{[][]int{{0, 15}, {15, -1}}, 0, []int{0},
			"LINESTRING ( 630 620, 710 640, 690 720, 760 690, 840 680, 950 720, 1020 720, 1040 660, 1080 620, 1050 570, 1040 510, 970 500, 940 450, 890 410, 830 440, 730 410, 780 460, 750 520, 810 510, 860 520, 880 490, 920 500, 930 540, 930 580, 920 600, 900 640, 890 640, 810 650, 800 610, 850 610, 870 610, 880 600, 860 570, 820 580, 800 580, 790 570, 810 540 )"},
		{[][]int{{0, 13}, {13, -1}}, 0, []int{},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 920 600, 960 580, 960 560, 960 540, 940 540, 910 550, 890 580, 870 600, 850 610, 840 570, 850 550, 860 530, 860 500, 930 490, 960 480, 990 480, 1010 500, 1060 530, 1090 580 )"},
		{[][]int{{0, 13}, {13, -1}}, 0, []int{},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 920 600, 960 580, 960 560, 960 540, 940 540, 910 550, 890 580, 870 600, 850 610, 840 570, 850 550, 860 530, 860 500, 820 470 )"},
		{[][]int{{0, 13}, {13, 20}, {20, -1}}, 2, []int{0},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 930 610, 960 580, 960 560, 960 540, 940 540, 910 540, 900 570, 910 580, 890 600, 870 600, 840 620, 820 590, 840 580, 850 570, 850 550, 860 540, 850 520, 870 510, 880 500, 860 480, 840 490, 820 460, 790 450 )"},
		{[][]int{{0, 12}, {12, 19}, {19, -1}}, 0, []int{2},
			"LINESTRING (790 450, 820 460, 840 490, 860 480, 880 500, 870 510, 850 520, 860 540, 850 550, 850 570, 840 580, 820 590, 840 620, 870 600, 890 600, 910 580, 900 570, 910 540, 940 540, 960 540, 960 560, 960 580, 930 610, 930 650, 910 690, 860 690, 780 680, 760 660, 780 630, 770 590, 750 550, 730 520, 730 490 )"},
		{[][]int{{0, 12}, {12, 19}, {19, -1}}, 0, []int{0, 2},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 930 610, 960 580, 960 560, 960 540, 940 510, 910 490, 900 500, 900 560, 870 550, 870 520, 840 520, 820 520, 800 570, 810 610, 830 630, 840 640, 850 650, 870 650, 910 660, 960 670, 1000 670, 1020 650, 1030 630 )"},
		{[][]int{{0, 12}, {12, 19}, {19, -1}}, 0, []int{0, 2},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 930 610, 960 580, 960 560, 960 540, 940 510, 910 490, 900 500, 900 560, 870 550, 870 520, 840 520, 820 520, 800 570, 810 610, 830 630, 840 640, 850 650, 870 650, 910 660, 960 670, 1020 650, 1040 620, 1060 570 )"},
		{[][]int{{12, 19}, {19, -1}}, 1, []int{1},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 930 610, 960 580, 960 560, 960 540, 940 510, 910 490, 900 500, 900 560, 870 550, 870 520, 840 520, 820 520, 800 570, 810 610, 830 630, 840 640, 850 650, 870 650, 910 660, 960 670, 1020 650, 1040 620, 1060 570 )"},
		{[][]int{{0, 12}, {12, 19}, {19, -1}}, 2, []int{},
			"LINESTRING ( 730 490, 730 520, 750 550, 770 590, 780 630, 760 660, 780 680, 860 690, 910 690, 930 650, 930 610, 960 580, 960 560, 960 540, 940 510, 910 490, 900 500, 900 560, 870 550, 870 520, 850 500, 830 500, 800 480, 740 460, 710 470, 670 500, 660 470, 670 440, 700 420, 730 400, 860 390, 890 390, 910 420 )"},
	}

	var createHullsDbTest = func(ranges [][]int, coordinates geom.Coords) ([]node.Node, *hdb.Hdb) {
		var inst = &testDP{0}
		var n = coordinates.Len()
		var polyline = geometry.CreatePolyline("0", coordinates, "")
		var hulls []node.Node
		for _, r := range ranges {
			var i, j = r[0], r[len(r)-1]
			if j == -1 {
				j = n - 1
			}
			var nr = rng.Range(i, j)
			var h = node.CreateNode(id, polyline.SubCoordinates(nr), nr, common.Geometry, inst)
			hulls = append(hulls, h)
		}

		return hulls, hdb.NewHdb().Load(hulls)
	}

	g.Describe("hull deformation", func() {
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

		var cdp = &dp.DouglasPeucker{Opts: options, Score: offset.MaxOffset}

		g.It("should test selection of hulls for deformation", func() {
			g.Timeout(60 * time.Minute)
			var contains = func(s *node.Node, slns []*node.Node) bool {
				bln := false
				for _, h := range slns {
					if s == h {
						bln = true
						break
					}
				}
				return bln
			}
			for _, o := range wktDat {
				var ranges, q, expects, wkt = o.ranges, o.q, o.expects, o.wkt
				var coords = geom.NewLineStringFromWKT(wkt).Coordinates
				var hulls, hulldb = createHullsDbTest(ranges, coords)

				var query = hulls[q]

				var slns = Select(cdp.Options(), hulldb, &query)
				//slns = select_hulls_to_deform(ha, hb, opts)
				if len(slns) != len(expects) {
					fmt.Println(slns)
				}
				g.Assert(len(slns)).Equal(len(expects))
				for _, i := range expects {
					g.Assert(contains(&hulls[i], slns))
				}
			}
		})
	})
}
