package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/hyperjiang/gin-skeleton/model"
)

type PostController struct {}

type UpdatePost struct {
	Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}

func (ctrl PostController) GetPostsPaginate(c *gin.Context) {
	post := new(model.Post)
	page := c.Query("page")
	perPage := c.Query("per_page")

	if page == "" {
		page = "1"
	}

	if perPage == "" {
		perPage = "10"
	}

	posts, err := post.GetPostsPaginate(page, perPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(c.Writer, posts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *PostController) StorePost(c *gin.Context) {
	post := new(model.Post)
	
	if err := jsonapi.UnmarshalPayload(c.Request.Body, post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := post.StorePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := jsonapi.MarshalPayload(c.Writer, post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *PostController) UpdatePost(c *gin.Context) {
	post := new(model.Post)
	// postValidation := new(UpdatePost)

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	if err := jsonapi.UnmarshalPayload(c.Request.Body, post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := post.Title
	body := post.Body

	if err := post.UpdatePost(title, body, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := jsonapi.MarshalPayload(c.Writer, post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *PostController) GetTodoById(c *gin.Context) {
	post := new(model.Post)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := post.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(c.Writer, post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ctrl *PostController) DeletePost(c *gin.Context) {
	post := new(model.Post)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := post.DeletePost(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}