package boilerplate

import (
	_ "embed"
)

//go:embed cmd-app-main.tpl
var CmdAppMainGo string

//go:embed go-mod.tpl
var GoMod string

//go:embed config.yaml.tpl
var ConfigExampleYaml string

//go:embed gitignore.tpl
var GitIgnore string

//go:embed migrations.go.tpl
var Migrations string

//go:embed api.tpl
var Api string

//go:embed ports-service.tpl
var PortsService string

//go:embed service.tpl
var Service string

//go:embed entity.tpl
var Entity string

//go:embed ports-repository.tpl
var PortsRepository string

//go:embed repository.tpl
var Repository string

//go:embed migrate.tpl
var Migrate string

//go:embed ports.tpl
var Ports string
