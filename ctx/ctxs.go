package ctx

import (
	"sort"
)

type ContextGeometries struct {
	list []*ContextGeometry
}

func NewContexts(geometries ...*ContextGeometry) *ContextGeometries {
	return &ContextGeometries{list: geometries}
}

func NewContextsFromObjects(objects []interface{}) *ContextGeometries {
	var n = len(objects)
	var contexts = &ContextGeometries{list: make([]*ContextGeometry, 0, n)}
	for i := range objects {
		contexts.list = append(contexts.list, objects[i].(*ContextGeometry))
	}
	return contexts
}

func (self *ContextGeometries) Len() int {
	return len(self.list)
}

func (self *ContextGeometries) Swap(i, j int) {
	self.list[i], self.list[j] = self.list[j], self.list[i]
}

func (self *ContextGeometries) Less(i, j int) bool {
	return self.list[i].I < self.list[j].I
}

func (self *ContextGeometries) Push(o *ContextGeometry) *ContextGeometries {
	self.list = append(self.list, o)
	return self
}

func (self *ContextGeometries) Extend(o []*ContextGeometry) *ContextGeometries {
	self.list = append(self.list, o...)
	return self
}

func (self *ContextGeometries) Sort() *ContextGeometries {
	sort.Sort(self)
	return self
}

func (self *ContextGeometries) DataView() []*ContextGeometry {
	return self.list
}

func (self *ContextGeometries) SetData(o []*ContextGeometry) *ContextGeometries{
	self.list = o
	return self
}
