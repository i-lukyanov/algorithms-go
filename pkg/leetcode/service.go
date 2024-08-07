package leetcode

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"two_sum":                          TwoSum{},
	"palindrome_number":                PalindromeNumber{},
	"reverse_linked_list":              ReverseLinkedList{},
	"reverse_polish_notation":          ReversePolishNotation{},
	"daily_temp":                       DailyTemp{},
	"reverse_int":                      ReverseInt{},
	"longest_palindrome_substr":        LongestPalindromeSubstr{},
	"valid_parentheses":                ValidParentheses{},
	"remove_node_from_linked_list_end": RemoveNodeFromLinkedListEnd{},
	"divide_arr_with_max_diff":         DivideArrWithMaxDiff{},
}

func RunExercise(cfg *config.Config, name string) error {
	if name == "" {
		return errors.New("no exercise name provided")
	}

	exer, ok := exerciseMap[name]
	if !ok {
		return errors.New("wrong exercise name")
	}

	exer.Run()

	return nil
}
