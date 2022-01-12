package constdp

import (
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/node"
)

func cleanUpDB(hulldb *hdb.Hdb, selections map[*node.Node]struct{}) {
	for o := range selections {
		hulldb.Remove(o)
	}
}
