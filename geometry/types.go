package geometry

type ILinear interface {
	LinearType() string
}

type LinearGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

func (ln *LinearGeometry) LinearType() string {
	return ln.Type
}

type MultiLinearGeometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

func (ln *MultiLinearGeometry) LinearType() string {
	return ln.Type
}

type LineStringFeature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   ILinear                `json:"geometry"`
}
