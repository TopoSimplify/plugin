package ctx

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)

const (
	Self             = "self"
	PlanarVertex     = "planar_vertex"
	NonPlanarVertex  = "non_planar_vertex"
	PlanarSegment    = "planar_segment"
	LinearSimple     = "linear_simple"
	ContextNeighbour = "context_neighbour"
)

type Meta struct {
	Planar    []int
	NonPlanar []int
}

type ContextGeometry struct {
	Geom    geom.Geometry
	CtxType string
	I       int
	J       int
	Meta    Meta
}

func New(g geom.Geometry, i, j int) *ContextGeometry {
	return &ContextGeometry{
		Geom:    g,
		CtxType: Self,
		I:       i,
		J:       j,
		Meta:    Meta{},
	}
}

//implements stringer interface
func (o *ContextGeometry) String() string {
	return o.Geom.WKT()
}

//implements IGeom interface
func (o *ContextGeometry) Geometry() geom.Geometry {
	return o.Geom
}

//BBox
func (o *ContextGeometry) BBox() *mbr.MBR {
	return o.Geom.BBox()
}

//Bounds
func (o *ContextGeometry) Bounds() mbr.MBR {
	return *o.BBox()
}

//AsLinear
func (o *ContextGeometry) AsLinear() []*geom.LineString {
	return o.Geom.AsLinear()
}

//Intersects
func (o *ContextGeometry) Intersects(other geom.Geometry) bool {
	return o.Geom.Intersects(other)
}

//Intersection
func (o *ContextGeometry) Intersection(other geom.Geometry) []geom.Point {
	return o.Geom.Intersection(other)
}

//Distance
func (o *ContextGeometry) Distance(other geom.Geometry) float64 {
	return o.Geom.Distance(other)
}

//Type
func (o *ContextGeometry) Type() geom.GeoType {
	return o.Geom.Type()
}

//Distance
func (o *ContextGeometry) WKT() string {
	return o.Geom.WKT()
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsSelf() *ContextGeometry {
	o.CtxType = Self
	return o
}

func (o *ContextGeometry) IsSelf() bool {
	return o.CtxType == Self
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsPlanarVertex() *ContextGeometry {
	o.CtxType = PlanarVertex
	return o
}

func (o *ContextGeometry) IsPlanarVertex() bool {
	return o.CtxType == PlanarVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsNonPlanarVertex() *ContextGeometry {
	o.CtxType = NonPlanarVertex
	return o
}

func (o *ContextGeometry) IsNonPlanarVertex() bool {
	return o.CtxType == NonPlanarVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsPlanarSegment() *ContextGeometry {
	o.CtxType = PlanarSegment
	return o
}

func (o *ContextGeometry) IsPlanarSegment() bool {
	return o.CtxType == PlanarSegment
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsLinearSimple() *ContextGeometry {
	o.CtxType = LinearSimple
	return o
}

func (o *ContextGeometry) IsLinearSimple() bool {
	return o.CtxType == LinearSimple
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsContextGeometries(objects ...*ContextGeometry) *ContextGeometries {
	return NewContexts(o).Extend(objects)
}

func (o *ContextGeometry) AsContextNeighbour() *ContextGeometry {
	o.CtxType = ContextNeighbour
	return o
}

func (o *ContextGeometry) IsContextNeighbour() bool {
	return o.CtxType == ContextNeighbour
}

//--------------------------------------------------------------------
