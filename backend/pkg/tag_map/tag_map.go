package tagmap

import "FilesWithTag/pkg/set"

type TagMap struct {
	Map     map[string]*set.Set
	Inverse map[string]*set.Set
}

func NewTagMap() *TagMap {
	return &TagMap{
		Map:     make(map[string]*set.Set),
		Inverse: make(map[string]*set.Set),
	}
}

func (t *TagMap) Add(key string, value string) {
	if _, ok := t.Map[key]; !ok {
		t.Map[key] = set.NewSet()
	}
	t.Map[key].Append(value)

	if _, ok := t.Inverse[value]; !ok {
		t.Inverse[value] = set.NewSet()
	}
	t.Inverse[value].Append(key)
}

func (t *TagMap) Remove(key string, value string) {
	t.Map[key].Remove(value)
	t.Inverse[value].Remove(key)
}

func (t *TagMap) Get(key string) []string {
	return t.Map[key].ToSlice()
}

func (t *TagMap) GetInverse(key string) []string {
	return t.Inverse[key].ToSlice()
}

func (t *TagMap) Contains(key string, value string) bool {
	return t.Map[key].Contains(value)
}

func (t *TagMap) ContainsInverse(key string, value string) bool {
	return t.Inverse[value].Contains(key)
}
