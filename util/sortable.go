package util

type SortableMap struct {
	Key   interface{}
	Value int64
}

type SortableMapList []SortableMap

func (p SortableMapList) Len() int           { return len(p) }
func (p SortableMapList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p SortableMapList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
