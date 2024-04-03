package yandex_alg101

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type DSqrtEquasion struct{}

func (e DSqrtEquasion) Run() {
	rd := bufio.NewReader(os.Stdin)

	sa, _ := rd.ReadString('\n')
	sb, _ := rd.ReadString('\n')
	sc, _ := rd.ReadString('\n')

	sa = strings.Replace(sa, "\n", "", -1)
	sb = strings.Replace(sb, "\n", "", -1)
	sc = strings.Replace(sc, "\n", "", -1)

	a, _ := strconv.Atoi(sa)
	b, _ := strconv.Atoi(sb)
	c, _ := strconv.Atoi(sc)

	fmt.Println(e.solution(a, b, c))
}

func (e DSqrtEquasion) solution(a, b, c int) string {
	if c < 0 {
		return "NO SOLUTION"
	}

	if float64(c) >= math.Sqrt(float64(math.MaxInt32)) {
		return "NO SOLUTION"
	}

	if c*c-b >= math.MaxInt32 {
		return "NO SOLUTION"
	}

	if a == 0 && ((c*c - b) == 0) {
		return "MANY SOLUTIONS"
	}

	if a == 0 {
		return "NO SOLUTION"
	}

	if ((c*c - b) % a) != 0 {
		return "NO SOLUTION"
	}

	return strconv.Itoa((c*c - b) / a)
}
