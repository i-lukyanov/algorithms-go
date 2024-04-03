package yandex_alg101

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"errors"
)

var exerciseMap = map[string]common.Exercise{
	"a_conditioner":   AConditioner{},
	"b_triangle":      BTriangle{},
	"c_phone_numbers": CPhoneNumbers{},
	"d_sqrt_equasion": DSqrtEquasion{},
	"e_ambulance":     EAmbulance{},
	"f_notebooks":     FNotebooks{},
	"g_details":       GDetails{},
	"h_metro":         HMetro{},
	"i_if":            IIf{},
	"j_system":        JSystem{},
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
