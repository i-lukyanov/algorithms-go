package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type G2Biggest struct{}

func (e G2Biggest) Run() {
	rd := bufio.NewReader(os.Stdin)

	sList, _ := rd.ReadString('\n')

	sList = strings.Replace(sList, "\n", "", -1)

	listSlice := strings.Split(strings.TrimSpace(sList), " ")
	if len(listSlice) == 1 && listSlice[0] == "" {
		listSlice = []string{}
	}

	res := e.getNumsBiggest(listSlice)

	fmt.Println(res)
}

func (e G2Biggest) getNumsBiggest(lst []string) string {
	ln := len(lst)
	lstInt := make([]int, ln)

	for k, v := range lst {
		num, _ := strconv.Atoi(v)
		lstInt[k] = num
	}

	max1, max2 := lstInt[0], lstInt[1]
	min1, min2 := lstInt[1], lstInt[0]

	if lstInt[0] < lstInt[1] {
		max1, max2 = lstInt[1], lstInt[0]
		min1, min2 = lstInt[0], lstInt[1]
	}

	for i := 2; i < ln; i++ {
		if lstInt[i] >= max1 {
			max2 = max1
			max1 = lstInt[i]
		} else if lstInt[i] <= max1 && lstInt[i] >= max2 {
			max2 = lstInt[i]
		} else if lstInt[i] <= min1 {
			min2 = min1
			min1 = lstInt[i]
		} else if lstInt[i] >= min1 && lstInt[i] <= min2 {
			min2 = lstInt[i]
		}
	}

	if min1 < 0 && min2 < 0 && (max1*max2) < (min1*min2) {
		return fmt.Sprintf("%d %d", min1, min2)
	}

	return fmt.Sprintf("%d %d", max2, max1)
}
