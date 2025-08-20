package create

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/go-liquor/liquor/v3/internal/boilerplate"
	"github.com/go-liquor/liquor/v3/internal/execcm"
	"github.com/go-liquor/liquor/v3/internal/stdout"
	"github.com/go-liquor/liquor/v3/internal/templates"
	"github.com/spf13/cobra"
)

func createProjectCmd() *cobra.Command {
	var projectName string
	var module string
	cm := &cobra.Command{
		Use:   "project [name]",
		Short: "Create a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				projectName = args[0]
			}

			var fields []huh.Field = []huh.Field{}
			if projectName == "" {
				fields = append(fields, huh.NewInput().
					Title("set the project name").
					Value(&projectName))
			}
			if module == "" {
				fields = append(fields, huh.NewInput().
					Title("set the go mod name").
					Placeholder("github.com/example/my-project").
					Value(&module),
				)
			}
			if len(fields) > 0 {
				form := huh.NewForm(
					huh.NewGroup(
						fields...,
					),
				)
				if err := form.Run(); err != nil {
					return err
				}
			}

			if err := os.MkdirAll(projectName, 0755); err != nil {
				return err
			}

			dirs := []string{
				"project",
			}
			dstRoot := path.Join(rootPath, projectName)

			for i := 0; i < len(dirs); i++ {
				root := dirs[i]
				files, err := boilerplate.ProjectFiles.ReadDir(root)
				if err != nil {
					return err
				}
				for _, f := range files {
					if f.IsDir() {
						nDir := path.Join(root, f.Name())
						dirs = append(dirs, nDir)
						os.MkdirAll(path.Join(dstRoot, strings.Replace(nDir, "project/", "", 1)), 0755)
						continue
					}
					if f.Name() == "gitkeep.tpl" {
						continue
					}
					name := strings.Replace(f.Name(), ".tpl", "", 1)
					dstFile := path.Join(dstRoot, strings.Replace(root, "project", "", 1), name)
					templateFile := path.Join(root, f.Name())
					contentOriginalFile, err := boilerplate.ProjectFiles.ReadFile(templateFile)
					if err != nil {
						return fmt.Errorf("failed to find template file %s: %v", templateFile, err)
					}
					if err := templates.ParseTemplate(string(contentOriginalFile), dstFile, map[string]any{
						"name":      projectName,
						"module":    module,
						"goversion": strings.Replace(runtime.Version(), "go", "", 1),
					}); err != nil {
						return err
					}

				}
			}

			if err := execcm.Command(dstRoot, "go", "get", "github.com/go-liquor/liquor/v3@latest"); err != nil {
				return err
			}
			if err := execcm.Command(dstRoot, "go", "mod", "tidy"); err != nil {
				return err
			}

			stdout.Success("created project %v", projectName)
			return nil
		},
	}
	cm.Flags().StringVarP(&projectName, "name", "n", "", "Project name")
	cm.Flags().StringVarP(&module, "module", "m", "", "Module name (eg: github.com/myuser/myproject)")
	return cm
}
