package main

import (
	"github.com/luizweitz/hello-gin/initializers"
	"github.com/luizweitz/hello-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
