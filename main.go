package main

import (
	"flag"
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
var argJSONFile string
var argBase64String string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&argBase64String, "b", "-", "base64 string")
	flag.StringVar(&argJSONFile, "f", "-", "json config file")

	flag.Parse()

	if argBase64String == "-" && argJSONFile == "-" {
		log.Println("\nHow to use:")
		flag.PrintDefaults()
		fmt.Println("Invalid arguments!")
		os.Exit(13)
	}

	if argBase64String != "-" && argJSONFile != "-" {
		log.Println("")
		log.Println("\nHow to use:")
		flag.PrintDefaults()
		fmt.Println("Ambigous arguments! Expects only one flag: -b or -f")
		os.Exit(13)
	}

	if argJSONFile != "-" {
		var jsonConfig, err = fileutil.ReadAllOfFile(argJSONFile)
		if err != nil {
			panic(err)
		}
		argBase64String = encode64(strings.TrimSpace(jsonConfig))
	}

	var argObj = parseInput(strings.TrimSpace(argBase64String))
	var options = optsFromCfg(argObj)

	//config simplification type
	argObj.SimplificationType = strings.ToLower(strings.TrimSpace(argObj.SimplificationType))
	var offsetFn = offsetDictionary[argObj.SimplificationType]
	if offsetFn == nil {
		log.Println(`Supported Simplification Types : "DP" or "SED", Fix input`)
		os.Exit(13)
	}

	//lines
	var polylines = make([]*geometry.Polyline, 0)
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
		simplifyFeatureClass(polylines, &options, constraints, offsetFn)
	} else {
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
