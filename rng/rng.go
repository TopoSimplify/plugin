package rng

import (
	"fmt"
)

//Rng
type Rng struct {
	I int
	J int
}

//New Rng
func Range(i, j int) Rng {
	return Rng{i, j}
}

//Stringer interface
func (o Rng) String() string {
	return fmt.Sprintf("Rng(i=%v, j=%v)", o.I, o.J)
}

//compare equality of two ranges
func (o Rng) Equals(r Rng) bool {
	return (o.I == r.I) && (o.J == r.J)
}

//Size
func (o Rng) Size() int {
	return o.J - o.I
}
