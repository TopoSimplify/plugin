package main

import (
	"encoding/json"
	"log"
)

//Cfg struct
type Cfg struct {
	Input                  string   `json:"input"`
	Output                 string   `json:"output"`
	Constraints            []string `json:"constraints"`
	SimplificationType     string   `json:"simplification_type"`
	Threshold              float64  `json:"threshold"`
	MinDist                float64  `json:"minimum_distance"`
	RelaxDist              float64  `json:"relax_distance"`
	IsFeatureClass         bool     `json:"is_feature_class"`
	PlanarSelf             bool     `json:"planar_self"`
	NonPlanarSelf          bool     `json:"non_planar_self"`
	AvoidNewSelfIntersects bool     `json:"avoid_new_self_intersects"`
	GeomRelation           bool     `json:"geometric_relation"`
	DistRelation           bool     `json:"distance_relation"`
	SideRelation           bool     `json:"homotopy_relation"`
}

func (opt Cfg) String() string {
	var cfgbytes, err = json.Marshal(opt)
	if err != nil {
		panic(err)
	}
	return string(cfgbytes)
}

func parseInput(arg string) Cfg {
	var config = Cfg{}
	var jsonInput = decode64(arg)
	if err := json.Unmarshal([]byte(jsonInput), &config); err != nil {
		log.Println("invalid input:")
		log.Fatalln(err)
	}

	//if !fileutil.IsFile(config.Input) {
	//	log.Println("input file not found")
	//	usageHelp()
	//	os.Exit(13)
	//}
	//
	//if strings.TrimSpace(config.Constraints) != "" && !fileutil.IsFile(config.Constraints) {
	//	log.Println("constraints file not found")
	//	usageHelp()
	//	os.Exit(13)
	//}

	return config
}
