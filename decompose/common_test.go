package decompose

import (
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/state"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
)

//hull geom
func hullGeom(coords geom.Coords) geom.Geometry {
	var g geom.Geometry
	var n = coords.Len()
	if n > 2 {
		g = geom.NewPolygon(coords)
	} else if n == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords.Pt(0)
	}
	return g
}

//Type DP
type dpTest struct {
	id        int
	Hulls     *deque.Deque
	Pln       pln.Polyline
	Meta      map[string]interface{}
	Opts      *opts.Opts
	ScoreFn   lnr.ScoreFn
	SimpleSet *sset.SSet
}

func (self *dpTest) Id() int {
	return self.id
}

func (self *dpTest) State() *state.State {
	var s state.State
	return &s
}

func (self *dpTest) Options() *opts.Opts {
	return self.Opts
}

func (self *dpTest) NodeQueue() *deque.Deque {
	return self.Hulls
}

func (self *dpTest) Simple() []int {
	return []int{}
}

func (self *dpTest) Coordinates() geom.Coords {
	return self.Pln.Coordinates
}

func (self *dpTest) Polyline() pln.Polyline {
	return self.Pln
}
