package main

import (
	"fmt"
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
	var cfg = parseInput(strings.TrimSpace(args[0]))
	var options = optsFromCfg(cfg)
	var constraints []geom.Geometry
	var simpleCoords []geom.Coords

	//config simplification type
	cfg.SimplificationType = strings.ToLower(strings.TrimSpace(cfg.SimplificationType))
	var offsetFn = offsetDictionary[cfg.SimplificationType]
	if offsetFn == nil {
		log.Println(`Supported Simplification Types : "DP" or "SED", Fix input`)
		os.Exit(1)
	}

	var err error
	var polyCoords []geom.Coords
	if isShapeFile(cfg.Input) {
		polyCoords, err = readWKTInput(cfg.Input)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
			os.Exit(1)
		}
	} else {
		panic("unknown file type, expects a shapefile: /path/to/name.shp")
	}

	// config output
	cfg.Output = strings.TrimSpace(cfg.Output)
	if cfg.Output == "" {
		panic("output path shapefile (*.shp) path required !")
	}

	// read constraints
	cfg.Constraints[0] = strings.TrimSpace(cfg.Constraints[0])

	if cfg.Constraints[0] != "" && isShapeFile(cfg.Constraints[0]) {
		constraints, err = readConstraints(cfg.Constraints[0])
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Constraints, err))
			os.Exit(1)
		}
	}

	// simplify
	log.Println("starting simplification ")
	var t0 = time.Now()
	if cfg.IsFeatureClass {
		simpleCoords = simplifyFeatureClass(polyCoords, &options, constraints, offsetFn)
	} else {
		simpleCoords = simplifyInstances(polyCoords, &options, constraints, offsetFn)
	}
	var t1 = time.Now()
	log.Println("done simplification ")
	log.Println(fmt.Sprintf("elapsed time: %v seconds", math.Round(t1.Sub(t0).Seconds(), 6)))

	var saved bool
	//Save output
	if isShapeFile(cfg.Input) {
		switch cfg.SimplificationType {
		case "dp":
			err = writeCoords(cfg.Output, simpleCoords, geom.WriteWKT)
		case "sed":
			err = writeCoords(cfg.Output, simpleCoords, geom.WriteWKT3D)
		}
		if err != nil {
			panic(err)
		}
		saved = true
	} else {
		panic("unknown file type, expects output as shapefile: /path/to/name.shp")
	}
	if saved {
		log.Println("simplification save to file :", cfg.Output)
	}
}

func optsFromCfg(cfg Cfg) opts.Opts {
	return opts.Opts{
		Threshold:              cfg.Threshold,
		MinDist:                cfg.MinDist,
		RelaxDist:              cfg.RelaxDist,
		PlanarSelf:             cfg.PlanarSelf,
		AvoidNewSelfIntersects: cfg.AvoidNewSelfIntersects,
		GeomRelation:           cfg.GeomRelation,
		DistRelation:           cfg.DistRelation,
		DirRelation:            cfg.SideRelation,
	}
}

func isShapeFile(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".shp"
}
