package main

import (
	"fmt"
	"github.com/TopoSimplify/plugin/geometry"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/math"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var wktPoint = []byte("point")
var wktPolygon = []byte("polygon")
var wktLinestring = []byte("linestring")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var args = os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("input base64 string not provided !")
	}

	var argObj = parseInput(strings.TrimSpace(args[0]))
	var options = optsFromCfg(argObj)

	//var polyline = geometry.ReadInputPolylines("data/input.json")
	//var constraints = geometry.ReadInputConstraints("data/constraints.json")

	//config simplification type
	argObj.SimplificationType = strings.ToLower(strings.TrimSpace(argObj.SimplificationType))
	var offsetFn = offsetDictionary[argObj.SimplificationType]
	if offsetFn == nil {
		log.Println(`Supported Simplification Types : "DP" or "SED", Fix input`)
		os.Exit(1)
	}

	var polylines = make([]*geometry.Polyline, 0)
	//var outputPolylines = make([]geometry.Polyline, 0)

	//lines
	argObj.Input = strings.TrimSpace(argObj.Input)
	if fileutil.IsFile(argObj.Input) {
		polylines = geometry.ReadInputPolylines(argObj.Input)
	} else {
		panic("input file not provided")
	}

	//output
	argObj.Output = strings.TrimSpace(argObj.Output)
	if argObj.Output == "" {
		panic("output filepath required !")
	}

	//constraints
	var constraints = make([]geometry.IGeometry, 0)
	argObj.Constraints = strings.TrimSpace(argObj.Constraints)
	if fileutil.IsFile(argObj.Constraints) {
		constraints = geometry.ReadInputConstraints(argObj.Constraints)
	} else {
		if argObj.Constraints != "" {
			panic("constraint file not found")
		}
	}

	var t0 = time.Now()
	if argObj.IsFeatureClass {
		//outputPolylines = simplifyFeatureClass(polylines, &options, constraints, offsetFn)
		simplifyFeatureClass(polylines, &options, constraints, offsetFn)
	} else {
		//outputPolylines = simplifyInstances(polylines, &options, constraints, offsetFn)
		simplifyInstances(polylines, &options, constraints, offsetFn)
	}
	var t1 = time.Now()

	var output = groupPlnsAsGeoJSONS(polylines)
	var err = fileutil.SaveText(argObj.Output, strings.Join(output, "\n"))
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("elapsed time: %v seconds", math.Round(t1.Sub(t0).Seconds(), 6)))
}
