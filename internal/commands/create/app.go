package create

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-liquor/liquor/v2/internal/boilerplate"
	"github.com/go-liquor/liquor/v2/internal/constants"
	"github.com/go-liquor/liquor/v2/internal/execcm"
	"github.com/go-liquor/liquor/v2/internal/gomod"
	"github.com/go-liquor/liquor/v2/internal/models/choice"
	"github.com/go-liquor/liquor/v2/internal/models/input"
	"github.com/go-liquor/liquor/v2/internal/stdout"
	"github.com/go-liquor/liquor/v2/internal/templates"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"github.com/spf13/cobra"
)

var createAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Create a new application",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		var pkg, _ = cmd.Flags().GetString("pkg")
		var database, _ = cmd.Flags().GetString("database")
		var err error
		if name == "" {
			name, err = input.NewInput("choice the application name", "")
			if err != nil {
				return err
			}
		}
		if pkg == "" {
			pkg, err = input.NewInput("choice the application package", "e.g github.com/"+name)
			if err != nil {
				return err
			}
		}

		if database == "" {
			database, err = choice.NewChoice("select the database driver", []choice.ChoiceOption{
				{
					Title: "None",
					Value: constants.None,
				},
				{
					Title: "Postgres",
					Value: constants.Postgres,
				},
				{
					Title: "MySQL",
					Value: constants.MySQL,
				},
				{
					Title: "SQLite",
					Value: constants.SQLite,
				},
				{
					Title: "MongoDB",
					Value: constants.MongoDB,
				},
			})
		}

		if name == "" || pkg == "" || database == "" {
			return fmt.Errorf("name, pkg and database are required")
		}

		projectPath := name
		appPath := path.Join(projectPath, "app")
		adaptersPath := path.Join(appPath, "adapters")
		databasePath := path.Join(adaptersPath, "database")
		restPath := path.Join(adaptersPath, "rest")
		domainPath := path.Join(appPath, "domain")
		entityPath := path.Join(domainPath, "entity")
		portsPath := path.Join(appPath, "ports")
		servicesPath := path.Join(appPath, "services")
		cmdPath := path.Join(projectPath, "cmd")
		appCmdPath := path.Join(cmdPath, "app")
		migrationsPath := path.Join(projectPath, "migrations")

		pathsToCreate := []string{
			projectPath,
			appPath,
			adaptersPath,
			databasePath,
			restPath,
			domainPath,
			entityPath,
			portsPath,
			servicesPath,
			cmdPath,
			appCmdPath,
			migrationsPath,
		}

		for _, p := range pathsToCreate {
			if err := os.Mkdir(p, 0755); err != nil {
				stdout.Error("failed to create %v", p)
				continue
			}
			stdout.Success("created %v", p)
		}

		files := map[string]string{
			path.Join(appCmdPath, "main.go"):                       boilerplate.CmdAppMainGo,
			path.Join(projectPath, "go.mod"):                       boilerplate.GoMod,
			path.Join(projectPath, "config.example.yaml"):          boilerplate.ConfigExampleYaml,
			path.Join(projectPath, ".gitignore"):                   boilerplate.GitIgnore,
			path.Join(restPath, "api.go"):                          boilerplate.Api,
			path.Join(portsPath, "service.go"):                     boilerplate.PortsService,
			path.Join(servicesPath, "service.go"):                  boilerplate.Service,
			path.Join(entityPath, lqstring.ToSingular(name)+".go"): boilerplate.Entity,
		}

		if database != "none" {
			if database != "mongodb" {
				files[path.Join(migrationsPath, "migrations.go")] = boilerplate.Migrations
				files[path.Join(migrationsPath, time.Now().Format(formatMigration)+"_create_table.go")] = boilerplate.Migrate
			}
			files[path.Join(portsPath, lqstring.ToSingular(name)+"_repository.go")] = boilerplate.PortsRepository
			files[path.Join(databasePath, lqstring.ToSingular(name)+"_repository.go")] = boilerplate.Repository
		}

		templates.ParseTemplates(files, map[string]any{
			"Package":                pkg,
			"Name":                   name,
			"Database":               database,
			"PascalCaseName":         lqstring.ToPascalCase(name),
			"SingularPascalCaseName": lqstring.ToPascalCase(lqstring.ToSingular(name)),
			"MigrateName":            "CreateTable",
			"NewProject":             true,
		})

		mod, err := gomod.GetModFile(name)
		if err != nil {
			return fmt.Errorf("failed to get mod file: %v", err)
		}

		mod.DropRequire("github.com/go-liquor/liquor/v2")
		mod.AddRequire("github.com/go-liquor/liquor/v2", "latest")
		mod.DropReplace("github.com/go-liquor/liquor/v2", "")
		content, err := mod.Format()
		if err != nil {
			return fmt.Errorf("failed in format mod file: %v", err)
		}
		if err := os.WriteFile(path.Join(name, "go.mod"), content, 0755); err != nil {
			return fmt.Errorf("failed to recreate go.mod: %v", err)
		}

		execcm.Command(projectPath, "go", "mod", "tidy")

		return nil
	},
}

func init() {
	createAppCmd.Flags().StringP("name", "n", "", "application name")
	createAppCmd.Flags().StringP("pkg", "p", "", "application pkg (e.g. github.com/lbernardo/example)")
	createAppCmd.Flags().StringP("database", "d", "", "database driver (e.g mysql, postgres, sqlite)")
	CreateCmd.AddCommand(
		createAppCmd,
	)
}
