package yandex_alg103

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"a_different_numbers": ADifferentNumbers{},
	"b_intersect":         BIntersect{},
	"c_dice":              CDice{},
	"d_words":             DWords{},
	"e_opencalc":          EOpenCalc{},
	"f_alien":             FAlien{},
	"g_turtles":           GTurtles{},
	"h_pigs":              HPigs{},
	"i_lang":              ILang{},
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
