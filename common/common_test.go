package common

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"testing"
	"time"
)

func TestCommon(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("DP2", func() {
		g.It("dp with self intersection", func() {
			g.Timeout(1 * time.Hour)
			var data = []geom.Point{
				{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8},
				{0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5},
			}
			var ptCoords = geom.Coordinates([]geom.Point{data[0]})
			g.Assert(Geometry(ptCoords).(*geom.Point).Equals2D(&data[0])).IsTrue()
		})
	})
}
