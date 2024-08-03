package common

type SetInt struct {
	data map[int]struct{}
}

func NewSetInt() *SetInt {
	return &SetInt{data: make(map[int]struct{})}
}

func (s *SetInt) Add(item int) {
	s.data[item] = struct{}{}
}

func (s *SetInt) Remove(item int) {
	delete(s.data, item)
}

func (s *SetInt) Contains(item int) bool {
	_, ok := s.data[item]

	return ok
}

func (s *SetInt) Size() int {
	return len(s.data)
}

func (s *SetInt) Intersect(other *SetInt) *SetInt {
	intersection := NewSetInt()

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

func (s *SetInt) Diff(other *SetInt) *SetInt {
	diff := NewSetInt()

	for item := range s.data {
		if !other.Contains(item) {
			diff.Add(item)
		}
	}

	return diff
}

func (s *SetInt) ToMap() map[int]int {
	sl := make(map[int]int, 0)
	for dd := range s.data {
		sl[dd] = dd
	}

	return sl
}

func (s *SetInt) ToSlice() []int {
	sl := make([]int, 0)
	for dd := range s.data {
		sl = append(sl, dd)
	}

	return sl
}
