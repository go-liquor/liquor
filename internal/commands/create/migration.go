package create

import (
	"path"
	"time"

	"github.com/go-liquor/liquor/v2/internal/boilerplate"
	"github.com/go-liquor/liquor/v2/internal/stdout"
	"github.com/go-liquor/liquor/v2/internal/templates"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/spf13/cobra"
)

const (
	formatMigration = "20060102150405"
)

var createMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Create a new migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		if err := templates.ParseTemplate(boilerplate.Migrate, path.Join("migrations",
			time.Now().Format(formatMigration)+"_"+lqstring.ToSnakeCase(name)+".go"),
			map[string]any{
				"MigrateName": lqstring.ToPascalCase(name),
			}); err != nil {
			return err
		}
		stdout.Success("Migration created successfully")
		return nil
	},
}

func init() {
	createMigrationCmd.Flags().StringP("name", "n", "", "Migration name")
	createMigrationCmd.MarkFlagRequired("name")
	CreateCmd.AddCommand(createMigrationCmd)
}
