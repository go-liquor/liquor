package create

import (
	"fmt"
	"path"
	"strings"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

var createServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a new service (internal/app/services)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")

		name = strings.ReplaceAll(name, "service", "")
		name = strings.ReplaceAll(name, "Service", "")

		if name == "" {
			return fmt.Errorf("the name %v can't have 'Service' in name. We already put", name)
		}

		var filename string = commons.ToFilename(name)
		var outputFile string = path.Join("internal/app/services", filename)

		if commons.IsExist(outputFile) {
			return fmt.Errorf("file %v already exists", outputFile)
		}

		if err := templates.ParseTemplate(templates.Service, outputFile, map[string]string{
			"PascalCaseName": textcase.PascalCase(name),
		}); err != nil {
			return err
		}
		message.Success("created %v", outputFile)
		fmt.Println("You need register service in cmd/app/main.go:")
		commons.PrintCode(fmt.Sprintf(`
func main() {
	app.NewApp(
		http.Server,
		app.RegisterService(
			services.NewInitialService,
			// here
			services.New%vService
		),
	)
}
	`, textcase.PascalCase(name)))
		return nil
	},
}

func init() {
	createServiceCmd.Flags().StringP("name", "n", "", "Service name")
	createServiceCmd.MarkFlagRequired("name")
}
