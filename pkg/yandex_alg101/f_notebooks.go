package yandex_alg101

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FNotebooks struct{}

func (e FNotebooks) Run() {
	rd := bufio.NewReader(os.Stdin)

	sizes, _ := rd.ReadString('\n')

	sizes = strings.Replace(sizes, "\n", "", -1)

	szSlice := strings.Split(sizes, " ")
	res := e.getMinTable(szSlice)

	fmt.Println(res)
}

func (e FNotebooks) getMinTable(sizes []string) string {
	type rect struct {
		len, wid, sq int
	}

	var n1, n2, table rect
	lens := make([]int, len(sizes))
	tables := make([]rect, 2*len(sizes))
	minSq := 0

	for k := range sizes {
		lens[k], _ = strconv.Atoi(sizes[k])
	}

	n1.len = lens[0]
	n1.wid = lens[1]
	n1.sq = n1.len * n1.wid
	n2.len = lens[2]
	n2.wid = lens[3]
	n2.sq = n2.len * n2.wid

	tables[0] = rect{len: n1.len + n2.len, wid: n1.wid, sq: (n1.len + n2.len) * n1.wid}
	tables[1] = rect{len: n1.len + n2.len, wid: n2.wid, sq: (n1.len + n2.len) * n2.wid}
	tables[2] = rect{len: n1.len + n2.wid, wid: n1.wid, sq: (n1.len + n2.wid) * n1.wid}
	tables[3] = rect{len: n1.len + n2.wid, wid: n2.len, sq: (n1.len + n2.wid) * n2.len}
	tables[4] = rect{len: n1.wid + n2.len, wid: n1.len, sq: (n1.wid + n2.len) * n1.len}
	tables[5] = rect{len: n1.wid + n2.len, wid: n2.wid, sq: (n1.wid + n2.len) * n2.wid}
	tables[6] = rect{len: n1.wid + n2.wid, wid: n1.len, sq: (n1.wid + n2.wid) * n1.len}
	tables[7] = rect{len: n1.wid + n2.wid, wid: n2.len, sq: (n1.wid + n2.wid) * n2.len}

	for _, t := range tables {
		if t.sq < (n1.sq + n2.sq) {
			continue
		}

		if minSq == 0 || t.sq < minSq {
			minSq = t.sq
			table = t
		}
	}

	return fmt.Sprintf("%d %d", table.len, table.wid)
}
