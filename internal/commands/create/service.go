package create

import (
	"fmt"
	"path"

	"github.com/go-liquor/liquor/v2/internal/boilerplate"
	"github.com/go-liquor/liquor/v2/internal/gomod"
	"github.com/go-liquor/liquor/v2/internal/stdout"
	"github.com/go-liquor/liquor/v2/internal/templates"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/spf13/cobra"
)

var createServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a new service",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")

		gm, err := gomod.GetModFile(".")
		if err != nil {
			return fmt.Errorf("failed to get go.mod file: %v", err)
		}

		files := map[string]string{
			path.Join("app/services", lqstring.ToSnakeCase(name)+"_service.go"): boilerplate.Service,
			path.Join("app/ports", lqstring.ToSnakeCase(name)+"_service.go"):    boilerplate.PortsService,
		}

		if err := templates.ParseTemplates(files, map[string]any{
			"PascalCaseName": lqstring.ToPascalCase(name),
			"Package":        gm.Module.Mod.Path,
		}); err != nil {
			return err
		}

		stdout.Success("service created successfully")
		return nil
	},
}

func init() {
	createServiceCmd.Flags().StringP("name", "n", "", "Service name")
	createServiceCmd.MarkFlagRequired("name")
	CreateCmd.AddCommand(createServiceCmd)
}
