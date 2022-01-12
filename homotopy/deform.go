package homotopy

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
)

//func printChain(chain *Chain) {
//	var coords []geom.Point
//	var link = chain.link
//	for link != nil {
//		coords = append(coords, *link.Point)
//		link = link.next
//	}
//	fmt.Println(geom.NewLineString(coords).WKT())
//}

//deforms a polyline given coordinates and
// disjoint context neighbours
func chainDeformation(
	coordinates geom.Coords,
	contexts *ctx.ContextGeometries) *Chain {

	var db = contextDB(contexts)
	var chain = NewChain(coordinates)
	var deformable = true

	for deformable && chain.size > 2 {
		deformable = false
		var link = chain.link
		for link != nil {
			if collapseVertex(link, db) {
				remove(link)
				chain.size += -1
				deformable = true
			}
			link = link.next
		}
		//printChain(chain)
	}
	return chain
}

func contextDB(contexts *ctx.ContextGeometries) *rtree.RTree {
	var db = rtree.NewRTree()
	var view = contexts.DataView()
	var objects = make([]rtree.BoxObject, 0, len(view))
	for i := range view {
		objects = append(objects, view[i])
	}
	db.Load(objects)
	return db
}

func collapseVertex(v *Vertex, db *rtree.RTree) bool {
	var va, vb, vc = v.prev, v, v.next
	if va == nil || vb == nil || vc == nil {
		return false
	}

	var bln = true
	var a, b, c = va.Point, vb.Point, vc.Point

	var box = a.BBox().ExpandIncludeXY(b[geom.X], b[geom.Y]).ExpandIncludeXY(
		c[geom.X], c[geom.Y],
	)
	var neighbours = db.Search(*box)
	if len(neighbours) > 0 {
		bln = isTriangleCollapsible(a, b, c, neighbours)
	}
	return bln
}

//find if intersects simple
func isTriangleCollapsible(a, b, c *geom.Point, neighbours []rtree.BoxObject) bool {
	var bln = true
	var triangle = geom.NewPolygon(geom.Coordinates([]geom.Point{*a, *b, *c, *a}))
	for i, n := 0, len(neighbours); bln && i < n; i++ {
		c := neighbours[i].(*ctx.ContextGeometry)
		bln = !triangle.Intersects(c.Geom) //disjoint
	}
	return bln
}
