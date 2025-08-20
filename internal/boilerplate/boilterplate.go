package boilerplate

import (
	"embed"
	_ "embed"
)

//go:embed project/*
var ProjectFiles embed.FS

//go:embed domain.go.tpl
var DomainFile string

//go:embed model.go.tpl
var ModelFile string

//go:embed usecase.go.tpl
var UsecaseFile string

//go:embed usecase_port.go.tpl
var UsecasePortFile string

//go:embed repository_port.go.tpl
var RepositoryPortFile string

//go:embed repository/*.tpl
var RepositoryImplFiles embed.FS

//go:embed rest.go.tpl
var RestApiFile string
