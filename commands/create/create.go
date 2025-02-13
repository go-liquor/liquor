package create

import "github.com/spf13/cobra"

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create objects",
}

func init() {
	CreateCmd.AddCommand(
		createModuleCmd,
	)
}
