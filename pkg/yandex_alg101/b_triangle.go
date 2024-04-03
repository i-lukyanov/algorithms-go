package yandex_alg101

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BTriangle struct{}

func (e BTriangle) Run() {
	rd := bufio.NewReader(os.Stdin)

	s1, _ := rd.ReadString('\n')
	s2, _ := rd.ReadString('\n')
	s3, _ := rd.ReadString('\n')

	s1 = strings.Replace(s1, "\n", "", -1)
	s2 = strings.Replace(s2, "\n", "", -1)
	s3 = strings.Replace(s3, "\n", "", -1)

	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	i3, _ := strconv.Atoi(s3)

	fmt.Println(isTriangle(i1, i2, i3))
}

func isTriangle(i1, i2, i3 int) string {
	if i1+i2 > i3 && i1+i3 > i2 && i2+i3 > i1 {
		return common.Yes
	}

	return common.No
}
