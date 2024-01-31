package leetcode

import (
	"algorithms/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

type ReverseLinkedList struct{}

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := &common.ListNode{
		Val: head.Val,
	}

	for head.Next != nil {
		head = head.Next

		tail = &common.ListNode{
			Val:  head.Val,
			Next: tail,
		}
	}

	return tail
}

var linkedListCases = []*common.ListNode{
	{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val: 4,
					Next: &common.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	},
	{
		Val: 1,
		Next: &common.ListNode{
			Val:  2,
			Next: nil,
		},
	},
	nil,
	{
		Val:  1,
		Next: nil,
	},
}

func (e ReverseLinkedList) Run() {
	for _, ex := range linkedListCases {
		res := reverseList(ex)
		var resSlc []string
		for res != nil {
			resSlc = append(resSlc, strconv.Itoa(res.Val))
			res = res.Next
		}
		fmt.Println(fmt.Sprintf("----- result: %s", strings.Join(resSlc, " - ")))
	}
}
