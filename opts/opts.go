package opts

import "encoding/json"

//Opts - input options
type Opts struct {
	Threshold              float64 `json:"threshold"`
	MinDist                float64 `json:"mindist"`
	RelaxDist              float64 `json:"relaxdist"`
	PlanarSelf             bool    `json:"planarself"`
	NonPlanarSelf          bool    `json:"nonplanarself"`
	AvoidNewSelfIntersects bool    `json:"avoidself"`
	GeomRelation           bool    `json:"geomrelate"`
	DistRelation           bool    `json:"distrelate"`
	DirRelation            bool    `json:"dirrelate"`
}

func (opt Opts) String() string {
	var bytes, err = json.Marshal(opt)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
