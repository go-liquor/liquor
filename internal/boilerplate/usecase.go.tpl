package application

import (
    "go.uber.org/zap"
    "{{.module}}/internal/ports"
)

type {{.useCaseName}}Service struct {
    logger *zap.Logger
}

func New{{.useCaseName}}Service(logger *zap.Logger) ports.{{.useCaseName}}Service {
    return &{{.useCaseName}}Service{
        logger: logger,
    }
}

