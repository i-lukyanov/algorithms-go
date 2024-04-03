package yandex_alg101

import (
	"algorithms/pkg/common"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	cCode  = "495"
	pLen11 = 11
)

type CPhoneNumbers struct{}

func (e CPhoneNumbers) Run() {
	rd := bufio.NewReader(os.Stdin)

	input, _ := rd.ReadString('\n')

	s1, _ := rd.ReadString('\n')
	s2, _ := rd.ReadString('\n')
	s3, _ := rd.ReadString('\n')

	input = strings.Replace(input, "\n", "", -1)
	s1 = strings.Replace(s1, "\n", "", -1)
	s2 = strings.Replace(s2, "\n", "", -1)
	s3 = strings.Replace(s3, "\n", "", -1)

	res := make([]string, 3)
	res = append(
		res,
		e.match(input, s1),
		e.match(input, s2),
		e.match(input, s3),
	)

	for _, r := range res {
		fmt.Println(r)
	}
}

func (e CPhoneNumbers) match(phone1, phone2 string) string {
	re := regexp.MustCompile(`\D+`)
	p1 := re.ReplaceAllString(phone1, "")
	p2 := re.ReplaceAllString(phone2, "")

	if len(p1) < pLen11 {
		p1 = cCode + p1
	} else {
		p1 = p1[1:]
	}

	if len(p2) < pLen11 {
		p2 = cCode + p2
	} else {
		p2 = p2[1:]
	}

	if p1 == p2 {
		return common.Yes
	}

	return common.No
}
