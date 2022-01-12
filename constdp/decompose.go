package constdp

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/decompose"
	"github.com/TopoSimplify/plugin/node"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/intdxdt/iter"
)

func (self *ConstDP) Decompose(id *iter.Igen) []node.Node {
	var score = self.Score
	var relation = self.ScoreRelation

	if self.SquareScore != nil {
		score = self.SquareScore
		relation = self.SquareScoreRelation
	}

	var decomp = offset.EpsilonDecomposition{
		ScoreFn:  score,
		Relation: relation,
	}

	return decompose.DouglasPeucker(
		id, self.Polyline(), decomp, common.Geometry, self,
	)
}
