package knn

import (
	"github.com/TopoSimplify/plugin/hdb"
	"github.com/TopoSimplify/plugin/node"
	"github.com/intdxdt/geom"
)

//ContextNeighbours - find context neighbours by a certain distance
func ContextNeighbours(database *hdb.Hdb, query geom.Geometry, dist float64) []*node.Node {
	return find(database, query, dist, ScoreFn(query))
}

//NodeNeighbours - find context hulls
func NodeNeighbours(database *hdb.Hdb, hull *node.Node, dist float64) []*node.Node {
	var inters = database.Search(*hull.BBox())
	var ns = make([]*node.Node, 0, len(inters))
	for _, nd := range inters {
		if nd.Id != hull.Id && nd.Geom.Distance(hull.Geom) <= dist {
			ns = append(ns, nd)
		}
	}
	return ns
}
