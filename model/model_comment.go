package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      uint  `jsonapi:"primary,comments"`
	PostID 	uint  `jsonapi:"attr,post_id"`	
	Body	string    `jsonapi:"attr,body"`
	CreatedAt time.Time `jsonapi:"attr,createdAt"`
	UpdatedAt time.Time `jsonapi:"attr,updatedAt"`
	DeletedAt time.Time `jsonapi:"attr,deletedAt"`
}

func (Comment) TableName() string {
	return "comment"
}