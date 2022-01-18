Table of contents
=================
<!--ts-->
   * [Installation](#introduction)
   * [Usage](#non_planar_self)
      * [STDIN](#stdin)
      * [Local files](#local-files)
      * [Remote files](#remote-files)
      * [Multiple files](#multiple-files)
      * [Combo](#combo)
      * [Auto insert and update TOC](#auto-insert-and-update-toc)
      * [GitHub token](#github-token)
      * [TOC generation with Github Actions](#toc-generation-with-github-actions)
   * [Tests](#tests)
   * [Dependency](#dependency)
   * [Docker](#docker)
     * [Local](#local)
     * [Public](#public)
<!--te-->


# Introduction 

For many applications in spatial and spatio-temporal GIS, it is beneficial 
to collect data once the highest possible resolution as a master database. 
Simplification can be used as a tool to derive data at a coarser resolution 
from the master database for other applications.  Simplification can also be 
used  as a pre-processing tool before data mining, visualization, data 
transmission, and data exploration.

Line simplification is the process of deforming the shape of a geometry by 
removing excessive detail. There are two methods:
 - Line smoothing (curve fitting) 
 - Vertex reduction

Either method, based on some criteria, may produce a simplification that has 
much less detail compared to the original polyline. If topologically constrained, 
the simplification will result in an output that maintains some topological 
consistency with the original input.

Unconstrained simplification can lead to various topological errors in the output 
geometry. In circumstances where the input geometry has spatial relationships with 
other features in the same layer or other layers, the spatial context under which 
the linear feature to be simplified needs consideration. 
Out of context simplification can lead to change in contextual representation and 
meaning - simplified geometry may be self topologically consistent but spatially 
invalid relative to other features that constrain its shape. 

The goal of this plugin is to support contextual simplification of linear features
and spatio-temporal trajectories. 

## Geometry Types  



## Consistent Line Simplification in the Context of Planar Constraints

Constrained simplification of arbitrary polylines in the context of arbitrary planar geometries.

### How to build

Install the latest version of [Go](https://golang.org/dl/)
Clone the repository, open a terminal/command prompt, change director (cd) to the cloned directory. Enter the command:

```shell
go build -o simplify.exe
```

### How to use

To perform a contextual simplification, a simplification configuration is need. A configuration is a
`JSON` string containing multiple simplification options. `simplify`[.exe] requires a configuration json in a text file
or a base64 encoded string passed at the command line.

Use the command line option `-b` for

```bash
./simplify.exe -b ewogICAgICAgICAgImlucHV0IiAgICAgICAgICAgICAgICAgICAgIDogImRhdGEvZm...
```

Use the command line option `-f` for configuration in a file

```bash
./simplify.exe -f "c:/path/to/config.json"
```

### Options:

Simplification options are a set of key value pairs

#### Sample JSON configuration

```json
{
  "input": "data/feature_class.json",
  "output": "output/out_feat_class.json",
  "constraints": "data/feature_class_const.json",
  "simplification_type": "DP",
  "threshold": 50.0,
  "minimum_distance": 20.0,
  "relax_distance": 10.0,
  "is_feature_class": false,
  "planar_self": true,
  "non_planar_self": true,
  "avoid_new_self_intersects": true,
  "geometric_relation": true,
  "distance_relation": true,
  "homotopy_relation": true
}
```

#### input

GeoJSON file with new line delimited (each line is linestring geojson feature) of JSON features (`LineString`
or `MultiLineString`)

```text
"input" : "data/input.json"
```

#### output

Path to simplification output(GeoJSON) as newline delimited simplification of `input`

```text 
"output" : "output/output.json"
```

#### constraints

GeoJSON file with new line delimited of JSON `Point`/`MultiPoint`, `LineString`/`MultiLineString` or `Polygon`
/`MultiPolygon`
geometries.

```text
"constraints" : "data/constraints.json" 
```

#### simplification_type

Type of simplification: `"DP"` or `"SED"`

```text
"simplification_type" : "DP"
```

#### threshold

Simplification distance threshold - in same units as input planar coordinates

```text
"threshold" : 0.0
```

#### minimum_distance

Minimum distance from planar contraints - provide a value if `"distance_relation": true`

```text
"minimum_distance" : 0.0
```

#### relax_distance

Relax distance for non-planar intersections - provide value if `NonPlanarSelf = true`

```text
"relax_distance" : 0.0
```

#### is_feature_class

Are polylines independent or a feature class ? if `false` planar and non-planar intersections between polylines are not
observed. If set to `true` the relations between each feature in the class of linestrings are preserved based on options
provided.

```text
"is_feature_class" : false
```

#### planar_self

Observe planar self-intersection - preserves planar intersection (vertex with degree greater than 2).
If `is_feature_class`
preserves planar intersections between features of a feature class.

```text
"planar_self" : false
```

#### non_planar_self

Observe non-planar self-intersection - preserves non-planar intersection (overlaps between lines that do not introduce
an intersection). If `is_feature_class` preserves non-planar intersections between features of a feature class based on
a relaxation distance.

```text
"non_planar_self" : false
```

#### avoid_new_self_intersects

Avoid introducing new self-intersections as a result of simplification algorithm.

```text
"avoid_new_self_intersects" : false
```

#### geometric_relation

Observe geometric relation (intersect / disjoint) to planar objects serving as constraints.

```text
"geometric_relation" : false
```

#### distance_relation

Observe distance relation (minimum distance) to planar objects serving as constraints.

```text
distance_relation : false
```

#### homotopy_relation

Observe homotopic (sidedness) relation to planar objects serving as constraints.

```text
"homotopy_relation" : false
```

