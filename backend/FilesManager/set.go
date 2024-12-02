package filesmanager

import "strings"

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[string]struct{}, 0)}
}

func (s *Set) Append(value string) {
	s.data[value] = struct{}{}
}

func (s *Set) Contains(value string) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set) Remove(value string) {
	delete(s.data, value)
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) AppendSlice(values []string) {
	for _, value := range values {
		s.Append(value)
	}
}

func (s *Set) ToSlice() []string {
	result := make([]string, 0)
	for value := range s.data {
		result = append(result, value)
	}
	return result
}

func (s *Set) ToString() string {
	result := "Tags:"
	for value := range s.data {
		result += value + ","
	}

	return strings.TrimSuffix(result, ",")
}
