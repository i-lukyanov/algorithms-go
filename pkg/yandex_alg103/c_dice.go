package yandex_alg103

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type CDice struct{}

func (e CDice) Run() {
	rd := bufio.NewReader(os.Stdin)

	countsStr, _ := rd.ReadString('\n')
	countsStr = strings.Replace(countsStr, "\n", "", -1)

	counts := strings.Split(strings.TrimSpace(countsStr), " ")

	cnt1, _ := strconv.Atoi(counts[0])
	cnt2, _ := strconv.Atoi(counts[1])

	s1 := common.NewSetInt()
	s2 := common.NewSetInt()

	for i := 0; i < cnt1; i++ {
		nums1Str, _ := rd.ReadString('\n')
		nums1Str = strings.Replace(nums1Str, "\n", "", -1)
		num1, _ := strconv.Atoi(nums1Str)
		s1.Add(num1)
	}

	for j := 0; j < cnt2; j++ {
		nums2Str, _ := rd.ReadString('\n')
		nums2Str = strings.Replace(nums2Str, "\n", "", -1)
		num2, _ := strconv.Atoi(nums2Str)
		s2.Add(num2)
	}

	inter := s1.Intersect(s2).ToSlice()
	slices.Sort(inter)

	s1Diff := s1.Diff(s2).ToSlice()
	slices.Sort(s1Diff)

	s2Diff := s2.Diff(s1).ToSlice()
	slices.Sort(s2Diff)

	interStr := make([]string, 0)
	s1DiffStr := make([]string, 0)
	s2DiffStr := make([]string, 0)

	for i := 0; i < len(inter); i++ {
		interStr = append(interStr, fmt.Sprintf("%d", inter[i]))
	}

	for j1 := 0; j1 < len(s1Diff); j1++ {
		s1DiffStr = append(s1DiffStr, fmt.Sprintf("%d", s1Diff[j1]))
	}

	for j2 := 0; j2 < len(s2Diff); j2++ {
		s2DiffStr = append(s2DiffStr, fmt.Sprintf("%d", s2Diff[j2]))
	}

	fmt.Println(fmt.Sprintf("%d", len(interStr)))

	fmt.Println(strings.Join(interStr, " "))

	fmt.Println(fmt.Sprintf("%d", len(s1DiffStr)))

	fmt.Println(strings.Join(s1DiffStr, " "))

	fmt.Println(fmt.Sprintf("%d", len(s2DiffStr)))

	fmt.Println(strings.Join(s2DiffStr, " "))
}
