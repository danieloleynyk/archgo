package cmd

import (
	"fmt"
	"os"

	"../utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goupdate",
		Short: "An auto updater for linux",
		Long: `An open source auto updater for the different linux distro built with go, 
complete documentation can be found on https://github.com/danieloleynyk/autoupdate-cli`,
	}

	distribution string
	scriptsPath  string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&distribution, "distro", "", "distribution name [arch, ubuntu, centos]")
	rootCmd.AddCommand(updateCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/autoupdate-cli/")
	viper.AddConfigPath(".")

	utils.Catch(viper.ReadInConfig(), "Error reading config")

	distribution = viper.GetString("distribution")
	scriptsPath = viper.GetString("scripts_path")
}

// Execute function for root
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
