package box

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
	"testing"
)

func TestMBRToPolygon(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("box as polygon", func() {
		g.It("should test mbr to polygon conversion", func() {
			var box = mbr.CreateMBR(0.25, 0.5, 5, 5)
			var pts = make([]geom.Point, 0)
			for _, pt := range box.AsPolyArray() {
				pts = append(pts, geom.CreatePoint(pt))
			}

			ply := geom.NewPolygon(geom.Coordinates(pts))
			g.Assert(box.Area()).Equal(ply.Area())
			g.Assert(MBRToPolygon(box).Area()).Equal(ply.Area())
		})
	})
}
