package leetcode

import (
	"fmt"
)

type DailyTemp struct{}

func dailyTemperatures(temperatures []int) []int {
	tlen := len(temperatures)
	td := make([]int, tlen)
	stack := make([]int, 0)

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			td[index] = i - index
		}

		stack = append(stack, i)
	}

	return td
}

var dailyTempCases = [][]int{
	{73, 74, 75, 71, 69, 72, 76, 73},
	{30, 40, 50, 60},
	{30, 60, 90},
	{90, 60, 30, 1},
}

func (e DailyTemp) Run() {
	for _, ex := range dailyTempCases {
		res := dailyTemperatures(ex)
		fmt.Println(fmt.Sprintf("result: %+v", res))
	}
}
