package state

type State bool

func (s *State) MarkClean() {
	*s = true
}

func (s *State) MarkDirty() {
	*s =  false
}

func (s State) IsClean() bool {
	return s == true
}

func (s State) IsDirty() bool {
	return s == false
}
