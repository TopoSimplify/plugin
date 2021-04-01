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
Open a terminal (command prompt) from the directory containing an executable 
(constdp[.exe] for 64bit, constdp_32bit[.exe] for 32bit systems). 
Simplification options are made available through  the use of a [TOML][0] file (config.toml).
Execute `constdp` with the following command :

```bash
./constdp -c ./config.toml 
```

If a `-c` option is not provided at the terminal e.g. `./constdp`, it assumes `./config.toml` 
as the default configuration file. Change `config.toml` to configure your simplification. 

#### config file 

```toml
# input file is required
Input                  = "/path/to/input.[wkt]" 
# output is optional, defaults to ./out.txt
Output                 = "" 
# this is optional
Constraints            = "/path/to/file.[wkt]" 
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
