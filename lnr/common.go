package lnr

import (
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/state"
	"github.com/intdxdt/geom"
)

const NullFId = -9

const (
	x = iota
	y
)

type ScoreFn func(geom.Coords) (int, float64)

type Polygonal interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
}

type Linegen interface {
	Id() int
	Options() *opts.Opts
	Simple() []int
	State() *state.State
}

type Linear interface {
	Polygonal
	Linegen
}
