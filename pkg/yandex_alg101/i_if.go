package yandex_alg101

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IIf struct{}

func (e IIf) Run() {
	rd := bufio.NewReader(os.Stdin)

	b1s, _ := rd.ReadString('\n')
	b2s, _ := rd.ReadString('\n')
	b3s, _ := rd.ReadString('\n')
	h1s, _ := rd.ReadString('\n')
	h2s, _ := rd.ReadString('\n')

	b1s = strings.Replace(b1s, "\n", "", -1)
	b2s = strings.Replace(b2s, "\n", "", -1)
	b3s = strings.Replace(b3s, "\n", "", -1)
	h1s = strings.Replace(h1s, "\n", "", -1)
	h2s = strings.Replace(h2s, "\n", "", -1)

	fmt.Println(e.isSuitable(b1s, b2s, b3s, h1s, h2s))
}

func (e IIf) isSuitable(b1s, b2s, b3s, h1s, h2s string) string {
	b1, _ := strconv.Atoi(b1s)
	b2, _ := strconv.Atoi(b2s)
	b3, _ := strconv.Atoi(b3s)
	h1, _ := strconv.Atoi(h1s)
	h2, _ := strconv.Atoi(h2s)

	if (b1 <= h1 && (b2 <= h2 || b3 <= h2)) || (b2 <= h1 && (b1 <= h2 || b3 <= h2)) || (b3 <= h1 && (b2 <= h2 || b1 <= h2)) {
		return common.Yes
	}

	return common.No
}
