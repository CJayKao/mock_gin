package main

import (
	"fmt"
	"mock_gin/pkg/repository/database"
	"mock_gin/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	return r
}
