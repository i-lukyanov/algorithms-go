package yandex_alg101

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type JSystem struct{}

func (e JSystem) Run() {
	rd := bufio.NewReader(os.Stdin)

	as, _ := rd.ReadString('\n')
	bs, _ := rd.ReadString('\n')
	cs, _ := rd.ReadString('\n')
	ds, _ := rd.ReadString('\n')
	es, _ := rd.ReadString('\n')
	fs, _ := rd.ReadString('\n')

	as = strings.Replace(as, "\n", "", -1)
	bs = strings.Replace(bs, "\n", "", -1)
	cs = strings.Replace(cs, "\n", "", -1)
	ds = strings.Replace(ds, "\n", "", -1)
	es = strings.Replace(es, "\n", "", -1)
	fs = strings.Replace(fs, "\n", "", -1)

	fmt.Println(e.getSystemSolution(as, bs, cs, ds, es, fs))
}

func (e JSystem) getSystemSolution(as, bs, cs, ds, es, fs string) string {
	var x, y float64

	aa, _ := strconv.Atoi(as)
	bb, _ := strconv.Atoi(bs)
	cc, _ := strconv.Atoi(cs)
	dd, _ := strconv.Atoi(ds)
	ee, _ := strconv.Atoi(es)
	ff, _ := strconv.Atoi(fs)

	if aa == 0 && cc == 0 && bb == 0 && dd == 0 {
		return "5"
	}

	disc := aa*dd - bb*cc
	discX := ee*dd - ff*bb
	discY := aa*ff - ee*cc

	if disc != 0 {
		x := float64(discX) / float64(disc)
		y := float64(discY) / float64(disc)

		return fmt.Sprintf("2 %.5f %.5f", x, y)
	}

	if discX != 0 && discY != 0 {
		return "0"
	}

	if discX == 0 && discY == 0 {
		if bb == 0 && dd == 0 {
			if aa == 0 {
				x = float64(ff) / float64(cc)
			} else {
				x = float64(ee) / float64(aa)
			}

			return fmt.Sprintf("3 %.5f", x)
		}

		if aa == 0 && cc == 0 {
			if bb == 0 {
				y = float64(ff) / float64(dd)
			} else {
				y = float64(ee) / float64(bb)
			}

			return fmt.Sprintf("4 %.5f", y)
		}

		var k, b float64
		if bb == 0 {
			k = float64((-1)*cc) / float64(dd)
			b = float64(ff) / float64(dd)
		} else {
			k = float64((-1)*aa) / float64(bb)
			b = float64(ee) / float64(bb)
		}

		return fmt.Sprintf("1 %.5f %.5f", k, b)
	}

	return "0"
}
