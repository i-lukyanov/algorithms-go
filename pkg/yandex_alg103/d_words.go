package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type DWords struct{}

func (e DWords) Run() {
	rd := bufio.NewReader(os.Stdin)
	fullStr := ""

	for {
		wordsStr, err := rd.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		wordsStr = strings.Replace(wordsStr, "\n", " ", -1)
		if len(wordsStr) <= 1 {
			break
		}

		fullStr += wordsStr
	}

	words := strings.Fields(fullStr)

	s := common.NewSet()

	for _, w := range words {
		if !s.Contains(w) {
			s.Add(w)
		}
	}

	fmt.Println(strconv.Itoa(s.Size()))
}
