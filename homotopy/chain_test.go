package homotopy

import (
	"time"
	"testing"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/ctx"
	"github.com/intdxdt/geom"
)

func reverse(coordinates geom.Coords) geom.Coords {
	for i, j := 0, len(coordinates.Idxs)-1; i < j; i, j = i+1, j-1 {
		coordinates.Idxs[i], coordinates.Idxs[j] = coordinates.Idxs[j], coordinates.Idxs[i]
	}
	return coordinates
}


func contextGeoms(wkts []string) *ctx.ContextGeometries {
	var contexts = ctx.NewContexts()
	for _, wkt := range wkts {
		contexts.Push(ctx.New(geom.ReadGeometry(wkt), 0, -1))
	}
	return contexts
}

func TestChain(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("chain", func() {
		g.It("should test chain contruction", func() {
			g.Timeout(1 * time.Hour)
			var wkt = "LINESTRING ( 600 0, 600 100, 500 100, 500 -100, 400 100, 600 200, 900 300, 1200 200, 1400 100, 1400 1.7158013051015217, 1400 -100, 1336.2430712729156 -100, 1332.8103356508254 97.83239872362749, 1200 100, 1200 -100, 1236.6937382322994 -87.53532486924402, 1226.3955313660288 72.08688155795089, 1291.6175081857427 66.93777812481557, 1278.0572495001297 -101.81210774059059, 1000 -300, 739.3977399886022 -182.83688839549262, 739.3977399886022 117.16311160450738, 839.3977399886022 117.16311160450738, 841.9291416919249 -102.98263516864998, 883.251804604375 -40.347537079926326, 884.4792054022216 117.8277156849893, 939.3977399886022 117.16311160450738, 939.762106921496 -102.98263516864998, 1039.3114399621122 -102.98263516864998, 1099.384313348691 -48.05886521520657, 1099.6824840778606 5.139171483985346 )"
			var coords = loads(wkt)
			var ch = NewChain(coords)

			g.Assert(ch.size).Equal(coords.Len())

			var link = ch.link
			var indices, expect []int

			for link != nil {
				indices = append(indices, link.i)
				link = link.next
			}

			for i := 0; i < coords.Len(); i++ {
				expect = append(expect, i)
			}

			g.Assert(indices).Equal(expect)
		})

	})
}
