package rng

//are two ranges contiguous
func (o Rng) Contiguous(r Rng) bool {
	return (o.I < r.J && o.J == r.I) || (r.I < o.J && r.J == o.I)
}
