package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ILang struct{}

func (e ILang) Run() {
	rd := bufio.NewReader(os.Stdin)

	cntStr, _ := rd.ReadString('\n')
	cntStr = strings.TrimSpace(strings.Replace(cntStr, "\n", "", -1))
	cnt, _ := strconv.Atoi(cntStr)

	langSet := common.NewMultiSetString()
	for i := 0; i < cnt; i++ {
		studCntStr, _ := rd.ReadString('\n')
		studCntStr = strings.TrimSpace(strings.Replace(studCntStr, "\n", "", -1))
		studCnt, _ := strconv.Atoi(studCntStr)

		for j := 0; j < studCnt; j++ {
			langStr, _ := rd.ReadString('\n')
			lang := strings.TrimSpace(strings.Replace(langStr, "\n", "", -1))

			langSet.Add(lang)
		}
	}

	cntAll := 0
	cntAny := langSet.Count()
	langAll := make([]string, 0, cntAny)
	langAny := make([]string, 0, cntAny)

	for l, lc := range langSet.ToMap() {
		langAny = append(langAny, l)

		if lc == cnt {
			cntAll++
			langAll = append(langAll, l)
		}
	}

	fmt.Println(fmt.Sprintf("%d", cntAll))
	for _, lal := range langAll {
		fmt.Println(lal)
	}

	fmt.Println(fmt.Sprintf("%d", cntAny))
	for _, lan := range langAny {
		fmt.Println(lan)
	}
}
