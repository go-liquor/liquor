package create

import (
	"github.com/spf13/cobra"
)

var createResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create a new resource (routes, handlers, entity, service, repository)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		createEntityRun(name)
		createRouteRun(name, "", true, true)
		createRepositoryRun(name)
		createServiceRun(name, true)
		return nil
	},
}

func init() {
	createResourceCmd.Flags().StringP("name", "n", "", "Resource name")
	createResourceCmd.MarkFlagRequired("name")
}
