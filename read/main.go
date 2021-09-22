package main

import (
	"github.com/intdxdt/fileutil"
	geojson "github.com/paulmach/go.geojson"
	"strings"
)

func main() {
	readJsonFile("input.json")
	//fc2 := geojson.NewFeature()

}

func parseFeatures(inputs []string){
	for _, fjson  := range inputs {
		feat, err := geojson.UnmarshalFeature([]byte(fjson))
		checkError(err)
		println(feat.Geometry.IsLineString() || feat.Geometry.IsMultiLineString())
	}
}



func readJsonFile(file string) {
	var data, err = fileutil.ReadAllOfFile(file)
	checkError(err)
	var tokens  = strings.Split(strings.TrimSpace(data), "\n")
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	parseFeatures(tokens)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
