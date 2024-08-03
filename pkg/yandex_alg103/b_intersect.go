package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type BIntersect struct{}

func (e BIntersect) Run() {
	rd := bufio.NewReader(os.Stdin)

	nums1Str, _ := rd.ReadString('\n')
	nums2Str, _ := rd.ReadString('\n')

	nums1Str = strings.Replace(nums1Str, "\n", "", -1)
	nums2Str = strings.Replace(nums2Str, "\n", "", -1)

	nums1 := strings.Split(strings.TrimSpace(nums1Str), " ")
	nums2 := strings.Split(strings.TrimSpace(nums2Str), " ")

	s1 := common.NewSetInt()
	s2 := common.NewSetInt()

	for k1 := range nums1 {
		num1, _ := strconv.Atoi(nums1[k1])
		s1.Add(num1)
	}

	for k2 := range nums2 {
		num2, _ := strconv.Atoi(nums2[k2])
		s2.Add(num2)
	}

	inter := s1.Intersect(s2).ToMap()
	interStr := make([]string, 0)

	interMin := math.MaxInt
	interMax := math.MinInt

	for _, i := range inter {
		if i < interMin {
			interMin = i
		}
		if i > interMax {
			interMax = i
		}
	}

	for j := interMin; j <= interMax; j++ {
		if elem, ok := inter[j]; ok {
			interStr = append(interStr, fmt.Sprintf("%d", elem))
		}
	}

	fmt.Println(strings.Join(interStr, " "))
}
