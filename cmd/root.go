package cmd

import (
	"fmt"
	"os"

	utils "archgo/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "archgo",
		Short: "A tool for daily arch users problems",
		Long: `A command line tool, that solves some of the frequent problems 
that arch users deal with. documentation can be found on https://github.com/danieloleynyk/archgo`,
	}

	scriptsPath string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(updateCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/autoupdate-cli/")
	viper.AddConfigPath(".")

	utils.Catch(viper.ReadInConfig(), "Error reading config")

	scriptsPath = viper.GetString("scripts_path")
}

// Execute function for root
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
