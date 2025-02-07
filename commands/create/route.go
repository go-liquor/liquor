package create

import (
	"fmt"
	"path"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

var createRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Create a new route group",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		var group, _ = cmd.Flags().GetString("group")
		var crud, _ = cmd.Flags().GetBool("crud")
		return createRouteRun(name, group, crud, false)
	},
}

func createRouteRun(name string, group string, crud bool, byResource bool) error {
	if group == "" {
		group = "/" + textcase.KebabCase(name)
	}

	modFile, err := commons.GetModFile(".")
	if err != nil {
		return err
	}

	var routeFilename string = commons.ToFilename(name)
	var handlerFilename string = commons.ToFilename(name)
	var outputFileRoute string = path.Join("internal/adapters/server/http/routes", routeFilename)
	var outputFileHandler string = path.Join("internal/adapters/server/http/handlers", handlerFilename)

	if commons.IsExist(outputFileRoute) {
		return fmt.Errorf("file %v already exists", outputFileRoute)
	}
	if commons.IsExist(outputFileHandler) {
		return fmt.Errorf("file %v already exists", outputFileHandler)
	}

	files := map[string]string{
		outputFileRoute:   templates.Route,
		outputFileHandler: templates.Handler,
	}

	if err := templates.ParseTemplates(files, map[string]any{
		"PascalCaseName": textcase.PascalCase(name),
		"Package":        modFile.Module.Mod.Path,
		"Group":          group,
		"CRUD":           crud,
		"ByResource":     byResource,
	}); err != nil {
		return err
	}

	message.Success("created %v", outputFileRoute)
	message.Success("created %v", outputFileHandler)
	return nil
}

func init() {
	createRouteCmd.Flags().StringP("name", "n", "", "Route group name")
	createRouteCmd.Flags().StringP("group", "g", "", "Path to access route group (eg /api/users)")
	createRouteCmd.Flags().BoolP("crud", "c", false, "Create routes to CRUD")
	createRouteCmd.MarkFlagRequired("name")
}
