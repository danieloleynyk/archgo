package cmd

import (
	utils "archgo/utils"
	"errors"

	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update linux package",
		Long:  `Manually update a linux package`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("update requires a package name")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			packageName = args[0]
			scriptName, err := utils.GetScriptName(packageName)
			utils.Catch(err, "Error while getting script name")

			utils.RunBashScript(scriptsPath + scriptName)
		},
	}

	packageName string
)
