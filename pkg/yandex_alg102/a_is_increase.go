package yandex_alg102

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AIsIncrease struct{}

func (e AIsIncrease) Run() {
	rd := bufio.NewReader(os.Stdin)

	sList, _ := rd.ReadString('\n')

	sList = strings.Replace(sList, "\n", "", -1)

	listSlice := strings.Split(sList, " ")
	res := e.isListIncrease(listSlice)

	fmt.Println(res)
}

func (e AIsIncrease) isListIncrease(lst []string) string {
	lstInt := make([]int, len(lst))
	for k, itm := range lst {
		ii, _ := strconv.Atoi(itm)
		lstInt[k] = ii
	}

	for i := 1; i < len(lstInt); i++ {
		if lstInt[i-1] >= lstInt[i] {
			return common.No
		}
	}

	return common.Yes
}
