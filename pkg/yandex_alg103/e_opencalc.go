package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EOpenCalc struct{}

func (e EOpenCalc) Run() {
	rd := bufio.NewReader(os.Stdin)

	digitsStr, _ := rd.ReadString('\n')
	nmbrStr, _ := rd.ReadString('\n')

	digitsStr = strings.Replace(digitsStr, "\n", "", -1)
	nmbrStr = strings.Replace(nmbrStr, "\n", "", -1)

	digits := strings.Split(strings.TrimSpace(digitsStr), " ")
	nmbrs := strings.Split(strings.TrimSpace(nmbrStr), "")

	dSet := common.NewSetInt()
	nmbrSet := common.NewSetInt()

	for k1 := range digits {
		digit, _ := strconv.Atoi(digits[k1])
		dSet.Add(digit)
	}

	for nn := range nmbrs {
		nmbr, _ := strconv.Atoi(nmbrs[nn])
		if !dSet.Contains(nmbr) {
			nmbrSet.Add(nmbr)
		}
	}

	fmt.Println(fmt.Sprintf("%d", len(nmbrSet.ToSlice())))
}
