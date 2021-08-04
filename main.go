package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// r := setupRouter()
	// r.Run(":8080")
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
	})
	return r
}
