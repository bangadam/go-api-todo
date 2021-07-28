package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/jsonapi"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID	uint	`jsonapi:"primary,posts"`
	Title	string	`jsonapi:"attr,title"`
	Body	string	`jsonapi:"attr,body"`
	Comments []*Comment `jsonapi:"relation,comments,omitempty"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	DeletedAt time.Time `jsonapi:"attr,deletedAt"`
}

func (Post) TableName() string {
	return "post"
}

func (p *Post) GetPostById(id string) error {
	err := DB().Preload("Comments").First(&p, id).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}


func (post Post) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://localhost:8080/api/post/%d", post.ID),
	}
}


func (post Post) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	if relation == "comments" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://localhost:8080/api/post/%d/comments", post.ID),
		}
	}
	return nil
}

// JSONAPIMeta implements the Metable interface for a blog
func (post Post) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"includes": []string{"comments"},
		"detail": "extra details regarding the post",
	}
}

// JSONAPIRelationshipMeta implements the RelationshipMetable interface for a blog
func (post Post) JSONAPIRelationshipMeta(relation string) *jsonapi.Meta {
	if relation == "comments" {
		return &jsonapi.Meta{
			"detail": "comments meta information",
		}
	}
	return nil
}