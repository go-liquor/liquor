package create

import (
	"fmt"
	"os"
	"path"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/project"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

func createMigrateCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "migrate",
		Short: "Create a new migrate",
		RunE:  createMigrateRun,
	}
	cmd.Flags().StringP("name", "n", "", "Migrate name")
	cmd.MarkFlagRequired("name")
	return &cmd
}

func createMigrateRun(cmd *cobra.Command, args []string) error {
	var name, _ = cmd.Flags().GetString("name")

	proj := project.GetProject()

	if proj.DatabaseDriver == "" {
		selected := commons.SelectDatabaseDriver()
		proj.DatabaseDriver = selected
		project.UpdateProject(proj)
	}

	var migrateFilename string = textcase.SnakeCase(name) + "_migrate.go"
	var outputMigrate string = path.Join("internal/adapters/database/migrations", migrateFilename)

	if _, err := os.Stat(outputMigrate); err == nil {
		return fmt.Errorf("file %v already exists", outputMigrate)
	}

	if err := templates.ParseTemplate(templates.Migrate, outputMigrate, map[string]any{
		"CamelCaseName":  textcase.CamelCase(name),
		"DatabaseDriver": proj.DatabaseDriver,
	}); err != nil {
		return err
	}

	if err := commons.ReplaceAnnotations("internal/adapters/database/migrations/migrations.go",
		"//go:liquor:migrations",
		fmt.Sprintf("\t%vMigrate,", textcase.CamelCase(name))); err != nil {
		return err
	}

	message.Success("updated migrations.go")

	return nil
}
