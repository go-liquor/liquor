package create

import "github.com/spf13/cobra"

func CreateCmd() *cobra.Command {
	cm := &cobra.Command{
		Use:   "create",
		Short: "create command",
	}
	cm.AddCommand(
		createProjectCmd(),
		createDomain(),
		createModel(),
		createUseCase(),
	)
	return cm
}
