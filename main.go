package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizweitz/hello-gin/controllers"
	"github.com/luizweitz/hello-gin/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/user", controllers.UserCreate)
	r.GET("/user", controllers.UserListAll)
	r.GET("/user/:id", controllers.UserShow)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)

	r.Run()
}
