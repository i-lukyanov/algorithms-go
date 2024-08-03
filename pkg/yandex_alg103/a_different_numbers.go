package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ADifferentNumbers struct{}

func (e ADifferentNumbers) Run() {
	rd := bufio.NewReader(os.Stdin)

	numsStr, _ := rd.ReadString('\n')

	numsStr = strings.Replace(numsStr, "\n", "", -1)

	nums := strings.Split(strings.TrimSpace(numsStr), " ")
	ln := len(nums)
	numsInt := make([]int, ln)

	for k, v := range nums {
		num, _ := strconv.Atoi(v)
		numsInt[k] = num
	}

	s := common.NewSet()

	for _, n := range numsInt {
		s.Add(n)
	}

	fmt.Println(strconv.Itoa(s.Size()))
}
