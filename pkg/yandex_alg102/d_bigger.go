package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DBigger struct{}

func (e DBigger) Run() {
	rd := bufio.NewReader(os.Stdin)

	sList, _ := rd.ReadString('\n')

	sList = strings.Replace(sList, "\n", "", -1)

	listSlice := strings.Split(sList, " ")

	res := e.isBiggerNeighbors(listSlice)

	fmt.Println(res)
}

func (e DBigger) isBiggerNeighbors(lst []string) string {
	lstInt := make([]int, len(lst))
	for k, itm := range lst {
		ii, _ := strconv.Atoi(itm)
		lstInt[k] = ii
	}

	cnt := 0

	for i := 1; i < len(lstInt)-1; i++ {
		if lstInt[i] > lstInt[i-1] && lstInt[i] > lstInt[i+1] {
			cnt++
		}
	}

	return strconv.Itoa(cnt)
}
