package leetcode

import (
	"fmt"
	"strings"
)

type LongestPalindromeSubstr struct{}

func (e LongestPalindromeSubstr) longestPalindrome(s string) string {
	ts := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	tsLen := len(ts)
	pal := make([]int, tsLen)
	center, radius := 0, 0

	for i := 1; i < tsLen-1; i++ {
		if i < radius {
			pal[i] = e.min(radius-i, pal[2*center-i])
		}

		for ts[i+1+pal[i]] == ts[i-1-pal[i]] {
			pal[i]++
		}

		if i+pal[i] > radius {
			center, radius = i, i+pal[i]
		}
	}

	maxLen := 0
	centerIndex := 0

	for k, p := range pal {
		if p > maxLen {
			maxLen = p
			centerIndex = k
		}
	}

	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func (e LongestPalindromeSubstr) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func (e LongestPalindromeSubstr) longestPalindrome(s string) string {
//	ls := len(s)
//	if ls == 0 {
//		return ""
//	}
//
//	res := ""
//
//	for i, j := 0, ls; i < j; {
//		ps := s[i:j]
//
//		if !e.isPalindrome(ps) {
//			j--
//
//			continue
//		}
//
//		if len(ps) > len(res) {
//			res = ps
//		}
//
//		i++
//		j = ls
//	}
//
//	return res
//}
//
//func (e LongestPalindromeSubstr) isPalindrome(s string) bool {
//	var rs string
//	for _, l := range s {
//		rs = string(l) + rs
//	}
//
//	return s == rs
//}

var palSubstrCases = []string{
	"babad",
	"cbbd",
	"abc",
	"",
	"123ytrrtyytrurlty321",
}

func (e LongestPalindromeSubstr) Run() {
	for _, ex := range palSubstrCases {
		res := e.longestPalindrome(ex)
		fmt.Println(fmt.Sprintf("result: %s", res))
	}
}
