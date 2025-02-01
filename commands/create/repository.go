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

func createRepositoryCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "repository",
		Short: "Create a new repository",
		RunE:  createRepositoryRun,
	}
	cmd.Flags().StringP("name", "n", "", "Repository name")
	cmd.MarkFlagRequired("name")
	return &cmd
}

func createRepositoryRun(cmd *cobra.Command, args []string) error {
	var name, _ = cmd.Flags().GetString("name")

	name = strings.ReplaceAll(name, "repository", "")
	name = strings.ReplaceAll(name, "Repository", "")

	if name == "" {
		return fmt.Errorf("the name can't have 'Repository' in name. We already put")
	}

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

	os.MkdirAll("internal/adapters/database/repository", 0755)
	os.MkdirAll("internal/app/ports", 0755)

	var repoFilename string = textcase.SnakeCase(name) + "_repository.go"
	var portsFilename string = textcase.SnakeCase(name) + "_repository.go"
	var outputFileRepo string = path.Join("internal/adapters/database/repository", repoFilename)
	var outputFilePort string = path.Join("internal/app/ports", portsFilename)

	if _, err := os.Stat(outputFileRepo); err == nil {
		return fmt.Errorf("file %v already exists", outputFileRepo)
	}
	if _, err := os.Stat(outputFilePort); err == nil {
		return fmt.Errorf("file %v already exists", outputFilePort)
	}

	files := map[string]string{
		outputFileRepo: templates.Repository,
		outputFilePort: templates.RepositoryPorts,
	}

	if err := templates.ParseTemplates(files, map[string]any{
		"PascalCaseName": textcase.PascalCase(name),
		"Package":        modFile.Module.Mod.Path,
		"DatabaseDriver": proj.DatabaseDriver,
	}); err != nil {
		return err
	}

	message.Success("created %v", outputFileRepo)
	message.Success("created %v", outputFilePort)

	if err := commons.ReplaceAnnotations("cmd/app/main.go",
		"//go:liquor:repositories",
		fmt.Sprintf("\trepository.New%vRepository,", textcase.PascalCase(name))); err != nil {
		return err
	}

	return nil
}
