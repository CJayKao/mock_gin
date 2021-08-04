package main

import (
	"context"
	"fmt"
	mocks "mock_gin/mocks/repository"
	"mock_gin/pkg/repository/database"
	"mock_gin/pkg/repository/model"
	"mock_gin/pkg/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	db, _ := database.NewDb("")
	userService := service.NewUserService(db)

	r.POST("/user", func(c *gin.Context) {
		type UserData struct {
			Name string `json:"name"`
		}
		var userData UserData
		c.BindJSON(&userData)
		user := userService.NewUser(c, userData.Name)
		fmt.Println("userData", userData)
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s", user))
	})

	r.POST("/mockUser", func(c *gin.Context) {
		mockUser := mocks.User{}
		mockUser.On("NewUser",
			mock.MatchedBy(func(context.Context) bool { return true }),
			mock.AnythingOfType("string")).
			Return(
				func(ctx context.Context, name string) model.User {
					t := time.Now()
					return model.User{
						Name:      name,
						CreatedAt: t,
						UpdatedAt: t,
					}
				},
			)

		data := mockUser.NewUser(context.TODO(), "nameXXX")
		c.String(http.StatusOK, fmt.Sprintf(" %v", data))
	})
	return r
}
