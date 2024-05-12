package yandex_alg102

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type JTriangle struct{}

type note struct {
	freq float64
	dist string
}

func (e JTriangle) Run() {
	rd := bufio.NewReader(os.Stdin)

	cntStr, _ := rd.ReadString('\n')
	cntStr = strings.Replace(cntStr, "\n", "", -1)
	cnt, _ := strconv.Atoi(cntStr)

	notes := make([]note, 0)

	for i := 0; i < cnt; i++ {
		n, _ := rd.ReadString('\n')

		n = strings.Replace(n, "\n", "", -1)
		ns := strings.Split(strings.TrimSpace(n), " ")

		var nn note
		nn.freq, _ = common.StringToFloat(ns[0])
		if len(ns) == 2 {
			nn.dist = ns[1]
		}

		notes = append(notes, nn)
	}

	res := e.getFreqRange(notes)

	fmt.Println(res)
}

func (e JTriangle) getFreqRange(notes []note) string {
	left, right := 30.0, 4000.0
	prev := notes[0].freq

	for i := 1; i < len(notes); i++ {
		now := notes[i].freq

		if math.Abs(now-prev) < math.Pow(10.0, -6.0) {
			continue
		}

		if notes[i].dist == "closer" {
			if now > prev {
				left = math.Max(left, (now+prev)/2)
			} else {
				right = math.Min(right, (now+prev)/2)
			}
		} else {
			if now < prev {
				left = math.Max(left, (now+prev)/2)
			} else {
				right = math.Min(right, (now+prev)/2)
			}
		}

		prev = now
	}

	return fmt.Sprintf("%.6f %.6f", left, right)
}
