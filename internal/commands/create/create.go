package create

import "github.com/spf13/cobra"

var rootPath string

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
		createRepository(),
		createRest(),
	)
	cm.PersistentFlags().StringVarP(&rootPath, "root", "r", ".", "Root path")
	return cm
}
