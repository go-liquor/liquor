package create

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/project"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

var createModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Create module in app/<module-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		return createModuleRun(name, "")
	},
}

func createModuleRun(name string, routeGroup string) error {
	modulePath := path.Join("app", name)

	restServerPath := path.Join(modulePath, "adapters/server/rest")
	databasePath := path.Join(modulePath, "adapters/database")
	repositoriesPath := path.Join(databasePath, "repositories")
	migrationsPath := path.Join(databasePath, "migrations")
	entityPath := path.Join(modulePath, "domain/entity")
	portsPath := path.Join(modulePath, "domain/ports")
	servicePath := path.Join(modulePath, "services")

	proj := project.GetProject()

	if proj.DatabaseDriver == "" {
		selected := commons.SelectDatabaseDriver()
		proj.DatabaseDriver = selected
		project.UpdateProject(proj)
	}

	modFile, err := commons.GetModFile(".")
	if err != nil {
		return err
	}

	if commons.IsExist(modulePath) {
		return fmt.Errorf("module %s already exists", name)
	}

	paths := [7]string{
		modulePath,
		restServerPath,
		repositoriesPath,
		migrationsPath,
		entityPath,
		portsPath,
		servicePath,
	}

	for _, p := range paths {
		if err := os.MkdirAll(p, 0755); err != nil {
			return err
		}
		message.Success("created %s", p)
	}

	if routeGroup == "" {
		routeGroup = fmt.Sprintf("/%s", name)
	}

	moduleApplicationFileOutput := path.Join(modulePath, commons.ToFilename(name, "_app"))
	serviceOutput := path.Join(servicePath, commons.ToFilename(name))
	repositoryOutput := path.Join(repositoriesPath, commons.ToFilename(name, "_repository"))
	repositoryPortsOutput := path.Join(portsPath, commons.ToFilename(name, "_repository"))
	migrationOutput := path.Join(migrationsPath, commons.ToFilename(name, "_migration"))
	databaseModuleOutput := path.Join(databasePath, "module.go")
	routeOutput := path.Join(restServerPath, commons.ToFilename(name, "_route"))
	handlerOutput := path.Join(restServerPath, commons.ToFilename(name, "_handler"))
	restModuleOutput := path.Join(restServerPath, "module.go")
	entityOutput := path.Join(entityPath, commons.ToFilename(name))

	files := map[string]string{
		moduleApplicationFileOutput: templates.ModuleApp,
		serviceOutput:               templates.Service,
		repositoryOutput:            templates.Repository,
		repositoryPortsOutput:       templates.RepositoryPorts,
		routeOutput:                 templates.Route,
		handlerOutput:               templates.Handler,
		restModuleOutput:            templates.RestModule,
		entityOutput:                templates.Entity,
		migrationOutput:             templates.Migrate,
		databaseModuleOutput:        templates.DatabaseModule,
	}

	if err := templates.ParseTemplates(files, map[string]string{
		"ModuleName":     strings.ToLower(name),
		"PascalCaseName": textcase.PascalCase(name),
		"Package":        modFile.Module.Mod.Path,
		"DatabaseDriver": proj.DatabaseDriver,
		"Group":          routeGroup,
	}); err != nil {
		return err
	}

	message.Info("Update you cmd/app/main.go with this:")

	commons.PrintCode(`
package main

import (
	"` + modFile.Module.Mod.Path + `/app/` + strings.ToLower(name) + `" // add this

	"github.com/go-liquor/liquor-sdk/app"
)

func main() {
	app.NewApp(
  		` + strings.ToLower(name) + `.Module, // add this
	)
}`)

	return nil
}

func init() {
	createModuleCmd.Flags().StringP("name", "n", "", "module name")
	createModuleCmd.MarkFlagRequired("name")
}
