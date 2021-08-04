package repository

import (
	"context"
	"mock_gin/pkg/repository/model"
)

type User interface {
	NewUser(ctx context.Context, name string) model.User
	UpdateUser(ctx context.Context, id int64, name string) (model.User, error)
}
