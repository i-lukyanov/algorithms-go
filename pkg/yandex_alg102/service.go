package yandex_alg102

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"a_is_increase":   AIsIncrease{},
	"b_sequence_type": BSequenceType{},
	"c_closest_num":   CClosestNum{},
	"d_bigger":        DBigger{},
	"e_cow":           ECow{},
	"f_symmetr":       FSymmetr{},
	"g_2biggest":      G2Biggest{},
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
