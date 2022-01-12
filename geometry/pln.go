package geometry

import (
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

//Polyline Type
type Polyline struct {
	G    *geom.LineString
	Id   string
	Meta string
}

//Geometry interface
func (pln Polyline) Geometry() geom.Geometry {
	return pln.G
}

//CreatePolyline construct new polyline
func CreatePolyline(id string, coordinates geom.Coords, meta string) Polyline {
	return Polyline{geom.NewLineString(coordinates), id, meta}
}

//SegmentBounds computes segment bounds
func (pln *Polyline) SegmentBounds() []mono.MBR {
	var I, J int
	var n = pln.Len() - 1
	var a, b *geom.Point
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		a, b = pln.G.Coordinates.Pt(i), pln.G.Coordinates.Pt(i+1)
		I, J = pln.G.Coordinates.Idxs[i], pln.G.Coordinates.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(a[geom.X], a[geom.Y], b[geom.X], b[geom.Y]),
			I:   I, J: J,
		})
	}
	return items
}

//Range of entire polyline
func (pln *Polyline) Range() rng.Rng {
	return rng.Range(pln.G.Coordinates.FirstIndex(), pln.G.Coordinates.LastIndex())
}

//Segment given range
func (pln *Polyline) Segment(i, j int) *geom.Segment {
	return geom.NewSegment(pln.G.Coordinates, i, j)
}

//SubPolyline - generates sub polyline from generator indices
func (pln *Polyline) SubPolyline(rng rng.Rng) Polyline {
	return CreatePolyline(pln.Id, pln.SubCoordinates(rng), pln.Meta)
}

//SubCoordinates - generates sub polyline from generator indices
func (pln *Polyline) SubCoordinates(rng rng.Rng) geom.Coords {
	var coords = pln.G.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

//Len - number of coordinates in polyline
func (pln *Polyline) Len() int {
	return pln.G.Coordinates.Len()
}
