package service

import "github.com/precioussilence/go-gin-starter/internal/repository"

type Services struct {
	User *UserService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repositories.User),
	}
}
