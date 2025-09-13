package repository

import (
	"context"

	"github.com/precioussilence/go-gin-starter/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, id uint, field string, value any) (int, error)
	UpdateMultiple(ctx context.Context, id uint, values map[string]any) error
	Delete(ctx context.Context, id uint) (int, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	return gorm.G[domain.User](r.db).Create(ctx, user)
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := gorm.G[domain.User](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := gorm.G[domain.User](r.db).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := gorm.G[domain.User](r.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*domain.User, len(users))
	for i, v := range users {
		result[i] = &v
	}
	return result, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, id uint, field string, value any) (int, error) {
	return gorm.G[domain.User](r.db).Where("id = ?", id).Update(ctx, field, value)
}

func (r *userRepositoryImpl) UpdateMultiple(ctx context.Context, id uint, values map[string]any) error {
	result := r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", id).
		Updates(values)
	return result.Error
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id uint) (int, error) {
	return gorm.G[domain.User](r.db).Where("id = ?", id).Delete(ctx)
}
