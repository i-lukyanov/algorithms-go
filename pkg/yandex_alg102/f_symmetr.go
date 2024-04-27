package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FSymmetr struct{}

func (e FSymmetr) Run() {
	rd := bufio.NewReader(os.Stdin)

	sCnt, _ := rd.ReadString('\n')
	sList, _ := rd.ReadString('\n')

	sCnt = strings.Replace(sCnt, "\n", "", -1)
	sList = strings.Replace(sList, "\n", "", -1)

	listSlice := strings.Split(strings.TrimSpace(sList), " ")
	if len(listSlice) == 1 && listSlice[0] == "" {
		listSlice = []string{}
	}

	cnt, _ := strconv.Atoi(strings.TrimSpace(sCnt))
	if cnt != len(listSlice) {
		fmt.Println("error")

		return
	}

	resCnt, resNums := e.getSymmetrNums(listSlice)

	fmt.Println(resCnt)
	if len(resNums) > 0 {
		fmt.Println(resNums)
	}
}

func (e FSymmetr) getSymmetrNums(lst []string) (string, string) {
	ln := len(lst)

	if e.isSymmetr(lst) {
		return "0", ""
	}

	addonLen := 1
	var addon []string

	for i := 0; i < ln-1; i++ {
		a := make([]string, 0)
		for j := i; j >= 0; j-- {
			a = append(a, lst[j])
		}

		symm := append(lst, a...)
		if e.isSymmetr(symm) {
			addon = a
			addonLen = len(addon)

			break
		}
	}

	return strconv.Itoa(addonLen), strings.Join(addon, " ")
}

func (e FSymmetr) isSymmetr(lst []string) bool {
	ln := len(lst)
	hf := ln / 2
	if ln%2 == 1 {
		hf = (ln - 1) / 2
	}

	left := lst[:hf]
	right := make([]string, hf)
	for i := 0; i < hf; i++ {
		right[i] = lst[ln-i-1]
	}

	eq := true
	for i := 0; i < hf; i++ {
		if left[i] != right[i] {
			eq = false

			break
		}
	}

	return eq
}
