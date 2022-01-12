package lnr

import (
	"github.com/intdxdt/geom"
)

type FC struct {
	Coordinates geom.Coords
	Fid         int
}

func NewFC(coordinates geom.Coords, fid int) *FC {
	return &FC{coordinates, fid}
}

//Planar self-intersection FCPlanarSelfIntersection
func FCPlanarSelfIntersection(featureClass []*FC) map[int][]int {
	//preallocate points size
	var n = 0
	for _, self := range featureClass {
		n +=  self.Coordinates.Len()
	}
	var points = make([]vertex, 0, n)

	for _, self := range featureClass {
		points = appendVertices(points, self.Coordinates, self.Fid)
	}

	vertices(points).Sort() //O(nlogn)

	var d = 0
	var indexes []*vertex
	var results = make(map[int][]int, 0)
	var a, b *vertex

	var bln bool
	for i, n := 0, len(points); i < n-1; i++ { //O(n)
		a, b = &points[i], &points[i+1]
		bln = a.Equals2D(b.Point)
		if bln {
			if d == 0 {
				indexes = append(indexes, a, b)
				d += 2
			} else {
				indexes = append(indexes, b)
				d += 1
			}
			continue
		}

		if d > 1 {
			for _, v := range indexes {
				results[v.fid] = append(results[v.fid], v.index)
			}
		}
		d = 0
		indexes = indexes[:0]
	}
	return results
}
