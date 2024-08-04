package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type GTurtles struct{}

func (e GTurtles) Run() {
	rd := bufio.NewReader(os.Stdin)

	cntStr, _ := rd.ReadString('\n')
	cntStr = strings.TrimSpace(strings.Replace(cntStr, "\n", "", -1))
	cnt, _ := strconv.Atoi(cntStr)

	truthSet := common.NewSetInt()
	for i := 0; i < cnt; i++ {
		turtleStr, _ := rd.ReadString('\n')
		turtleStr = strings.TrimSpace(strings.Replace(turtleStr, "\n", "", -1))

		turtle := strings.Split(strings.TrimSpace(turtleStr), " ")
		tFront, _ := strconv.Atoi(turtle[0])
		tBack, _ := strconv.Atoi(turtle[1])

		if tFront >= 0 && tBack >= 0 && math.Abs(float64(tFront+tBack)) == float64(cnt-1) {
			truthSet.Add(tFront)
		}
	}

	fmt.Println(fmt.Sprintf("%d", truthSet.Size()))
}
