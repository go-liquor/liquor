package create

import "github.com/spf13/cobra"

func CreateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Command to create objects",
	}
	cmd.AddCommand(
		createServiceCmd(),
		createRouteCmd(),
		createRepositoryCmd(),
		createMigrateCmd(),
	)
	return &cmd
}
