package hdb

import (
	"fmt"
	"github.com/intdxdt/mbr"
	"github.com/TopoSimplify/node"
)

//KObj instance struct
type KObj struct {
	dbNode   *dbNode
	MBR      *mbr.MBR
	IsItem   bool
	Distance float64
}

func (kobj *KObj) GetNode() *node.Node {
	return kobj.dbNode.item
}

//String representation of knn object
func (kobj *KObj) String() string {
	return fmt.Sprintf("%v -> %v", kobj.dbNode.bbox.String(), kobj.Distance)
}

//Compare - cmp interface
func kobjCmp(a interface{}, b interface{}) int {
	var self, other = a.(*KObj), b.(*KObj)
	var dx = self.Distance - other.Distance
	var r = 1
	if feq(dx, 0) {
		r = 0
	} else if dx < 0 {
		r = -1
	}
	return r
}
