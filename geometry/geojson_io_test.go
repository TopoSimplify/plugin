package geometry

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestGeoJSON_IO_File_Constraints(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GeoJSON - IO", func() {
		g.It("consts from file", func() {
			g.Timeout(1 * time.Hour)
			var igeoms = ReadInputConstraints("data/consts.txt")
			g.Assert(len(igeoms) > 11).IsTrue()
		})
		g.It("linestrings from file", func() {
			g.Timeout(1 * time.Hour)
			var igeoms = ReadInputPolylines("data/lines.txt")
			g.Assert(len(igeoms)).Equal(4)
		})
	})
}
