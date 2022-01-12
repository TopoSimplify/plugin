package node

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
)

type NodePtrs []*Node

func (nptrs NodePtrs) Len() int {
	return len(nptrs)
}

func (nptrs NodePtrs) Swap(i, j int) {
	nptrs[i], nptrs[j] = nptrs[j], nptrs[i]
}

func (nptrs NodePtrs) Less(i, j int) bool {
	return nptrs[i].Range.I < nptrs[j].Range.I
}

type Nodes []Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (nodes Nodes) Less(i, j int) bool {
	return nodes[i].Range.I < nodes[j].Range.I
}

func (nodes Nodes) AsPointSet() *sset.SSet {
	var set = sset.NewSSet(cmp.Int)
	for _, o := range nodes {
		set.Extend(o.Range.I, o.Range.J)
	}
	return set
}

func Pop(self *[]Node) []Node {
	var nodes = *self
	if len(nodes) != 0 {
		n := len(nodes) - 1
		nodes[n] = Node{}
		*self = nodes[:n]
	}
	return nodes
}

func Clear(self *[]Node){
	*self = (*self)[:0]
}
