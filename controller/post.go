package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/hyperjiang/gin-skeleton/model"
)

type PostController struct {}

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