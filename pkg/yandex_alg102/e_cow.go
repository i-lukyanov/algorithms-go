package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ECow struct{}

func (e ECow) Run() {
	rd := bufio.NewReader(os.Stdin)

	sCnt, _ := rd.ReadString('\n')
	sList, _ := rd.ReadString('\n')

	sCnt = strings.Replace(sCnt, "\n", "", -1)
	sList = strings.Replace(sList, "\n", "", -1)

	listSlice := strings.Split(sList, " ")
	cnt, _ := strconv.Atoi(sCnt)
	if cnt != len(listSlice) {
		fmt.Println("error")

		return
	}

	res := e.getMaxPos(listSlice)

	fmt.Println(res)
}

func (e ECow) getMaxPos(dst []string) string {
	maxDist, winDist := 0, 0
	dstInt := make([]int, len(dst))

	for k, itm := range dst {
		ii, _ := strconv.Atoi(itm)
		dstInt[k] = ii

		if ii > winDist {
			winDist = ii
		}
	}

	for i := 1; i < len(dstInt)-1; i++ {
		if dst[i][len(dst[i])-1:] == "5" && dstInt[i] > maxDist && dstInt[i-1] >= dstInt[i] &&
			dstInt[i-1] == winDist && dstInt[i+1] < dstInt[i] {
			maxDist = dstInt[i]
		}
	}

	if maxDist == 0 {
		return "0"
	}

	if maxDist == winDist {
		return "1"
	}

	pos := 1

	for _, dd := range dstInt {
		if dd > maxDist {
			pos++
		}
	}

	return strconv.Itoa(pos)
}
