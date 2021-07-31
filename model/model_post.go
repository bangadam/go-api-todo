package model

import (
	"errors"
	"fmt"
	"strconv"
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
}

func (Post) TableName() string {
	return "post"
}

func (p *Post) GetPostsPaginate(page string, perPage string) (post []*Post, err error) {

	newpPage, err := strconv.Atoi(page)
	newpPerPage, err := strconv.Atoi(perPage)	

	db := DB().Preload("Comments").Limit(newpPerPage).Offset(newpPerPage * (newpPage - 1)).Find(&post)

	if db.Error != nil {
		return nil, db.Error
	}

	return post, nil
}

func (p *Post) StorePost() error {
	db := DB().Create(p)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}
	return nil
}

func (p *Post) UpdatePost(title string, description string, id string) error {
	db := DB().Where("id = ?", id).First(&p)

	if db.Error != nil {
		return db.Error
	}

	p.Title = title
	p.Body = description
	db.Save(p)

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (p *Post) GetPostById(id string) error {
	err := DB().Preload("Comments").First(&p, id).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

func (p *Post) DeletePost(id string) error {
	db := DB().Where("id = ?", id).Delete(&p)

	if db.Error != nil {
		return db.Error
	}

	return nil
	
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