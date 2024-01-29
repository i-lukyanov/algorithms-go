package cmd

import (
	"algorithms/config"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string             //nolint:gochecknoglobals
	rootCmd = &cobra.Command{} //nolint:gochecknoglobals
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execution error", err.Error())
	}
}

func init() { //nolint:gochecknoinits
	cobra.OnInitialize(initConfig)

	// Init command
	addCommand()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.yaml", "path to config file")
}

func addCommand() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "exercise",
		Short: "Start exercise",
		Run:   runExercise,
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfg := config.GetConfig()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		log.Fatal("config file not provided")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("error reading config file [%s]", viper.ConfigFileUsed()))
	}

	// override default config
	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println(fmt.Sprintf("unable to unmarshal config [%s]", viper.ConfigFileUsed()))
	}
}
