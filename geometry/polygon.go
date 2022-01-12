package geometry

import "github.com/intdxdt/geom"

type Polygon struct {
	G    *geom.Polygon
	Id   string
	Meta string
}

//Geometry interface
func (g Polygon) Geometry() geom.Geometry {
	return g.G
}

//CreatePolygon creates a new Polygon
func CreatePolygon(id string, coordinates []geom.Coords, meta string) Polygon {
	return Polygon{geom.NewPolygon(coordinates...), id, meta}
}
