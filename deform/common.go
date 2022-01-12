package deform

import (
	"github.com/TopoSimplify/plugin/lnr"
)

func isSame(a, b lnr.Linegen) bool {
	return a.Id() == b.Id()
}
