package homotopy

import (
	"github.com/intdxdt/geom"
)

type Vertex struct {
	Point *geom.Point
	i     int
	prev  *Vertex
	next  *Vertex
}

type Chain struct {
	link *Vertex
	size int
	simple *geom.Segment
}

func NewChain(coordinates geom.Coords) *Chain {
	var n = coordinates.Len()
	var chain = &Chain{
		size: n,
		link: &Vertex{Point: coordinates.Pt(0)},
		simple: geom.NewSegment(coordinates, 0, n-1),
	}

	var prev, next *Vertex
	prev = chain.link
	for i := 1; i < chain.size; i++ {
		next = &Vertex{Point: coordinates.Pt(i), i: i, prev: prev}
		ptrs(prev, next)
		prev = next
	}
	return chain
}

func remove(link *Vertex) {
	if link != nil {
		ptrs(link.prev, link.next)
	}
}

//Updates next and prev pointers
func ptrs(prev, next *Vertex) {
	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}
}
