package box

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)

func MBRToPolygon(o mbr.MBR) *geom.Polygon {
	return geom.NewPolygon(geom.Coordinates(
		[]geom.Point{
			{o.MinX, o.MinY},
			{o.MinX, o.MaxY},
			{o.MaxX, o.MaxY},
			{o.MaxX, o.MinY},
			{o.MinX, o.MinY},
		},
	))
}
