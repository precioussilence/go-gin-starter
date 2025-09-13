package service

import (
	"context"
	"errors"

	"github.com/precioussilence/go-gin-starter/internal/domain"
	"github.com/precioussilence/go-gin-starter/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userReposity repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userReposity,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
