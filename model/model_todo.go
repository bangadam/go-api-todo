package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID		uint `gorm:"primary_key" json:"id"`
	Title	string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Tablename gorm
func (t *Todo) TableName() string {
	return "todos"
}

// Create
func (t *Todo) Create() error {
	db := DB().Create(t)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}

func (t *Todo) GetTodoById(id string) error {
	err := DB().Where("id = ?", id).First(&t).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

func (t *Todo) Update(title string, description string, id string) error {
	db := DB().Where("id = ?", id).First(&t)

	if db.Error != nil {
		return db.Error
	}
	
	t.Title = title
	t.Description = description
	db.Save(&t)

	if db.Error != nil {
		return  db.Error
	}

	return nil
}

func (t *Todo) DeleteTodo(id string) error {
	db := DB().Where("id = ?", id).Delete(&t)

	if db.Error != nil {
		return db.Error
	}

	return nil
}