package main

import (
	"apiSwagger/controllers"
	"apiSwagger/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsController)
	r.GET("/posts", controllers.GetAllPost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
