package rng

//As Array
func (o Rng) AsArray() [2]int {
	return [2]int{o.I, o.J}
}

//As Slice
func (o Rng) AsSlice() []int {
	var a = o.AsArray()
	return a[:]
}

