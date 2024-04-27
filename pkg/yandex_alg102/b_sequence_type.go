package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	seqEnd = -2000000000

	seqConst = "CONSTANT"
	seqAsc   = "ASCENDING"
	seqWAsc  = "WEAKLY ASCENDING"
	seqDesc  = "DESCENDING"
	seqWDesc = "WEAKLY DESCENDING"
	seqRand  = "RANDOM"
)

type BSequenceType struct{}

func (e BSequenceType) Run() {
	rd := bufio.NewReader(os.Stdin)
	seq := make([]int, 0)

	for {
		s, _ := rd.ReadString('\n')

		s = strings.Replace(s, "\n", "", -1)

		ss, _ := strconv.Atoi(s)

		if ss == seqEnd {
			break
		}

		seq = append(seq, ss)
	}

	res := e.getSequenceType(seq)

	fmt.Println(res)
}

func (e BSequenceType) getSequenceType(seq []int) string {
	seqType := ""

	for i := 1; i < len(seq); i++ {
		if seq[i-1] == seq[i] {
			if seqType == seqConst || seqType == "" {
				seqType = seqConst

				continue
			}

			if seqType == seqWAsc || seqType == seqWDesc {
				continue
			}

			if seqType == seqAsc {
				seqType = seqWAsc

				continue
			}

			if seqType == seqDesc {
				seqType = seqWDesc

				continue
			}
		}

		if seq[i-1] < seq[i] {
			if seqType == "" {
				seqType = seqAsc

				continue
			}

			if seqType == seqDesc || seqType == seqWDesc {
				return seqRand
			}

			if seqType == seqConst {
				seqType = seqWAsc

				continue
			}

			if seqType == seqAsc || seqType == seqWAsc {
				continue
			}
		}

		if seq[i-1] > seq[i] {
			if seqType == "" {
				seqType = seqDesc

				continue
			}

			if seqType == seqAsc || seqType == seqWAsc {
				return seqRand
			}

			if seqType == seqConst {
				seqType = seqWDesc

				continue
			}

			if seqType == seqDesc || seqType == seqWDesc {
				continue
			}
		}
	}

	return seqType
}
