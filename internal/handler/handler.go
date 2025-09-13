package handler

import "github.com/precioussilence/go-gin-starter/internal/service"

type Handlers struct {
	User *UserHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User: NewUserHandler(services.User),
	}
}
