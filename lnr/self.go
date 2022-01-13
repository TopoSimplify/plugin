package lnr

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/geom/index"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/iter"
)

//SelfIntersection - Planar and non-planar intersections
func SelfIntersection(polyline geometry.Polyline, planar, nonPlanar bool) *ctx.ContextGeometries {
	var inters = ctx.NewContexts()
	if planar {
		inters.Extend(planarIntersects(polyline).DataView())
	}
	if nonPlanar {
		inters.Extend(nonPlanarIntersection(polyline).DataView())
	}
	return inters
}

//Planar self-intersection
func planarIntersects(polyline geometry.Polyline) *ctx.ContextGeometries {
	var points = make([]vertex, 0, polyline.Coordinates.Len())
	for i := range polyline.Coordinates.Idxs {
		points = append(points, vertex{polyline.Pt(i), i, NullFId})
	}
	vertices(points).Sort() //O(nlogn)

	var d = 0
	var a, b *vertex
	var indices []int
	var results = ctx.NewContexts()

	var bln bool
	for i, n := 0, len(points); i < n-1; i++ { //O(n)
		a, b = &points[i], &points[i+1]
		bln = a.Equals2D(b.Point)
		if bln {
			if d == 0 {
				indices = append(indices, a.index, b.index)
				d += 2
			} else {
				indices = append(indices, b.index)
				d += 1
			}
			continue
		}

		if d > 1 {
			var cg = ctx.New(points[i].Point, 0, -1).AsPlanarVertex()
			cg.Meta.Planar = iter.SortedIntsSet(indices)
			results.Push(cg)
		}
		d = 0
		indices = indices[:0]
	}
	return results
}

func nonPlanarIntersection(polyline geometry.Polyline) *ctx.ContextGeometries {
	var s *mono.MBR
	var cache = make(map[[4]int]bool)
	var tree, data = segmentDB(polyline)
	var results = ctx.NewContexts()
	var neighbours []*mono.MBR
	var sa, sb, oa, ob *geom.Point

	for d := range data {
		s = &data[d]
		neighbours = tree.Search(s.MBR)

		for _, o := range neighbours {
			//var o = obj.Object.(*geom.Segment)
			if s == o {
				continue
			}

			var k = cacheKey(s, o)
			if cache[k] {
				continue
			}
			cache[k] = true

			sa, sb = polyline.Pt(s.I), polyline.Pt(s.J)
			oa, ob = polyline.Pt(o.I), polyline.Pt(o.J)
			var intersects = geom.SegSegIntersection(sa, sb, oa, ob)
			var pt *geom.InterPoint
			for idx := range intersects {
				pt = &intersects[idx]
				if pt.IsVertex() && !pt.IsVerteXOR() { //if not exclusive vertex
					continue
				}
				cg := ctx.New(pt.Point, 0, -1).AsNonPlanarVertex()
				cg.Meta.NonPlanar = iter.SortedIntsSet(k[:])
				results.Push(cg)
			}
		}
	}
	return results
}

//cache key: [0, 1, 9, 10] == [9, 10, 0, 1]
func cacheKey(a, b *mono.MBR) [4]int {
	if b.I < a.I {
		a, b = b, a
	}
	return [4]int{a.I, a.J, b.I, b.J}
}

func segmentDB(polyline geometry.Polyline) (*index.Index, []mono.MBR) {
	var tree = index.NewIndex()
	var data = polyline.SegmentBounds()
	tree.Load(data)
	return tree, data
}
