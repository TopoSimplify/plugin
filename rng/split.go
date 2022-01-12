package rng

import "github.com/intdxdt/iter"

//Split Rng at indices
func (o Rng) Split(idxs []int) []Rng {
	idxs = iter.SortedIntsSet(idxs)
	var i, j = o.I, o.J
	var sub = make([]Rng, 0)
	for _, idx := range idxs {
		if i < idx && idx < j {
			s := i
			if len(sub) > 0 {
				s = sub[len(sub)-1].J
			}
			sub = append(sub, Range(s, idx))
		}
	}
	//close rest of split
	if len(sub) > 0 {
		e := sub[len(sub)-1].J
		if e < j {
			sub = append(sub, Range(e, j))
		}
	}
	return sub
}

