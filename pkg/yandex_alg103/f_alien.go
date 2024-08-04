package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FAlien struct{}

func (e FAlien) Run() {
	rd := bufio.NewReader(os.Stdin)

	gen1Str, _ := rd.ReadString('\n')
	gen2Str, _ := rd.ReadString('\n')

	gen1Str = strings.TrimSpace(strings.Replace(gen1Str, "\n", "", -1))
	gen2Str = strings.TrimSpace(strings.Replace(gen2Str, "\n", "", -1))

	gen2Set := common.NewSet()

	for i := 0; i < len(gen2Str)-1; i++ {
		gen2Set.Add(gen2Str[i : i+2])
	}

	genCounter := 0
	for j := 0; j < len(gen1Str)-1; j++ {
		if gen2Set.Contains(gen1Str[j : j+2]) {
			genCounter++
		}
	}

	fmt.Println(fmt.Sprintf("%d", genCounter))
}
