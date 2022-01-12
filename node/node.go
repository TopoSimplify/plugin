package node

import (
	"github.com/TopoSimplify/plugin/lnr"
	"github.com/TopoSimplify/plugin/pln"
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/mbr"
)

// Node Type
type Node struct {
	Id       int
	Polyline pln.Polyline
	Range    rng.Rng
	MBR      mbr.MBR
	Geom     geom.Geometry
	Instance lnr.Linegen
}

// CreateNode Node
func CreateNode(id *iter.Igen, coordinates geom.Coords, rng rng.Rng,
	geomFn func(geom.Coords) geom.Geometry, instance lnr.Linegen) Node {
	var g = geomFn(geom.ConvexHull(coordinates))
	return Node{
		Id:       id.Next(),
		Polyline: pln.CreatePolyline(coordinates),
		Range:    rng,
		MBR:      g.Bounds(),
		Geom:     g,
		Instance: instance,
	}
}

// Implements bbox interface
func (node *Node) BBox() *mbr.MBR {
	return node.Geom.BBox()
}

// Implements bbox interface
func (node *Node) Bounds() mbr.MBR {
	return node.Geom.Bounds()
}

// Geometry bbox interface
func (node *Node) Geometry() geom.Geometry {
	return node.Geom
}

// stringer interface
func (node *Node) String() string {
	return node.Geom.WKT()
}

// coordinates
func (node *Node) Coordinates() geom.Coords {
	return node.Polyline.Coordinates
}

// first point in coordinates
func (node *Node) First() *geom.Point {
	return node.Polyline.Coordinates.First()
}

// last point in coordinates
func (node *Node) Last() *geom.Point {
	return node.Polyline.Coordinates.Last()
}

// as segment
func (node *Node) Segment() *geom.Segment {
	return geom.NewSegment(
		node.Polyline.Coordinates, 0, node.Polyline.Coordinates.Len()-1,
	)
}

// hull segment as polyline
func (node *Node) SegmentAsPolyline() pln.Polyline {
	var n = node.Polyline.Len()
	var coords = node.Polyline.Coordinates
	coords.Idxs = []int{coords.Idxs[0], coords.Idxs[n-1]}
	return pln.CreatePolyline(coords)
}

//SegmentPoints - segment points
func (node *Node) SegmentPoints() (*geom.Point, *geom.Point) {
	return node.First(), node.Last()
}
