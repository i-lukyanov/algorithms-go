package yandex_alg102

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CClosestNum struct{}

func (e CClosestNum) Run() {
	rd := bufio.NewReader(os.Stdin)

	sArrLen, _ := rd.ReadString('\n')
	sArr, _ := rd.ReadString('\n')
	sNum, _ := rd.ReadString('\n')

	sArrLen = strings.Replace(sArrLen, "\n", "", -1)
	sArr = strings.Replace(sArr, "\n", "", -1)
	sNum = strings.Replace(sNum, "\n", "", -1)

	arrLen, _ := strconv.Atoi(sArrLen)
	num, _ := strconv.Atoi(sNum)
	sSlice := strings.Split(sArr, " ")

	if arrLen != len(sSlice) {
		return
	}

	arr := make([]int, len(sSlice))
	for k, elem := range sSlice {
		arrElem, _ := strconv.Atoi(elem)
		arr[k] = arrElem
	}

	res := e.getClosestNum(arr, num)

	fmt.Println(strconv.Itoa(res))
}

func (e CClosestNum) getClosestNum(arr []int, num int) int {
	closest := arr[0]

	for i := 1; i < len(arr); i++ {
		if math.Abs(float64(arr[i]-num)) < math.Abs(float64(closest-num)) {
			closest = arr[i]
		}
	}

	return closest
}
