package geometry

import (
	"github.com/TopoSimplify/plugin/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

//Polyline Type
type Polyline struct {
	*geom.LineString
	Id     string
	Meta   string
	Simple geom.Coords
}

//CreatePolyline construct new polyline
func CreatePolyline(id string, coordinates geom.Coords, meta string) *Polyline {
	if coordinates.Len() < 2 {
		return &Polyline{nil, id, meta, geom.Coords{}}
	}
	return &Polyline{geom.NewLineString(coordinates), id, meta, geom.Coords{}}
}

//SegmentBounds computes segment bounds
func (pln *Polyline) SegmentBounds() []mono.MBR {
	var I, J int
	var n = pln.Len() - 1
	var a, b *geom.Point
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		a, b = pln.Coordinates.Pt(i), pln.Coordinates.Pt(i+1)
		I, J = pln.Coordinates.Idxs[i], pln.Coordinates.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(a[geom.X], a[geom.Y], b[geom.X], b[geom.Y]),
			I:   I, J: J,
		})
	}
	return items
}

//Range of entire polyline
func (pln *Polyline) Range() rng.Rng {
	return rng.Range(pln.Coordinates.FirstIndex(), pln.Coordinates.LastIndex())
}

//Segment given range
func (pln *Polyline) Segment(i, j int) *geom.Segment {
	return geom.NewSegment(pln.Coordinates, i, j)
}

//SubPolyline - generates sub polyline from generator indices
func (pln *Polyline) SubPolyline(rng rng.Rng) *Polyline {
	return CreatePolyline(pln.Id, pln.SubCoordinates(rng), pln.Meta)
}

//SubCoordinates - generates sub polyline from generator indices
func (pln *Polyline) SubCoordinates(rng rng.Rng) geom.Coords {
	var coords = pln.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

//Len - number of coordinates in polyline
func (pln *Polyline) Len() int {
	return pln.Coordinates.Len()
}
