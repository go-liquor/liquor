package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	*gin.Context
}

type handlerFunc func(*Request)

type Api interface {
	Routes(s *Server)
}

type Server struct {
	engine *gin.Engine
}

func newServer(e *gin.Engine) *Server {
	return &Server{
		engine: e,
	}
}

// Group creates a new router group with the given path prefix.
// Groups can be used to organize routes and apply middleware to specific route groups.
//
// Parameters:
//   - path: the path prefix for the group
//
// Returns:
//   - *gin.RouterGroup: a new router group instance
//
// Example:
//
//	api := server.Group("/api/v1")
//	api.Get("/users", handleUsers)
//	api.Get("/products", handleProducts)
func (l *Server) Group(path string) *gin.RouterGroup {
	return l.engine.Group(path)
}

// Handle registers a new request handler for the given HTTP method and path.
// The handler function receives a wrapped Request object that contains the gin.Context.
func (l *Server) Handle(method string, path string, handler handlerFunc) {
	l.engine.Handle(method, path, func(ctx *gin.Context) {
		handler(&Request{Context: ctx})
	})
}

// Get registers a new handler for GET HTTP method with the given path.
// It's a shorthand for Handle(http.MethodGet, path, handler).
func (l *Server) Get(path string, handler handlerFunc) {
	l.Handle(http.MethodGet, path, handler)
}

// Post registers a new handler for POST HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPost, path, handler).
func (l *Server) Post(path string, handler handlerFunc) {
	l.Handle(http.MethodPost, path, handler)
}

// Patch registers a new handler for PATCH HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPatch, path, handler).
func (l *Server) Patch(path string, handler handlerFunc) {
	l.Handle(http.MethodPatch, path, handler)
}

// Delete registers a new handler for DELETE HTTP method with the given path.
// It's a shorthand for Handle(http.MethodDelete, path, handler).
func (l *Server) Delete(path string, handler handlerFunc) {
	l.Handle(http.MethodDelete, path, handler)
}

// Put registers a new handler for PUT HTTP method with the given path.
// It's a shorthand for Handle(http.MethodPut, path, handler).
func (l *Server) Put(path string, handler handlerFunc) {
	l.Handle(http.MethodPut, path, handler)
}
