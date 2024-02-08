package pkg

import (
	"algorithms/config"
	"algorithms/pkg/common"
	"algorithms/pkg/go_tour"
	"algorithms/pkg/leetcode"
	"algorithms/pkg/yandex_alg10"
	"errors"
)

type Service struct {
	Args []string
	Cfg  *config.Config
}

func NewService(cfg *config.Config, args []string) *Service {
	return &Service{
		Args: args,
		Cfg:  cfg,
	}
}

func (s *Service) Run() error {
	if len(s.Args) == 0 {
		return errors.New("no package provided")
	}

	if len(s.Args) < 2 {
		s.Args = append(s.Args, "")
	}

	switch s.Args[0] {
	case common.PackageGoTour:
		return go_tour.RunExercise(s.Cfg, s.Args[1])
	case common.PackageLeetcode:
		return leetcode.RunExercise(s.Cfg, s.Args[1])
	case common.PackageYandexAlg10:
		return yandex_alg10.RunExercise(s.Cfg, s.Args[1])
	case common.PackageHackerrank:
	case common.PackageCodewars:
	default:
		return errors.New("wrong package name")
	}

	return nil
}
