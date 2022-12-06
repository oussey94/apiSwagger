package controllers

import (
	"apiSwagger/initializers"
	"apiSwagger/models"

	"github.com/gin-gonic/gin"
)

// PostPost             godoc
// @Summary      Store a new post
// @Description  Takes a post JSON and store in DB. Return saved JSON.
// @Tags         posts
// @Produce      json
// @Param        post  body      models.Post  true  "Add Post JSON"
// @Success      200   {object}  models.Post
// @Router       /posts [post]
func PostsController(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// GetAllPost             godoc
// @Summary      Get posts array
// @Description  Responds with the list of all posts as JSON.
// @Tags         posts
// @Produce      json
// @Success      200  {array}  models.Post
// @Router       /posts [get]
func GetAllPost(c *gin.Context) {
	var posts []models.Post
	//get all post
	initializers.DB.Find(&posts)
	//response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// GetPostById godoc
//
//	@Summary		Show a post
//	@Description	get string by ID
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int	true	"post ID"
//	@Success		200	{object}	models.Post
//	@Router			/posts/{id} [get]
func GetPostById(c *gin.Context) {
	//get id post
	id := c.Param("id")
	//get post
	var post models.Post
	initializers.DB.First(&post, id)
	//response
	c.JSON(200, gin.H{
		"post": post,
	})

}

// UpdatePost godoc
//
//	@Summary		Update a post
//	@Description	Update by json post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"post ID"
//	@Param			post	body		models.Post	true	"Update post"
//	@Success		200		{object}	models.Post
//	@Router			/posts/{id} [put]
func UpdatePost(c *gin.Context) {
	//get id off the url
	id := c.Param("id")
	//get data of the request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//fin the post wer updating
	var post models.Post
	initializers.DB.First(&post, id)
	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})
	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// DeleteAccount godoc
//
//	@Summary		Delete an account
//	@Description	Delete by account ID
//	@Tags			posts
//	@Accept			json
//  @Param			id	path		int	true	"post ID"
//	@Success		204	{object}	models.Post
//	@Router			/posts/{id} [delete]
func DeletePost(c *gin.Context) {
	//get id of the url
	id := c.Param("id")
	//delete the post
	initializers.DB.Delete(&models.Post{}, id)
	//respond
	c.Status(200)
}
