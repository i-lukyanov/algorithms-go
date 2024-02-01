package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ReverseInt struct{}

func reverse(x int) int {
	xs := strconv.Itoa(x)
	var rs, minus string
	var res int

	for _, d := range xs {
		if string(d) == "-" {
			minus = string(d)

			continue
		}

		if len(rs) == 0 && string(d) == "0" {
			continue
		}

		rs = string(d) + rs
	}

	if isBeyond(rs, minus) {
		return 0
	}

	res, _ = strconv.Atoi(minus + rs)

	return res
}

func isBeyond(s, minus string) bool {
	minInt := strconv.Itoa(math.MinInt32)
	minInt = minInt[1:]
	maxInt := strconv.Itoa(math.MaxInt32)

	compMax := strings.Compare(s, maxInt)
	compMin := strings.Compare(s, minInt)
	if (len(minus) == 0 && (len(s) > len(maxInt) || (len(s) == len(maxInt) && compMax == 1))) ||
		(len(minus) == 1 && (len(s) > len(minInt) || (len(s) == len(minInt) && compMin == 1))) {
		return true
	}

	return false
}

var reverseIntCases = []int{
	123,
	-123,
	120,
	-34560,
	595959595,
	-1234567898,
}

func (e ReverseInt) Run() {
	for _, ex := range reverseIntCases {
		res := reverse(ex)
		fmt.Println(fmt.Sprintf("result: %d", res))
	}
}
