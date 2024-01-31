package leetcode

import (
	"algorithms/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

type ReversePolishNotation struct{}

const (
	opPlus  = "+"
	opMinus = "-"
	opProd  = "*"
	opDiv   = "/"
)

func evalRPNwithStack(tokens []string) int {
	var st common.StackInt

	for _, token := range tokens {
		if isOperator(token) {
			op2, _ := st.Pop()
			op1, _ := st.Pop()

			st.Push(eval(op1, op2, token))
		} else {
			item, _ := strconv.Atoi(token)
			st.Push(item)
		}
	}

	item, _ := st.Pop()

	return item
}

func evalRPN(tokens []string) int {
	tl := len(tokens)
	if tl == 1 {
		r, _ := strconv.Atoi(tokens[0])

		return r
	}

	if tl < 3 || tl%2 == 0 {
		return 0
	}

	if isOperator(tokens[0]) || isOperator(tokens[1]) {
		return 0
	}

	resStr, err := strconv.Atoi(evalRPNrecursive(tokens))
	if err != nil {
		return 0
	}

	return resStr
}

func evalRPNrecursive(tokens []string) string {
	tl := len(tokens)

	if tl == 3 {
		op1i, _ := strconv.Atoi(tokens[0])
		op2i, _ := strconv.Atoi(tokens[1])
		ev := eval(op1i, op2i, tokens[2])
		return strconv.Itoa(ev)
	}

	var reducedTokens []string
	for i := 2; i < tl+2; {
		if i < tl && isOperator(tokens[i]) && !isOperator(tokens[i-1]) && !isOperator(tokens[i-2]) {
			op1i, _ := strconv.Atoi(tokens[i-2])
			op2i, _ := strconv.Atoi(tokens[i-1])
			reducedTokens = append(reducedTokens, strconv.Itoa(eval(op1i, op2i, tokens[i])))
			i += 3

			continue
		}

		reducedTokens = append(reducedTokens, tokens[i-2])
		i++
	}

	rr := evalRPNrecursive(reducedTokens)
	return rr
}

func isOperator(t string) bool {
	ops := opPlus + opMinus + opProd + opDiv

	return strings.Contains(ops, t)
}

func eval(op1, op2 int, operator string) int {
	if !isOperator(operator) {
		return 0
	}

	switch operator {
	case opPlus:
		return op1 + op2
	case opMinus:
		return op1 - op2
	case opProd:
		return op1 * op2
	case opDiv:
		return int(op1 / op2)
	}

	return 0
}

var RPNCases = [][]string{
	{"2", "1", "+", "3", "*"},
	{"4", "13", "5", "/", "+"},
	{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	{"18"},
}

func (e ReversePolishNotation) Run() {
	for _, ex := range RPNCases {
		res := evalRPNwithStack(ex)
		fmt.Println(fmt.Sprintf("result: %d", res))
	}
}
