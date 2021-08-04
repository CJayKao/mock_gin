package repository

// import (
// 	"context"
// 	"mock_gin/pkg/repository/model"

// 	"gorm.io/gorm"
// )

// type UserRepository struct {
// 	db *gorm.DB
// }

// // NewRepository repository new constructor
// func NewRepository(db *gorm.DB) *UserRepository {
// 	return &UserRepository{
// 		db: db,
// 	}
// }

// func (r *UserRepository) NewUser(ctx context.Context, name string) model.User {
// 	return model.User{}
// }
// func (r *UserRepository) UpdateUser(ctx context.Context, id int64, name string) (model.User, error) {
// 	return model.User{}, nil
// }
