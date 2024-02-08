package leetcode

import (
	"algorithms/pkg/common"
	"fmt"
)

const (
	parOpenRnd  = "("
	parOpenSq   = "["
	parOpenFig  = "{"
	parCloseRnd = ")"
	parCloseSq  = "]"
	parCloseFig = "}"
)

type ValidParentheses struct{}

func (e ValidParentheses) isValid(s string) bool {
	var pStack common.StackString
	sl := len(s)

	for i := 0; i < sl; i++ {
		p := string(s[i])

		if pStack.Len() == 0 && e.isParenthesisClose(p) {
			return false
		}

		if e.isParenthesisOpen(p) {
			pStack.Push(p)

			continue
		}

		if e.isParenthesisClose(p) {
			po, ok := pStack.Pop()
			if ok == true && e.match(po, p) {
				continue
			}
		}

		return false
	}

	return pStack.Len() == 0
}

func (e ValidParentheses) isParenthesisOpen(p string) bool {
	return p == parOpenRnd || p == parOpenSq || p == parOpenFig
}

func (e ValidParentheses) isParenthesisClose(p string) bool {
	return p == parCloseRnd || p == parCloseSq || p == parCloseFig
}

func (e ValidParentheses) match(openPar, closePar string) bool {
	if closePar == parCloseRnd {
		return openPar == parOpenRnd
	}

	if closePar == parCloseSq {
		return openPar == parOpenSq
	}

	if closePar == parCloseFig {
		return openPar == parOpenFig
	}

	return false
}

var validParenthesesCases = []string{
	"()",         // true
	"()[]{}",     // true
	"(}",         // false
	"][",         // false
	"([)]",       // false
	"()[({()})]", // true
}

func (e ValidParentheses) Run() {
	for _, ex := range validParenthesesCases {
		res := e.isValid(ex)
		fmt.Println(fmt.Sprintf("result: %t", res))
	}
}
