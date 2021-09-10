
# Introduction 

For many applications in spatial and spatio-temporal GIS, it is beneficial to collect data once the highest possible 
resolution as a master database. Simplification can be used as a tool to derive data at a coarser resolution from the 
master database for other applications.  Simplification can also be used  as a pre-processing tool before data mining, 
visualization, data transmission, and data exploration.

Line simplification is the process of deforming the shape of a geometry by removing excessive detail. There are two methods
 - Line smoothing (curve fitting) 
 - Vertex reduction

Either method, based on some criteria, may produce a simplification that has much less detail compared
to the original polyline. If topologically constrained, th simplification will result in an output that maintains some
topological fidelity with the input polyline.




