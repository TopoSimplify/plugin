package hdb

import (
	"sort"
	"github.com/intdxdt/mbr"
)

//calcBBox calculates its bbox from bboxes of its children.
func calcBBox(nd *dbNode) {
	nd.bbox = distBBox(nd, 0, len(nd.children))
}

//distBBox computes min bounding rectangle of dbNode children from k to p-1.
func distBBox(nd *dbNode, k, p int) mbr.MBR {
	var bbox = emptyMBR()
	for i := k; i < p; i++ {
		bbox.ExpandIncludeMBR(&nd.children[i].bbox)
	}
	return bbox
}

//allDistMargin computes total margin of all possible split distributions.
//Each dbNode is at least m full.
func (tree *Hdb) allDistMargin(nd *dbNode, m, M int, sortBy sortBy) float64 {
	if sortBy == byX {
		sort.Sort(xNodePath{nd.children})
		//bubbleAxis(*dbNode.getChildren(), byX, byY)
	} else if sortBy == byY {
		sort.Sort(yNodePath{nd.children})
		//bubbleAxis(*dbNode.getChildren(), byY, byX)
	}

	var i int
	var leftBBox = distBBox(nd, 0, m)
	var rightBBox = distBBox(nd, M-m, M)
	var margin = bboxMargin(&leftBBox) + bboxMargin(&rightBBox)

	for i = m; i < M-m; i++ {
		leftBBox.ExpandIncludeMBR(&nd.children[i].bbox)
		margin += bboxMargin(&leftBBox)
	}

	for i = M - m - 1; i >= m; i-- {
		rightBBox.ExpandIncludeMBR(&nd.children[i].bbox)
		margin += bboxMargin(&rightBBox)
	}
	return margin
}
