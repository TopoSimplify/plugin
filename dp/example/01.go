package main

import (
	"fmt"
	"github.com/TopoSimplify/plugin/dp"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/TopoSimplify/plugin/offset"
	"github.com/TopoSimplify/plugin/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

var data = []geom.Point{{0.5, 1.0}, {1.5, 2.0}, {2.5, 1.5}, {3.5, 2.5}, {4.0, 1.5}, {3.0, 1.0}}

func main() {
	//var gstr = 'LINESTRING (520.3891360357894 542.3912033070129, 506.3024618690045 551.4232473315985, 499.8456492240652 555.3948968460392, 492.961552805167 552.5004635914114, 489.3155900796462 547.0315195031302, 494.7910190818659 540.6453203655232, 503.2430235819369 542.0539877822016, 506.3024618690045 551.4232473315985, 505.72509579166825 560.3502151427206, 505.2252456091915 568.0786679640912)'
	//var g = geom.readwkt(gstr)
	var options = &opts.Opts{Threshold: 0.6}
	var coords = geom.Coordinates(data)
	var pln = geometry.CreatePolyline("0", coords, "")
	var tree = dp.New(0, pln, options, offset.MaxOffset)
	var id = iter.NewIgen()

	var o = geom.NewLineString(tree.Coordinates())
	fmt.Println(o.WKT())
	tree.Simplify(id)
	fmt.Println(tree.Simple())
	//o = geom.NewLineString()
	//fmt.Println (o.WKT())

}
