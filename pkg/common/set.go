package common

type Set struct {
	data map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[interface{}]struct{})}
}

func (s *Set) Add(item interface{}) {
	s.data[item] = struct{}{}
}

func (s *Set) Remove(item interface{}) {
	delete(s.data, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.data[item]

	return ok
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Intersect(other *Set) *Set {
	intersection := NewSet()

	if other.Size() > s.Size() {
		for item := range s.data {
			if other.Contains(item) {
				intersection.Add(item)
			}
		}
	} else {
		for item := range other.data {
			if s.Contains(item) {
				intersection.Add(item)
			}
		}
	}

	return intersection
}

func (s *Set) ToSlice() []interface{} {
	sl := make([]interface{}, 0)
	for dd := range s.data {
		sl = append(sl, dd)
	}

	return sl
}
