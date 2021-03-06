package database

import (
	"context"
	"mock_gin/pkg/repository/model"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) NewUser(ctx context.Context, name string) model.User {
	// r.db.Create("..")
	return model.User{
		ID:   1,
		Name: name,
	}
}
func (r *UserRepo) UpdateUser(ctx context.Context, id int64, name string) (model.User, error) {
	// r.db.Update("...")
	return model.User{}, nil
}
