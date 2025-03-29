package rest

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	*gin.Context
}

type handlerFunc func(*Request)

type Api interface {
	Routes(s *Route)
}
