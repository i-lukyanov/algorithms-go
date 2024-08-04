package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HPigs struct{}

func (e HPigs) Run() {
	rd := bufio.NewReader(os.Stdin)

	cntStr, _ := rd.ReadString('\n')
	cntStr = strings.TrimSpace(strings.Replace(cntStr, "\n", "", -1))
	cnt, _ := strconv.Atoi(cntStr)

	birdsSet := common.NewSetInt()
	for i := 0; i < cnt; i++ {
		birdStr, _ := rd.ReadString('\n')
		birdStr = strings.TrimSpace(strings.Replace(birdStr, "\n", "", -1))

		bird := strings.Split(strings.TrimSpace(birdStr), " ")
		birdX, _ := strconv.Atoi(bird[0])

		birdsSet.Add(birdX)
	}

	fmt.Println(fmt.Sprintf("%d", birdsSet.Size()))
}
