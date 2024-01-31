package leetcode

import (
	"fmt"
	"strconv"
)

type PalindromeNumber struct{}

func isPalindrome(x int) bool {
	xs := strconv.Itoa(x)
	xsr := ""

	for _, l := range xs {
		xsr = string(l) + xsr
	}

	return xs == xsr
}

var palNumCases = []int{
	121,
	-121,
	10,
	595959595,
}

func (e PalindromeNumber) Run() {
	for _, ex := range palNumCases {
		res := isPalindrome(ex)
		fmt.Println(fmt.Sprintf("result: %t", res))
	}
}
