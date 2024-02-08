package common

type StackInt struct {
	items []int
}

func (si *StackInt) Push(item int) {
	si.items = append(si.items, item)
}

func (si *StackInt) Pop() (item int, ok bool) {
	sl := len(si.items)
	if sl == 0 {
		return 0, false
	}

	item = si.items[sl-1]
	si.items = si.items[:sl-1]

	return item, true
}

type StackString struct {
	items []string
}

func (ss *StackString) Push(item string) {
	ss.items = append(ss.items, item)
}

func (ss *StackString) Pop() (item string, ok bool) {
	sl := len(ss.items)
	if sl == 0 {
		return "", false
	}

	item = ss.items[sl-1]
	ss.items = ss.items[:sl-1]

	return item, true
}

func (ss *StackString) Len() int {
	return len(ss.items)
}
