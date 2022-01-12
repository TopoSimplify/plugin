package constdp

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/intdxdt/geom"
)

//ConstDP Type
type ConstDP struct {
	*dp.DouglasPeucker
	ContextDB *hdb.Hdb
}

//NewConstDP - creates a new constrained DP Simplification instance
//	dp decomposition of linear geometries
func NewConstDP(id int, coordinates geom.Coords,
	constraints []geometry.IGeometry, options *opts.Opts,
	offsetScore lnr.ScoreFn, squareOffsetScore ...lnr.ScoreFn) *ConstDP {
	var sqrScore lnr.ScoreFn
	if len(squareOffsetScore) > 0 {
		sqrScore = squareOffsetScore[0]
	}

	var instance = ConstDP{
		DouglasPeucker: dp.New(id, coordinates, options, offsetScore, sqrScore),
		ContextDB:      hdb.NewHdb(),
	}
	instance.BuildContextDB(constraints)

	if coordinates.Len() > 1 {
		instance.Pln = pln.CreatePolyline(coordinates)
	}
	return &instance
}

//BuildContextDB - creates constraint db from geometries
func (cdp *ConstDP) BuildContextDB(geoms []geometry.IGeometry) *ConstDP {
	var lst = make([]node.Node, 0, len(geoms))
	for i := range geoms {
		cg := ctx.New(geoms[i], 0, -1).AsContextNeighbour()
		lst = append(lst, node.Node{
			MBR:      cg.Bounds(),
			Geom:     cg,
			Instance: cdp,
		})
	}
	cdp.ContextDB.Clear().Load(lst)
	return cdp
}
