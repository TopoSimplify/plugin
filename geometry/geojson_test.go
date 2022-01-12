package geometry

import (
	"encoding/json"
	"github.com/franela/goblin"
	"strings"
	"testing"
	"time"
)

func TestGeoJSON_IO(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GeoJSON - IO", func() {
		g.It("geojson IO", func() {
			g.Timeout(1 * time.Minute)
			var obj JSONType
			var dat = `{ "type": "Point", "coordinates": [30.0, 10.0] }`

			var err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsTrue()

			dat = `{ "type": "FeatureCollection", "features": [ { "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }, { "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }, { "type": "Feature", "geometry": { "type": "Polygon", "coordinates": [ [ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ] ] }, "properties": { "prop0": "value0", "prop1": {"this": "that"} } } ] }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsTrue()

			dat = `{"type": "Feature", "geometry": {"type": "Polygon", "coordinates": [[[100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]]]}, "properties": {"prop0": "value0", "prop1": {"this": "that"}}}`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

			dat = `{ "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

			dat = `{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }`
			err = json.Unmarshal([]byte(dat), &obj)
			g.Assert(err).IsNil()
			g.Assert(obj.isGeometryType()).IsFalse()
			g.Assert(obj.isFeatureCollectionType()).IsFalse()
			g.Assert(obj.isFeatureType()).IsTrue()

		})
	})
}

func TestGeoJSON_IO_Constraints(t *testing.T) {
	g := goblin.Goblin(t)
	var textFile = `
{ "type": "FeatureCollection", "features": [ { "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }, { "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }, { "type": "Feature", "geometry": { "type": "Polygon", "coordinates": [ [ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ] ] }, "properties": { "prop0": "value0", "prop1": {"this": "that"} } } ] }
{ "type": "Point", "coordinates": [30.0, 10.0] }
{ "type": "LineString", "coordinates": [ [30.0, 10.0], [10.0, 30.0], [40.0, 40.0] ] }
{ "type": "Polygon", "coordinates": [ [[30.0, 10.0], [40.0, 40.0], [20.0, 40.0], [10.0, 20.0], [30.0, 10.0]] ] }
{ "type": "Polygon", "coordinates": [ [[35.0, 10.0], [45.0, 45.0], [15.0, 40.0], [10.0, 20.0], [35.0, 10.0]], [[20.0, 30.0], [35.0, 35.0], [30.0, 20.0], [20.0, 30.0]] ] }
{ "type": "MultiPoint", "coordinates": [ [10.0, 40.0], [40.0, 30.0], [20.0, 20.0], [30.0, 10.0] ] }
{ "type": "MultiLineString", "coordinates": [ [[10.0, 10.0], [20.0, 20.0], [10.0, 40.0]], [[40.0, 40.0], [30.0, 30.0], [40.0, 20.0], [30.0, 10.0]] ] }
{ "type": "MultiPolygon", "coordinates": [ [ [[30.0, 20.0], [45.0, 40.0], [10.0, 40.0], [30.0, 20.0]] ], [ [[15.0, 5.0], [40.0, 10.0], [10.0, 20.0], [5.0, 10.0], [15.0, 5.0]] ] ] }
{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [102.0, 0.5] }, "properties": { "prop0": "value0" } }
{ "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }
{ "type": "Feature", "geometry": { "type": "Polygon", "coordinates": [ [ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ] ] }, "properties": { "prop0": "value0", "prop1": { "this": "that" } } }
`
	g.Describe("GeoJSON - IO", func() {
		g.It("geojson IO files from file", func() {
			g.Timeout(1 * time.Hour)
			textFile = strings.TrimSpace(textFile)
			var inputs = strings.Split(textFile, "\n")
			var igeoms = parseConstraintFeatures(inputs)
			g.Assert(len(igeoms)).Equal(len(igeoms))
		})
	})
}

func TestGeoJSON_IO_LineStrings(t *testing.T) {
	g := goblin.Goblin(t)
	var textFile = `
{ "type": "LineString", "coordinates": [ [30.0, 10.0], [10.0, 30.0], [40.0, 40.0] ] }
{ "type": "MultiLineString", "coordinates": [ [[10.0, 10.0], [20.0, 20.0], [10.0, 40.0]], [[40.0, 40.0], [30.0, 30.0], [40.0, 20.0], [30.0, 10.0]] ] }
{ "type": "Feature", "geometry": { "type": "LineString", "coordinates": [ [102.0, 0.0], [103.0, 1.0], [104.0, 0.0], [105.0, 1.0] ] }, "properties": { "prop0": "value0", "prop1": 0.0 } }
`
	g.Describe("GeoJSON - IO", func() {
		g.It("geojson IO files from file", func() {
			g.Timeout(1 * time.Hour)
			textFile = strings.TrimSpace(textFile)
			var inputs = strings.Split(textFile, "\n")
			var igeoms = parseInputLinearFeatures(inputs)
			g.Assert(len(igeoms)).Equal(4)
		})
	})
}
