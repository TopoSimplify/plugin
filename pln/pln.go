package pln

import (
	"github.com/TopoSimplify/rng"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

//Polyline Type
type Polyline struct {
	*geom.LineString
}

//construct new polyline
func CreatePolyline(coordinates geom.Coords) Polyline {
	return Polyline{geom.NewLineString(coordinates)}
}

//Polyline segments
func (ln *Polyline) SegmentBounds() []mono.MBR {
	var I, J int
	var n = ln.Len() - 1
	var a, b *geom.Point
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		a, b = ln.Coordinates.Pt(i), ln.Coordinates.Pt(i+1)
		I, J = ln.Coordinates.Idxs[i], ln.Coordinates.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(a[geom.X], a[geom.Y], b[geom.X], b[geom.Y]),
			I:   I, J: J,
		})
	}
	return items
}

//Range of entire polyline
func (ln *Polyline) Range() rng.Rng {
	return rng.Range(ln.Coordinates.FirstIndex(), ln.Coordinates.LastIndex())
}

//Segment given range
func (ln *Polyline) Segment(i, j int) *geom.Segment {
	return geom.NewSegment(ln.Coordinates, i, j)
}

//generates sub polyline from generator indices
func (ln *Polyline) SubPolyline(rng rng.Rng) Polyline {
	return CreatePolyline(ln.SubCoordinates(rng))
}

//generates sub polyline from generator indices
func (ln *Polyline) SubCoordinates(rng rng.Rng) geom.Coords {
	var coords = ln.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

//Length of coordinates in polyline
func (ln *Polyline) Len() int {
	return ln.Coordinates.Len()
}
