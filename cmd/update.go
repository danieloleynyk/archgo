package cmd

import (
	"../utils"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update linux package",
		Long:  `Manually update a linux package`,
		Run: func(cmd *cobra.Command, args []string) {
			scriptName, err := utils.GetScriptName(packageName, distribution)
			utils.Catch(err, "Error while getting script name")

			utils.RunBashScript(scriptsPath + scriptName)
		},
	}

	packageName string
)

func init() {
	packageNameFlag := "package"
	updateCmd.Flags().StringVar(&packageName, packageNameFlag, "", "Name of the package to update (required)")
	updateCmd.MarkFlagRequired(packageNameFlag)
}
