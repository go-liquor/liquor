package templates

import (
	_ "embed"
)

//go:embed service.tmpl
var Service string

//go:embed route.tmpl
var Route string

//go:embed handler.tmpl
var Handler string

//go:embed repository.tmpl
var Repository string

//go:embed ports_repository.tmpl
var RepositoryPorts string

//go:embed migrate.tmpl
var Migrate string

//go:embed entity.tmpl
var Entity string

//go:embed grpc_proto.tmpl
var GrpcProto string

//go:embed grpc_server.tmpl
var GrpcServer string
