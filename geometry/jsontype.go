package geometry

type JSONType struct {
	Type string `json:"type"`
}

//isGeometryType
func (o *JSONType) isGeometryType() bool {
	return o.Type == TypePoint || o.Type == TypeLineString || o.Type == TypePolygon ||
		o.Type == TypeMultiPoint || o.Type == TypeMultiLineString || o.Type == TypeMultiPolygon
}

//isFeatureType
func (o *JSONType) isFeatureType() bool {
	return o.Type == TypeFeature
}

//isFeatureCollectionType
func (o *JSONType) isFeatureCollectionType() bool {
	return o.Type == TypeFeatureCollection
}
