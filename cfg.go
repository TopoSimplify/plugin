package main

import (
	"encoding/json"
	"github.com/TopoSimplify/plugin/opts"
	"log"
)

//ArgObj struct
type ArgObj struct {
	Input                  string  `json:"input"`
	Output                 string  `json:"output"`
	Constraints            string  `json:"constraints"`
	SimplificationType     string  `json:"simplification_type"`
	Threshold              float64 `json:"threshold"`
	MinDist                float64 `json:"minimum_distance"`
	NonPlanarDisplacement  float64 `json:"non_planar_displacement"`
	IsFeatureClass         bool    `json:"is_feature_class"`
	PlanarSelf             bool    `json:"planar_self"`
	NonPlanarSelf          bool    `json:"non_planar_self"`
	AvoidNewSelfIntersects bool    `json:"avoid_new_self_intersects"`
	GeomRelation           bool    `json:"geometric_relation"`
	DistRelation           bool    `json:"distance_relation"`
	SideRelation           bool    `json:"homotopy_relation"`
}

func (opt ArgObj) String() string {
	var cfgbytes, err = json.Marshal(opt)
	if err != nil {
		panic(err)
	}
	return string(cfgbytes)
}

func optsFromCfg(obj ArgObj) opts.Opts {
	return opts.Opts{
		Threshold:              obj.Threshold,
		MinDist:                obj.MinDist,
		NonPlanarDisplacement:  obj.NonPlanarDisplacement,
		PlanarSelf:             obj.PlanarSelf,
		NonPlanarSelf:          obj.NonPlanarSelf,
		AvoidNewSelfIntersects: obj.AvoidNewSelfIntersects,
		GeomRelation:           obj.GeomRelation,
		DistRelation:           obj.DistRelation,
		DirRelation:            obj.SideRelation,
	}
}

func parseInput(arg string) ArgObj {
	var config = ArgObj{}
	var jsonInput = decode64(arg)
	if err := json.Unmarshal([]byte(jsonInput), &config); err != nil {
		log.Println("invalid input:")
		log.Fatalln(err)
	}
	return config
}
