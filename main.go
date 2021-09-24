package main

import (
	"fmt"
	"github.com/TopoSimplify/geometry"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"io"
	"log"
	"os"
	"path/filepath"
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
	var simpleCoords []geom.Coords

	var polyline = geometry.ReadInputPolylines("data/input.json")
	var constraints = geometry.ReadInputConstraints("data/constraints.json")
	fmt.Println(polyline)
	fmt.Println(constraints)

	//config simplification type
	argObj.SimplificationType = strings.ToLower(strings.TrimSpace(argObj.SimplificationType))
	var offsetFn = offsetDictionary[argObj.SimplificationType]
	if offsetFn == nil {
		log.Println(`Supported Simplification Types : "DP" or "SED", Fix input`)
		os.Exit(1)
	}

	var err error
	var polyCoords []geom.Coords
	if isShapeFile(argObj.Input) {
		polyCoords, err = readWKTInput(argObj.Input)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", argObj.Input, err))
			os.Exit(1)
		}
	} else {
		panic("unknown file type, expects a shapefile: /path/to/name.shp")
	}

	// config output
	argObj.Output = strings.TrimSpace(argObj.Output)
	if argObj.Output == "" {
		panic("output path shapefile (*.shp) path required !")
	}

	// read constraints
	argObj.Constraints[0] = strings.TrimSpace(argObj.Constraints[0])

	if argObj.Constraints[0] != "" && isShapeFile(argObj.Constraints[0]) {
		//constraints, err = readConstraints(argObj.Constraints[0])
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", argObj.Constraints, err))
			os.Exit(1)
		}
	}

	// simplify
	log.Println("starting simplification ")
	var t0 = time.Now()
	if argObj.IsFeatureClass {
		simpleCoords = simplifyFeatureClass(polyCoords, &options, constraints, offsetFn)
	} else {
		simpleCoords = simplifyInstances(polyCoords, &options, constraints, offsetFn)
	}
	var t1 = time.Now()
	log.Println("done simplification ")
	log.Println(fmt.Sprintf("elapsed time: %v seconds", math.Round(t1.Sub(t0).Seconds(), 6)))

	var saved bool
	//Save output
	if isShapeFile(argObj.Input) {
		switch argObj.SimplificationType {
		case "dp":
			err = writeCoords(argObj.Output, simpleCoords, geom.WriteWKT)
		case "sed":
			err = writeCoords(argObj.Output, simpleCoords, geom.WriteWKT3D)
		}
		if err != nil {
			panic(err)
		}
		saved = true
	} else {
		panic("unknown file type, expects output as shapefile: /path/to/name.shp")
	}
	if saved {
		log.Println("simplification save to file :", argObj.Output)
	}
}

func optsFromCfg(obj ArgObj) opts.Opts {
	return opts.Opts{
		Threshold:              obj.Threshold,
		MinDist:                obj.MinDist,
		RelaxDist:              obj.RelaxDist,
		PlanarSelf:             obj.PlanarSelf,
		AvoidNewSelfIntersects: obj.AvoidNewSelfIntersects,
		GeomRelation:           obj.GeomRelation,
		DistRelation:           obj.DistRelation,
		DirRelation:            obj.SideRelation,
	}
}

func isShapeFile(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".shp"
}
