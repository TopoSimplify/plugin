package state

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestState(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("State", func() {
		g.It("test clean and dirty", func() {
			g.Timeout(1 * time.Hour)
			var s State
			s.MarkClean()
			g.Assert(s.IsClean()).IsTrue()
			g.Assert(s.IsDirty()).IsFalse()
			s.MarkDirty()
			g.Assert(s.IsDirty()).IsTrue()
			g.Assert(s.IsClean()).IsFalse()
		})
	})
}
