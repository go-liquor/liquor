package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-liquor/liquor/v2/app/adapters/rest"
	"github.com/go-liquor/liquor/v2/example/app/entity"
	"github.com/go-liquor/liquor/v2/example/app/ports"
)

type UsersApi struct {
	svc ports.UserService
}

func NewUsersApi(svc ports.UserService) rest.Api {
	return &UsersApi{
		svc: svc,
	}
}

func (u *UsersApi) Routes(s *rest.Server) {
	s.Get("/", u.Get)
	s.Post("/", u.Create)
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
