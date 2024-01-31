package leetcode

import (
	"fmt"
)

type TwoSum struct{}

func twoSum(nums []int, target int) []int {
	nl := len(nums)
	if nl == 0 {
		return []int{}
	}

	nc := nl

	for nc > 1 {
		nc--
		last := nums[nc]
		nums := nums[:nc]

		for i, v := range nums {
			if v+last == target {
				return []int{i, nc}
			}
		}
	}

	return []int{}
}

type twoSumCase struct {
	nums []int
	tgt  int
}

var twoSumCases = []twoSumCase{
	{[]int{2, 7, 11, 15}, 9},
	{[]int{3, 2, 4}, 6},
	{[]int{3, 3}, 6},
	{[]int{2, 3}, 8},
}

func (e TwoSum) Run() {
	for _, ex := range twoSumCases {
		res := twoSum(ex.nums, ex.tgt)
		fmt.Println(fmt.Sprintf("nums: %v, target: %d - result: %v", ex.nums, ex.tgt, res))
	}
}
