package lnr

import (
	"github.com/intdxdt/geom"
	"sort"
)

type vertex struct {
	*geom.Point
	index int
	fid   int
}

type vertices []vertex

//lexical sort of x and y coordinates
func (v vertices) Less(i, j int) bool {
		return (v[i].Point[x] < v[j].Point[x]) || (
			v[i].Point[x] == v[j].Point[x] &&
				v[i].Point[y] < v[j].Point[y])
}

//Len for sort interface
func (v vertices) Len() int {
	return len(v)
}

//Swap for sort interface
func (v vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

//In place lexicographic sort
func (v vertices) Sort() {
	sort.Sort(v)
}

func appendVertices(points []vertex, coordinates geom.Coords, fid int) []vertex {
	for i := range coordinates.Idxs {
		points = append(points, vertex{coordinates.Pt(i), i, fid})
	}
	return points
}
