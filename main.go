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
	r.GET("/user", controllers.UserAll)
	r.GET("/user/:id", controllers.UserShow)

	r.Run() // listen and serve on 0.0.0.0:8080
}
