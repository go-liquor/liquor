package templates

import (
	_ "embed"
)

//go:embed app/module_app.tmpl
var ModuleApp string

//go:embed service/service.tmpl
var Service string

//go:embed rest/route.tmpl
var Route string

//go:embed rest/rest_module.tmpl
var RestModule string

//go:embed rest/handler.tmpl
var Handler string

//go:embed database/repository.tmpl
var Repository string

//go:embed database/ports_repository.tmpl
var RepositoryPorts string

//go:embed database/migrate.tmpl
var Migrate string

//go:embed database/module.tmpl
var DatabaseModule string

//go:embed entity/entity.tmpl
var Entity string

//go:embed grpc/grpc_proto.tmpl
var GrpcProto string

//go:embed grpc/grpc_server.tmpl
var GrpcServer string
