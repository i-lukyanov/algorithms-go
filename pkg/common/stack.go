package common

type StackInt struct {
	items []int
}

func (s *StackInt) Push(item int) {
	s.items = append(s.items, item)
}

func (s *StackInt) Pop() (item int, ok bool) {
	sl := len(s.items)
	if sl == 0 {
		return 0, false
	}

	item = s.items[sl-1]
	s.items = s.items[:sl-1]

	return item, true
}
