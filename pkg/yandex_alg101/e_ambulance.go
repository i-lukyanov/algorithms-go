package yandex_alg101

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EAmbulance struct{}

func (e EAmbulance) Run() {
	rd := bufio.NewReader(os.Stdin)

	apts, _ := rd.ReadString('\n')

	apts = strings.Replace(apts, "\n", "", -1)

	aptSlice := strings.Split(apts, " ")
	res := e.getGateFloor(aptSlice)

	fmt.Println(res)
}

func (e EAmbulance) getGateFloor(aptSlice []string) string {
	resContradict := "-1 -1"
	if len(aptSlice) < 5 {
		return resContradict
	}

	var aptNums []int
	for _, a := range aptSlice {
		aNum, _ := strconv.Atoi(a)

		if aNum <= 0 {
			return resContradict
		}

		aptNums = append(aptNums, aNum)
	}

	var res struct {
		gate, floor int
	}

	res.gate = -1
	res.floor = -1
	ok := false

	for aptsOnFloor := 1; aptsOnFloor < 20; aptsOnFloor++ {
		nGate, nFloor := e.check(aptNums, aptsOnFloor)
		if nGate != -1 {
			ok = true
			if res.gate == -1 {
				res.gate, res.floor = nGate, nFloor
			} else if res.gate != nGate && res.gate != 0 {
				res.gate = 0
			} else if res.floor != nFloor && res.floor != 0 {
				res.floor = 0
			}
		}
	}

	if !ok {
		return resContradict
	}

	return fmt.Sprintf("%d %d", res.gate, res.floor)
}

func (e EAmbulance) getGateFloorByApt(apt, aptsOnFloor, floors int) (int, int) {
	floorsBefore := (apt - 1) / aptsOnFloor
	gate := floorsBefore/floors + 1
	floor := floorsBefore%floors + 1

	return gate, floor
}

func (e EAmbulance) check(aptParams []int, aptsOnFloor int) (int, int) {
	gate2, floor2 := e.getGateFloorByApt(aptParams[2], aptsOnFloor, aptParams[1])

	if gate2 == aptParams[3] && floor2 == aptParams[4] {
		return e.getGateFloorByApt(aptParams[0], aptsOnFloor, aptParams[1])
	}

	return -1, -1
}
