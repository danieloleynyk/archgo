package cmd

import (
	utils "archgo/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var updateAll bool
var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update linux package",
		Long:  `Manually update a linux package`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, packageName := range args {
				scriptName, err := utils.GetScriptName(packageName)
				utils.Catch(err, "Error while getting script name")

				err = utils.RunBashScript(scriptsPath + scriptName)
				utils.Catch(err, fmt.Sprintf("An error occurred in %s update", packageName))
			}

			if updateAll && len(args) == 0 {
				updateAllScriptName := "updateAll.sh"
				err := utils.RunBashScript(scriptsPath + updateAllScriptName)
				utils.Catch(err, "An error occurred in update all script")
			}
		},
	}
)

func init() {
	updateCmd.Flags().BoolVarP(&updateAll, "all", "a", false, "Updates all packages")
}
