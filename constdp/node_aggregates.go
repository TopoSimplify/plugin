package constdp

import (
	"github.com/TopoSimplify/plugin/common"
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/knn"
	"github.com/TopoSimplify/plugin/merge"
	"github.com/TopoSimplify/plugin/node"
	"github.com/intdxdt/iter"
	"sort"
)

//Merge segment fragments where possible
func (self *ConstDP) AggregateSimpleSegments(
	id *iter.Igen, db *hdb.Hdb, constVertexSet []int,
	validateMerge func(*node.Node, *hdb.Hdb) bool,
) {
	var fragmentSize = 1
	var neighbours []*node.Node
	var cache = make(map[[CacheKeySize]int]bool)
	var objects = db.All()
	var hull *node.Node
	sort.Sort(node.NodePtrs(objects))

	var score = self.Score
	var relation = self.ScoreRelation

	if self.SquareScore != nil {
		score = self.SquareScore
		relation = self.SquareScoreRelation
	}

	for len(objects) != 0 {
		hull = objects[0]
		objects = objects[1:]
		if hull.Range.Size() != fragmentSize {
			continue
		}

		//make sure hull index is not part of vertex with degree > 2
		if iter.SortedSearchInts(constVertexSet, hull.Range.I) &&
			iter.SortedSearchInts(constVertexSet, hull.Range.J) {
			continue
		}
		var withNext, withPrev = !iter.SortedSearchInts(constVertexSet, hull.Range.J),
			!iter.SortedSearchInts(constVertexSet, hull.Range.I)

		db.Remove(hull)

		// find context neighbours
		//neighbours = common.NodesFromObjects(knn.FindNodeNeighbours(db, hull, knn.EpsilonDist))
		neighbours = knn.NodeNeighbours(db, hull, knn.EpsilonDist)

		// find context neighbours
		var prev, nxt = node.Neighbours(hull, neighbours)

		// find mergeable neighbours contiguous
		var key [CacheKeySize]int
		var mergePrev, mergeNxt node.Node
		var mergeprevBln, mergenxtBln bool

		if withPrev && prev != nil {
			key = CacheKey(prev, hull)
			if !cache[key] {
				cache[key] = true
				mergeprevBln, mergePrev = merge.ContiguousFragmentsAtThreshold(
					id, score, prev, hull, relation, common.Geometry,
				)
			}
		}

		if withNext && nxt != nil {
			key = CacheKey(hull, nxt)
			if !cache[key] {
				cache[key] = true
				mergenxtBln, mergeNxt = merge.ContiguousFragmentsAtThreshold(
					id, score, hull, nxt, relation, common.Geometry,
				)
			}
		}

		var insertList []node.Node
		var merged bool
		//nxt, prev
		if !merged && mergenxtBln {
			db.Remove(nxt)
			if validateMerge(&mergeNxt, db) {
				if objects[0] == nxt {
					objects = objects[1:]
				}
				insertList = append(insertList, mergeNxt)
				merged = true
			} else {
				merged = false
				insertList = append(insertList, *nxt)
			}
		}

		if !merged && mergeprevBln {
			db.Remove(prev)
			//prev cannot exist since moving from left --- right
			if validateMerge(&mergePrev, db) {
				insertList = append(insertList, mergePrev)
				merged = true
			} else {
				merged = false
				insertList = append(insertList, *prev)
			}
		}

		if !merged {
			insertList = append(insertList, *hull)
		}

		db.Load(insertList)
	}
}
