package yandex_alg101

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type HMetro struct{}

func (e HMetro) Run() {
	rd := bufio.NewReader(os.Stdin)

	int1s, _ := rd.ReadString('\n')
	int2s, _ := rd.ReadString('\n')
	cnt1s, _ := rd.ReadString('\n')
	cnt2s, _ := rd.ReadString('\n')

	int1s = strings.Replace(int1s, "\n", "", -1)
	int2s = strings.Replace(int2s, "\n", "", -1)
	cnt1s = strings.Replace(cnt1s, "\n", "", -1)
	cnt2s = strings.Replace(cnt2s, "\n", "", -1)

	fmt.Println(e.getMinMax(int1s, int2s, cnt1s, cnt2s))
}

func (e HMetro) getMinMax(int1s, int2s, cnt1s, cnt2s string) string {
	stop := 1

	int1, _ := strconv.Atoi(int1s)
	int2, _ := strconv.Atoi(int2s)
	cnt1, _ := strconv.Atoi(cnt1s)
	cnt2, _ := strconv.Atoi(cnt2s)

	if (int1 <= int2 && cnt1 < cnt2) || (int1 >= int2 && cnt1 > cnt2) {
		return "-1"
	}

	min1 := (stop+int1)*cnt1 - int1
	max1 := int1 + (stop+int1)*cnt1

	min2 := (stop+int2)*cnt2 - int2
	max2 := int2 + (stop+int2)*cnt2

	if (min1 > max2) || (max1 < min2) {
		return "-1"
	}

	totalMin := int(math.Max(float64(min1), float64(min2)))
	totalMax := int(math.Min(float64(max1), float64(max2)))

	return fmt.Sprintf("%d %d", totalMin, totalMax)
}
