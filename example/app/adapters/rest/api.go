package rest

import (
	"fmt"
	"net/http"

	"example/app/entity"
	"example/app/ports"
	"github.com/gin-gonic/gin"
	"github.com/go-liquor/liquor/v2/app/adapters/rest"
)

type UsersApi struct {
	svc ports.UserService
}

func NewUsersApi(svc ports.UserService) rest.Api {
	return &UsersApi{
		svc: svc,
	}
}

func (u *UsersApi) Routes(s *rest.Route) {
	s.Get("/", u.Get)
	s.Post("/", u.Create)
	gp := s.Group("/").Middleware(func(context *gin.Context) {
		fmt.Println("chamouy aquio", context.FullPath())
		context.Next()
	})
	{
		gp.Get("/tatu", func(request *rest.Request) {
			request.Status(http.StatusOK)
		})
	}
}

func (u *UsersApi) Get(r *rest.Request) {
	users := u.svc.Get(r.Request.Context())
	r.JSON(http.StatusOK, users)
}

func (u *UsersApi) Create(r *rest.Request) {
	var user *entity.User
	if err := r.ShouldBindJSON(&user); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := u.svc.Create(r.Request.Context(), user); err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.Status(http.StatusCreated)
}
