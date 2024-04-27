package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type H3Biggest struct{}

func (e H3Biggest) Run() {
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

func (e H3Biggest) getNumsBiggest(lst []string) string {
	ln := len(lst)
	lstInt := make([]int, ln)

	for k, v := range lst {
		num, _ := strconv.Atoi(v)
		lstInt[k] = num
	}

	max1, max2, max3 := lstInt[0], lstInt[1], lstInt[2]
	min1, min2 := lstInt[1], lstInt[0]

	if lstInt[0] < lstInt[1] {
		min1, min2 = lstInt[0], lstInt[1]
	}

	if lstInt[1] > lstInt[0] && lstInt[1] > lstInt[2] && lstInt[2] > lstInt[0] {
		max1, max2, max3 = lstInt[1], lstInt[2], lstInt[0]
	} else if lstInt[1] > lstInt[0] && lstInt[1] > lstInt[2] && lstInt[0] > lstInt[2] {
		max1, max2, max3 = lstInt[1], lstInt[0], lstInt[2]
	} else if lstInt[2] > lstInt[0] && lstInt[2] > lstInt[1] && lstInt[0] > lstInt[1] {
		max1, max2, max3 = lstInt[2], lstInt[0], lstInt[1]
	} else if lstInt[2] > lstInt[0] && lstInt[2] > lstInt[1] && lstInt[1] > lstInt[0] {
		max1, max2, max3 = lstInt[2], lstInt[1], lstInt[0]
	} else if lstInt[2] > lstInt[1] {
		max1, max2, max3 = lstInt[0], lstInt[2], lstInt[1]
	}

	for i := 2; i < ln; i++ {
		if lstInt[i] <= min1 {
			min2 = min1
			min1 = lstInt[i]
		} else if lstInt[i] >= min1 && lstInt[i] <= min2 {
			min2 = lstInt[i]
		}
	}

	for i := 3; i < ln; i++ {
		if lstInt[i] >= max1 {
			max3 = max2
			max2 = max1
			max1 = lstInt[i]
		} else if lstInt[i] <= max1 && lstInt[i] >= max2 {
			max3 = max2
			max2 = lstInt[i]
		} else if lstInt[i] <= max2 && lstInt[i] >= max3 {
			max3 = lstInt[i]
		}
	}

	if min1 < 0 && min2 < 0 && (max1*max2*max3) < (min1*min2*max1) {
		return fmt.Sprintf("%d %d %d", min1, min2, max1)
	}

	return fmt.Sprintf("%d %d %d", max3, max2, max1)
}
