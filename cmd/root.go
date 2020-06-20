package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Linux CLI Autoupdate",
	Short: "An auto updater for linux",
	Long: `An open source auto updater for the different linux distro built with go, 
complete documentation can be found on https://github.com/danieloleynyk/autoupdate-cli`,
}

// Execute function for root
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
