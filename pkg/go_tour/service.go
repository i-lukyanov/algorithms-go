package go_tour

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"fibonacci_range_and_close": FibonacciRangeAndClose{},
	"equivalent_binary_trees":   EquivalentBinaryTrees{},
	"web_crawler":               WebCrawler{},
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
