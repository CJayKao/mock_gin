package service

import (
	"context"
	"mock_gin/pkg/repository"
	"mock_gin/pkg/repository/database"
	"mock_gin/pkg/repository/model"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *database.UserRepo
}

func NewUserService(db *gorm.DB) repository.User {
	return &UserService{
		userRepo: database.NewUserRepo(db),
	}
}

func (s *UserService) NewUser(ctx context.Context, name string) model.User {
	return s.userRepo.NewUser(ctx, name)
}
func (s *UserService) UpdateUser(ctx context.Context, id int64, name string) (model.User, error) {
	return s.userRepo.UpdateUser(ctx, id, name)
}
