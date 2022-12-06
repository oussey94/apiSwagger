package main

import (
	"apiSwagger/controllers"
	_ "apiSwagger/docs"
	"apiSwagger/initializers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @title     Gin_Go Poststore API
// @version         1.0
// @description     A post management service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   mbodji ousseynonu
// @contact.url    https://twitter.com/weuz
// @contact.email  mbodjiousseynou94@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
func main() {
	r := gin.Default()
	r.POST("/api/v1/posts", controllers.PostsController)
	r.PUT("/api/v1/posts/:id", controllers.UpdatePost)
	r.GET("/api/v1/posts", controllers.GetAllPost)
	r.GET("/api/v1/posts/:id", controllers.GetPostById)
	r.DELETE("/api/v1/posts/:id", controllers.DeletePost)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
