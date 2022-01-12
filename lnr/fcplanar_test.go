package lnr

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"sort"
	"testing"
	"time"
)

func TestFCSelfPlanarIntersects(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("fc planar self intersection", func() {
		g.It("fc planar intersects case 1", func() {
			g.Timeout(1 * time.Hour)
			var wkts = []string{
				"LINESTRING ( -300 -500, -200 -600, -100 -500, -100 -700, 200 -700, 279.69268484251296 -831.3596549207016, 600 -700, 600 -499.3043108842079, 300 -400, 200 -500, 279.69268484251296 -831.3596549207016, 657.2870474897259 -844.6418686821614, 839.4431219326025 -745.9739950256032 )",
				"LINESTRING ( -300 -1100, -300 -900, -100 -900, -100 -800, 279.69268484251296 -831.3596549207016, 500 -1000, 400 -1100, 400 -1200 )",
				"LINESTRING ( 522.1337339611549 -1142.8294428402041, 657.2870474897259 -844.6418686821614, 839.4431219326025 -745.9739950256032, 972.4544063237287 -952.9896735474056, 1038.6763283544988 -1133.0556532167159 )",
			}

			var fcs []*FC
			for fid, wkt := range wkts {
				o := NewFC(geom.NewLineStringFromWKT(wkt).Coordinates, fid)
				fcs = append(fcs, o)
			}

			var inters = FCPlanarSelfIntersection(fcs)
			for _, o := range inters {
				sort.Ints(o)
			}
			g.Assert(inters[0]).Equal([]int{5, 10, 11, 12})
			g.Assert(inters[1]).Equal([]int{4})
			g.Assert(inters[2]).Equal([]int{1, 2})
		})

		g.It("fc planar intersects case 2", func() {
			g.Timeout(1 * time.Hour)
			var wkts = []string{
				"LINESTRING ( 300 0, 300 400, 600 600, 600 1000, 900 1000, 900 700, 1300 700, 1400 400, 1600 200, 1300 0, 800 100, 300 0 )",
				"LINESTRING ( 100 200, 0 300, 100 500, 100 700, 300 800, 300 1100, 333.48668893714955 1263.8423649803672, 400 1300, 800 1100, 1100 1100, 1100 900, 1200 900, 1300 700, 1600 700, 1500 500, 1700 400, 1630.634600565117 122.53840226046754, 1600 0, 1100 -200, 600 -200 )",
				"LINESTRING ( 100 -100, -100 0, -100 100, -200 200, -200 400, -400 500, -500 400, -600 300, -500 100, -300 100, -200 400, -300 700, -200 800, -200 900, 0 800, 300 1100, 300 1300, 600 1400, 900 1500, 1100 1300, 1400 900, 1700 900, 1800 600, 1800 -200 )",
			}

			var fcs []*FC
			for fid, wkt := range wkts {
				o := NewFC(geom.NewLineStringFromWKT(wkt).Coordinates, fid)
				fcs = append(fcs, o)
			}

			var inters = FCPlanarSelfIntersection(fcs)
			for _, o := range inters {
				sort.Ints(o)
			}
			g.Assert(inters[0]).Equal([]int{0, 6, 11})
			g.Assert(inters[1]).Equal([]int{5, 12})
			g.Assert(inters[2]).Equal([]int{4, 10, 15})
		})

		g.It("fc planar intersects case 3", func() {
			g.Timeout(1 * time.Hour)
			var wkts = []string{
				"LINESTRING ( 2014.9435396905058 234.97154854691937, 2093.0032307474535 278.8801247664523, 2187.8674386291605 239.85027923797858, 2113.6718459100757 121.76018213404694, 2237.738907915543 68.01054267511512, 2345.070983118846 176.96886144210424, 2478.9650365290267 117.33993077360273, 2351.575957373592 22.47572289189578, 2024.1730507699667 46.5467070328619, 1924.5190637422045 -30.93600141769867 )",
				"LINESTRING ( 2269.1796168134806 -83.77218993561601, 2365.670068258874 -38.77945134029214, 2478.9650365290267 117.33993077360273, 2334.9112017144043 278.14838499157116, 2187.8674386291605 239.85027923797858, 2258.337993055571 337.42489305916286, 2448.6084900068804 339.0511366228493, 2531.0048305669916 236.59779211060578 )",
			}

			var fcs []*FC
			for fid, wkt := range wkts {
				o := NewFC(geom.NewLineStringFromWKT(wkt).Coordinates, fid)
				fcs = append(fcs, o)
			}

			var inters = FCPlanarSelfIntersection(fcs)
			for _, o := range inters {
				sort.Ints(o)
			}
			g.Assert(inters[0]).Equal([]int{2, 6})
			g.Assert(inters[1]).Equal([]int{2, 4})
		})
	})
}
