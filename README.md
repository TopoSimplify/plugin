## Topologically Consistent Line Simplification in the Context of Planar Constraints
Constrained simplification of arbitrary polylines in the context of arbitrary planar geometries.

### How to build
Install the latest version of [Go](https://golang.org/dl/)
Clone the repository, open a terminal/command prompt, change director (cd) to the path of `plugin` directory.
Enter the command:
```shell
go build
```
or with an optional executable name:
```shell
go build -o plugin.exe
```

### how to use 
Open a command execute plugin[.exe] passing in a json string of configuration.
Make sure to properly escape input JSON. For example:
```bash
./plugin.exe "{\"a\": \"C:\\path\\to\\input.shp\", \"b\": 299.434}"
```

#### JSON input fields

```json
{
  "input": "/path/to/input.[shp]",
  "output": "/path/to/output.shp",
  "constraints": "/path/to/file.[shp]",
  "simplificationType": "DP",
  "threshold": 0.0,
  "minDist": 0.0,
  "relaxDist": 0.0,
  "isFeatureClass": false,
  "planarSelf": false,
  "nonPlanarSelf": false,
  "avoid_new_self_intersects": false,
  "geom_relation": false,
  "dist_relation": false,
  "side_relation" : false
}
```
Json input field descriptions:
```toml
# input file is required
Input                  = "/path/to/input.[shp]" 
# output is optional, defaults to ./out.txt
Output                 = "/path/to/output.shp" 
# this is optional
Constraints            = "/path/to/file.[shp]" 
# type of simplification, options : DP, SED
SimplificationType     = "DP"
# simplification threshold (in metric units as input geometric coordinates) 
Threshold              = 0.0
# minimum distance from planar contraints - provide value if `DistRelation = true`
MinDist                = 0.0
# relax distance for non-planar intersections - provide value if `NonPlanarSelf = true`
RelaxDist              = 0.0
# are polylines independent or a feature class ?
# if false planar and non-planar intersections between polylines are not observed
IsFeatureClass         = false
# observe planar self-intersection
PlanarSelf             = false
# observe non-planar self-intersection
NonPlanarSelf          = false
# avoid introducing new self-intersections as a result of simplification
AvoidNewSelfIntersects = false
# observe geometric relation (intersect / disjoint) to planar objects serving as constraints
GeomRelation           = false
# observe distance relation (minimum distance) to planar objects serving as constraints
DistRelation           = false
# observe homotopic (sidedness) relation to planar objects serving as constraints
SideRelation           = false
```
