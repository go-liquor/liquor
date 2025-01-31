package create

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

func createRouteCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "route",
		Short: "Create a new route group",
		RunE:  createRouteRun,
	}
	cmd.Flags().StringP("name", "n", "", "Route group name")
	cmd.Flags().StringP("group", "g", "", "Path to access route group (eg /api/users)")
	cmd.Flags().BoolP("crud", "c", false, "Create routes to CRUD")
	cmd.MarkFlagRequired("name")
	return &cmd
}

func createRouteRun(cmd *cobra.Command, args []string) error {
	var name, _ = cmd.Flags().GetString("name")
	var group, _ = cmd.Flags().GetString("group")
	var crud, _ = cmd.Flags().GetBool("crud")

	if group == "" {
		group = "/" + strings.ToLower(name)
	}

	modFile, err := commons.GetModFile(".")
	if err != nil {
		return err
	}

	var routeFilename string = textcase.SnakeCase(name) + ".go"
	var handlerFilename string = textcase.SnakeCase(name) + ".go"
	var outputFileRoute string = path.Join("internal/adapters/server/http/routes", routeFilename)
	var outputFileHandler string = path.Join("internal/adapters/server/http/handlers", handlerFilename)

	if _, err := os.Stat(outputFileRoute); err == nil {
		return fmt.Errorf("file %v already exists", outputFileRoute)
	}
	if _, err := os.Stat(outputFileHandler); err == nil {
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
	}); err != nil {
		return err
	}

	message.Success("created %v", outputFileRoute)
	message.Success("created %v", outputFileHandler)

	return nil
}
