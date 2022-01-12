package dp

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/decompose"
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/state"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/sset"
)

//Type DP
type DouglasPeucker struct {
	id          int
	Hulls       []node.Node
	Pln         pln.Polyline
	Meta        map[string]interface{}
	Opts        *opts.Opts
	Score       lnr.ScoreFn
	SquareScore lnr.ScoreFn
	SimpleSet   *sset.SSet
	state       state.State
}

//Creates a new constrained DP Simplification instance
func New(
	id int,
	coordinates geom.Coords,
	options *opts.Opts,
	offsetScore lnr.ScoreFn,
	squareOffsetScore ...lnr.ScoreFn,
) *DouglasPeucker {

	var sqrScore lnr.ScoreFn
	if len(squareOffsetScore) > 0 {
		sqrScore = squareOffsetScore[0]
	}

	var instance = DouglasPeucker{
		id:          id,
		Opts:        options,
		Meta:        make(map[string]interface{}, 0),
		SimpleSet:   sset.NewSSet(cmp.Int),
		Score:       offsetScore,
		SquareScore: sqrScore,
	}

	if coordinates.Len() > 1 {
		instance.Pln = pln.CreatePolyline(coordinates)
	}
	return &instance
}

func (dp *DouglasPeucker) ScoreRelation(val float64) bool {
	return val <= dp.Opts.Threshold
}

func (dp *DouglasPeucker) SquareScoreRelation(val float64) bool {
	return val <= (dp.Opts.Threshold * dp.Opts.Threshold)
}

func (dp *DouglasPeucker) Decompose(id *iter.Igen) []node.Node {
	var score = dp.Score
	var relation = dp.ScoreRelation
	if dp.SquareScore != nil {
		score = dp.SquareScore
		relation = dp.SquareScoreRelation
	}
	var decomp = offset.EpsilonDecomposition{ScoreFn: score, Relation: relation}
	return decompose.DouglasPeucker(
		id, dp.Polyline(), decomp, common.Geometry, dp,
	)
}

func (dp *DouglasPeucker) Simplify(id *iter.Igen) *DouglasPeucker {
	dp.Hulls = dp.Decompose(id)
	return dp
}

func (dp *DouglasPeucker) Simple() []int {
	dp.SimpleSet.Empty()
	for i := range dp.Hulls {
		dp.SimpleSet.Extend(dp.Hulls[i].Range.I, dp.Hulls[i].Range.J)
	}
	var indices = make([]int, dp.SimpleSet.Size())

	dp.SimpleSet.ForEach(func(v interface{}, i int) bool {
		indices[i] = v.(int)
		return true
	})
	return indices
}

func (dp *DouglasPeucker) Id() int {
	return dp.id
}

func (dp *DouglasPeucker) State() *state.State {
	return &dp.state
}

func (dp *DouglasPeucker) Options() *opts.Opts {
	return dp.Opts
}

func (dp *DouglasPeucker) Coordinates() geom.Coords {
	return dp.Pln.Coordinates
}

func (dp *DouglasPeucker) Polyline() pln.Polyline {
	return dp.Pln
}
