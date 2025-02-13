package create

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-liquor/liquor/internal/commons"
	"github.com/go-liquor/liquor/internal/message"
	"github.com/go-liquor/liquor/internal/templates"
	"github.com/golang-cz/textcase"
	"github.com/spf13/cobra"
)

var createGrpcService = &cobra.Command{
	Use:   "grpc",
	Short: "Create GRPC service",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, _ = cmd.Flags().GetString("name")
		var force, _ = cmd.Flags().GetBool("force")
		return createGRPCServiceRun(name, force)
	},
}

func createGRPCServiceRun(name string, force bool) error {

	if _, err := exec.LookPath("protoc"); err != nil {
		return fmt.Errorf("protoc command not found. Please install Protocol Buffers compile tool. See how in https://grpc.io/docs/languages/go/quickstart/")
	}

	modFile, err := commons.GetModFile(".")
	if err != nil {
		return err
	}

	os.MkdirAll("pkg/proto", 0755)
	os.MkdirAll("internal/adapters/server/grpc", 0755)

	var protoFile = path.Join("pkg/proto", textcase.SnakeCase(name)+".proto")
	var serverFile = path.Join("internal/adapters/server/grpc", "server.go")

	if !force {
		if commons.IsExist(protoFile) {
			return fmt.Errorf("file %v already exists", protoFile)
		}

		if commons.IsExist(serverFile) {
			return fmt.Errorf("file %v alread exists", serverFile)
		}
	}

	files := map[string]string{
		protoFile:  templates.GrpcProto,
		serverFile: templates.GrpcServer,
	}

	if err := templates.ParseTemplates(files, map[string]any{
		"PascalCaseName": textcase.PascalCase(name),
		"Package":        modFile.Module.Mod.Path,
	}); err != nil {
		return err
	}

	message.Success("created %v", protoFile)
	message.Success("created %v", serverFile)
	commons.Command(".", "make", "protogen")

	commons.PrintCode(`
	// add to cmd/app/main.go
	//
	// # imports
	import (
		"google.golang.org/grpc"
		"` + modFile.Module.Mod.Path + `/pkg/proto"
		grpcadapter "` + modFile.Module.Mod.Path + `/internal/adapters/server/grpc"
		liquorgrpc "github.com/go-liquor/liquor-sdk/server/grpc"
	)
		// Code
	liquorgrpc.RegisterGRPCServer(&grpcadapter.Server{},
		grpcadapter.NewServer,
		func(adapter *grpcadapter.Server, svc *grpc.Server) {
			proto.Register` + textcase.PascalCase(name) + `Server(svc, adapter)
		}),
	`)

	return nil
}

func init() {
	createGrpcService.Flags().StringP("name", "n", "", "grpc service name")
	createGrpcService.Flags().Bool("force", false, "force to create files")
	createGrpcService.MarkFlagRequired("name")
}
