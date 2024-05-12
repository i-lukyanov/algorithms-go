package yandex_alg102

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IMinesweeper struct{}

type mine struct {
	row, col int
}

func (e IMinesweeper) Run() {
	rd := bufio.NewReader(os.Stdin)

	fieldConfigStr, _ := rd.ReadString('\n')

	fieldConfigStr = strings.Replace(fieldConfigStr, "\n", "", -1)

	fieldConfig := strings.Split(strings.TrimSpace(fieldConfigStr), " ")
	ln := len(fieldConfig)
	fieldConfigInt := make([]int, ln)

	for k, v := range fieldConfig {
		num, _ := strconv.Atoi(v)
		fieldConfigInt[k] = num
	}

	if len(fieldConfigInt) < 3 {
		fmt.Println("error")

		return
	}

	mines := make([]mine, 0)

	for i := 0; i < fieldConfigInt[2]; i++ {
		m, _ := rd.ReadString('\n')

		m = strings.Replace(m, "\n", "", -1)
		ms := strings.Split(strings.TrimSpace(m), " ")

		if len(ms) < 2 {
			fmt.Println("error")

			return
		}

		var mm mine
		mm.row, _ = strconv.Atoi(ms[0])
		mm.col, _ = strconv.Atoi(ms[1])

		mines = append(mines, mm)
	}

	res := e.getMineField(fieldConfigInt, mines)

	for _, rr := range res {
		fmt.Println(rr)
	}
}

func (e IMinesweeper) getMineField(cnf []int, mines []mine) (res []string) {
	field := make(map[int]map[int]string, cnf[0])

	for i := 1; i <= cnf[0]; i++ {
		row := make(map[int]string, cnf[1])

		for j := 1; j <= cnf[1]; j++ {
			row[j] = "0"
		}

		field[i] = row
	}

	for _, m := range mines {
		field[m.row][m.col] = "*"

		lt, ok := field[m.row-1][m.col-1]
		if ok && lt != "*" {
			lti, _ := strconv.Atoi(lt)
			lti++
			field[m.row-1][m.col-1] = strconv.Itoa(lti)
		}

		t, ok := field[m.row-1][m.col]
		if ok && t != "*" {
			ti, _ := strconv.Atoi(t)
			ti++
			field[m.row-1][m.col] = strconv.Itoa(ti)
		}

		rt, ok := field[m.row-1][m.col+1]
		if ok && rt != "*" {
			rti, _ := strconv.Atoi(rt)
			rti++
			field[m.row-1][m.col+1] = strconv.Itoa(rti)
		}

		l, ok := field[m.row][m.col-1]
		if ok && l != "*" {
			li, _ := strconv.Atoi(l)
			li++
			field[m.row][m.col-1] = strconv.Itoa(li)
		}

		r, ok := field[m.row][m.col+1]
		if ok && r != "*" {
			ri, _ := strconv.Atoi(r)
			ri++
			field[m.row][m.col+1] = strconv.Itoa(ri)
		}

		lb, ok := field[m.row+1][m.col-1]
		if ok && lb != "*" {
			lbi, _ := strconv.Atoi(lb)
			lbi++
			field[m.row+1][m.col-1] = strconv.Itoa(lbi)
		}

		b, ok := field[m.row+1][m.col]
		if ok && b != "*" {
			bi, _ := strconv.Atoi(b)
			bi++
			field[m.row+1][m.col] = strconv.Itoa(bi)
		}

		rb, ok := field[m.row+1][m.col+1]
		if ok && rb != "*" {
			rbi, _ := strconv.Atoi(rb)
			rbi++
			field[m.row+1][m.col+1] = strconv.Itoa(rbi)
		}
	}

	for fi := 1; fi <= len(field); fi++ {
		rr := ""

		for fj := 1; fj <= len(field[fi]); fj++ {
			rr += field[fi][fj] + " "
		}

		rr = strings.TrimSpace(rr)
		res = append(res, rr)
	}

	return res
}
