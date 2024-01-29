package cmd

import (
	"algorithms/config"
	"algorithms/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

func runExercise(_ *cobra.Command, args []string) {
	cfg := config.GetConfig()

	fmt.Println("--- start exercise ---")

	exer := pkg.NewService(cfg, args)
	if err := exer.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
	}

	fmt.Println("--- finish exercise ---")
}
