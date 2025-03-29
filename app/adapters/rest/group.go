package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteGroup struct {
	group *gin.RouterGroup
}

func (rg *RouteGroup) Handle(method string, path string, handler handlerFunc) {
	rg.group.Handle(method, path, func(ctx *gin.Context) {
		handler(&Request{Context: ctx})
	})
}

func (rg *RouteGroup) Middleware(middleware ...gin.HandlerFunc) *RouteGroup {
	rg.group.Use(middleware...)
	return rg
}

// Get registers a new handler for GET HTTP method with the given path.
// It's a shorthand for Handle(http.MethodGet, path, handler).
func (rg *RouteGroup) Get(path string, handler handlerFunc) {
	rg.Handle(http.MethodGet, path, handler)
}

// Post registers a new handler for POST HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPost, path, handler).
func (rg *RouteGroup) Post(path string, handler handlerFunc) {
	rg.Handle(http.MethodPost, path, handler)
}

// Patch registers a new handler for PATCH HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPatch, path, handler).
func (rg *RouteGroup) Patch(path string, handler handlerFunc) {
	rg.Handle(http.MethodPatch, path, handler)
}

// Delete registers a new handler for DELETE HTTP method with the given path.
// It's a shorthand for Handle(http.MethodDelete, path, handler).
func (rg *RouteGroup) Delete(path string, handler handlerFunc) {
	rg.Handle(http.MethodDelete, path, handler)
}

// Put registers a new handler for PUT HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPut, path, handler).
func (rg *RouteGroup) Put(path string, handler handlerFunc) {
	rg.Handle(http.MethodPut, path, handler)
}
