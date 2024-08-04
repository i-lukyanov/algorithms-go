package common

type MultiSetString struct {
	data map[string]int
}

func NewMultiSetString() *MultiSetString {
	return &MultiSetString{data: make(map[string]int)}
}

func (mss *MultiSetString) Add(item string) {
	if mss.Contains(item) {
		mss.data[item]++
	} else {
		mss.data[item] = 1
	}
}

func (mss *MultiSetString) Remove(item string) {
	if mss.Contains(item) {
		mss.data[item]--

		if mss.data[item] <= 0 {
			delete(mss.data, item)
		}
	}
}

func (mss *MultiSetString) Contains(item string) bool {
	_, ok := mss.data[item]

	return ok
}

func (mss *MultiSetString) Size() int {
	cnt := 0
	for _, v := range mss.data {
		cnt += v
	}

	return cnt
}

func (mss *MultiSetString) Count() int {
	return len(mss.data)
}

func (mss *MultiSetString) ToMap() map[string]int {
	return mss.data
}

//func (mss *MultiSetString) Intersect(other *MultiSetString) *MultiSetString {
//	intersection := NewMultiSetString()
//
//	if other.Size() > mss.Size() {
//		for item := range mss.data {
//			if other.Contains(item) {
//				intersection.Add(item)
//			}
//		}
//	} else {
//		for item := range other.data {
//			if mss.Contains(item) {
//				intersection.Add(item)
//			}
//		}
//	}
//
//	return intersection
//}
//
//func (mss *MultiSetString) Diff(other *MultiSetString) *MultiSetString {
//	diff := NewMultiSetString()
//
//	for item := range mss.data {
//		if !other.Contains(item) {
//			diff.Add(item)
//		}
//	}
//
//	return diff
//}
