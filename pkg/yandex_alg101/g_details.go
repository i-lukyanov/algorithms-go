package yandex_alg101

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GDetails struct{}

func (e GDetails) Run() {
	rd := bufio.NewReader(os.Stdin)

	weights, _ := rd.ReadString('\n')

	weights = strings.Replace(weights, "\n", "", -1)

	wSlice := strings.Split(weights, " ")

	fmt.Println(e.getDetailsCount(wSlice))
}

func (e GDetails) getDetailsCount(wSlice []string) string {
	weights := make([]int, len(wSlice))

	for k := range wSlice {
		weights[k], _ = strconv.Atoi(wSlice[k])
	}

	if weights[0] < weights[1] || weights[0] < weights[2] || weights[1] < weights[2] {
		return "0"
	}

	initWeight := weights[0]
	detCnt := 0

	for initWeight >= weights[1] {
		detCnt = weights[1]/weights[2] + detCnt
		initWeight = initWeight - weights[1] + weights[1]%weights[2]
	}

	return fmt.Sprintf("%d", detCnt)
}
