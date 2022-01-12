package rng

import "github.com/intdxdt/iter"

//Stride
func (o Rng) Stride(step ...int) []int {
	var s = 1
	if len(step) > 0 {
		s = step[0]
	}
	return iter.NewGenerator(o.I, o.J+1, s).Values()
}

//Exclusive stride
func (o *Rng) ExclusiveStride(step ...int) []int {
	var s = 1
	if len(step) > 0 {
		s = step[0]
	}
	return iter.NewGenerator(o.I+1, o.J, s).Values()
}

