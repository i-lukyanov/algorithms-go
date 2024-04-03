package yandex_alg101

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	rHeat   = "heat"
	rFreeze = "freeze"
	rAuto   = "auto"
)

type AConditioner struct{}

func (e AConditioner) Run() {
	rd := bufio.NewReader(os.Stdin)

	temps, _ := rd.ReadString('\n')
	reg, _ := rd.ReadString('\n')

	temps = strings.Replace(temps, "\n", "", -1)
	reg = strings.Replace(reg, "\n", "", -1)

	tSlice := strings.Split(temps, " ")
	res := e.getTemp(tSlice, reg)

	fmt.Println(res)
}

func (e AConditioner) getTemp(tempSlice []string, regime string) int {
	tRoom, _ := strconv.Atoi(tempSlice[0])
	tCond, _ := strconv.Atoi(tempSlice[1])

	if regime == rAuto || (regime == rHeat && tCond >= tRoom) || (regime == rFreeze && tCond <= tRoom) {
		return tCond
	}

	return tRoom
}
