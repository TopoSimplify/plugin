package relate

import (
	"github.com/TopoSimplify/plugin/ctx"
	"github.com/TopoSimplify/plugin/node"
)

//IsGeomRelateValid - Node geometric relation to other context geometries
func IsGeomRelateValid(hull *node.Node, contexts *ctx.ContextGeometries) bool {
	var seg = hull.Segment()
	var lnGeom = hull.Polyline.Geometry()
	var segGeom = seg
	var lnGInter, segGInter bool
	var g *ctx.ContextGeometry

	var bln = true
	var geometries = contexts.DataView()

	for i, n := 0, contexts.Len(); bln && i < n; i++ {
		g = geometries[i]
		lnGInter = lnGeom.Intersects(g.Geom.Geometry())
		segGInter = segGeom.Intersects(g.Geom.Geometry())

		bln = !((segGInter && !lnGInter) || (!segGInter && lnGInter))
	}

	return bln
}
