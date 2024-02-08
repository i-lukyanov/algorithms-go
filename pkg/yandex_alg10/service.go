package yandex_alg10

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"a_conditioner": AConditioner{},
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
