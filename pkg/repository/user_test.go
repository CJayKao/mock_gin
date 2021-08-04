package repository

import (
	"context"
	"fmt"
	mocks "mock_gin/mocks/repository"
	"mock_gin/pkg/repository/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUser(t *testing.T) {
	mockUser := mocks.User{}
	mockUser.On("NewUser",
		mock.MatchedBy(func(context.Context) bool { return true }),
		mock.AnythingOfType("string")).
		Return(
			func(ctx context.Context, name string) model.User {
				return model.User{Name: name}
			},
		)
	data := mockUser.NewUser(context.TODO(), "name")
	fmt.Println("data=>", data.Name)
}
