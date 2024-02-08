package leetcode

import (
	"algorithms/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

type RemoveNodeFromLinkedListEnd struct{}

func (e *RemoveNodeFromLinkedListEnd) removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	var listSlc []int
	for head != nil {
		listSlc = append(listSlc, head.Val)
		head = head.Next
	}

	ll := len(listSlc)
	if ll == 1 && n >= 1 {
		return nil
	}

	listSlc = append(listSlc[:ll-n], listSlc[ll-n+1:]...)
	ll = len(listSlc)

	var res = &common.ListNode{
		Val:  listSlc[ll-1],
		Next: nil,
	}
	for i := ll - 2; i >= 0; i-- {
		res = &common.ListNode{
			Val:  listSlc[i],
			Next: res,
		}
	}

	return res
}

//func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
//	cnt := 1
//	last := head
//
//	for {
//		if last.Next != nil {
//			cnt++
//			last = last.Next
//
//			continue
//		}
//
//		break
//	}
//
//	if head.Next == nil || n >= cnt {
//		return nil
//	}
//
//	nn := head
//	prev := head
//	for cnt >= (n + 1) {
//		prev = nn
//		nn = nn.Next
//		cnt--
//	}
//
//	prev.Next = nn.Next
//
//	return head
//}

func (e RemoveNodeFromLinkedListEnd) Run() {
	type llCase struct {
		list *common.ListNode
		num  int
	}

	llCases := []llCase{
		{
			list: &common.ListNode{
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
			num: 2,
		}, // 1-2-3-5
		{
			list: &common.ListNode{
				Val:  1,
				Next: nil,
			},
			num: 1,
		}, // []
		{
			list: &common.ListNode{
				Val: 1,
				Next: &common.ListNode{
					Val:  2,
					Next: nil,
				},
			},
			num: 1,
		}, // 1
		{
			list: &common.ListNode{
				Val: 1,
				Next: &common.ListNode{
					Val:  2,
					Next: nil,
				},
			},
			num: 2,
		}, // 2
	}

	for _, ex := range llCases {
		res := e.removeNthFromEnd(ex.list, ex.num)
		var resSlc []string
		for res != nil {
			resSlc = append(resSlc, strconv.Itoa(res.Val))
			res = res.Next
		}
		fmt.Println(fmt.Sprintf("----- result: %s", strings.Join(resSlc, "-")))
	}
}
