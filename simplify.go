package main

import (
	"github.com/TopoSimplify/plugin/constdp"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

func simplifyInstances(plns []*geometry.Polyline, opts *opts.Opts,
	constraints []geometry.IGeometry, offsetFn func(geom.Coords) (int, float64)) {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	var junctions = make(map[int][]int, 0)

	for _, pln := range plns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), pln, constraints, opts, offsetFn,
		))
	}

	constdp.SimplifyInstances(id, forest, junctions)
	setSimpleCoords(forest)
}

func simplifyFeatureClass(
	lns []*geometry.Polyline, opts *opts.Opts, constraints []geometry.IGeometry,
	offsetFn func(geom.Coords) (int, float64)) {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	for _, ln := range lns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), ln, constraints, opts, offsetFn,
		))
	}

	constdp.SimplifyFeatureClass(id, forest, opts)
	setSimpleCoords(forest)
}

func setSimpleCoords(forest []*constdp.ConstDP) {
	for _, tree := range forest {
		var coords = tree.Coordinates().Clone()
		coords.Idxs = make([]int, 0, tree.SimpleSet.Size())
		for _, o := range tree.SimpleSet.Values() {
			coords.Idxs = append(coords.Idxs, o.(int))
		}
		tree.Polyline.Simple = coords
	}
}
