package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/jsonapi"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User the user model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `jsonapi:"name"`
	Email     string    `jsonapi:"email"`
	Password  string    `jsonapi:"-"`
	CreatedAt time.Time `jsonapi:"createdAt"`
	UpdatedAt time.Time `jsonapi:"updatedAt"`
}

// TableName for gorm
func (User) TableName() string {
	return "users"
}

// GetFirstByID gets the user by his ID
func (u *User) GetFirstByID(id string) error {
	err := DB().Where("id=?", id).First(u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

// GetFirstByEmail gets the user by his email
func (u *User) GetFirstByEmail(email string) error {
	err := DB().Where("email=?", email).First(u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

// Create a new user
func (u *User) Create() error {
	db := DB().Create(u)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}

// Signup a new user
func (u *User) Signup() error {
	var user User
	err := user.GetFirstByEmail(u.Email)

	if err == nil {
		return ErrUserExists
	} else if err != ErrDataNotFound {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// replace the plaintext password with ciphertext password
	u.Password = string(hash)

	return u.Create()
}

func (u *User) SignupJson() error {
	var user User
	err := user.GetFirstByEmail(u.Email)

	if err == nil {
		return ErrUserExists
	} else if err != ErrDataNotFound {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// replace the plaintext password with ciphertext password
	u.Password = string(hash)

	return u.Create()
}

// Login a user
func (u *User) Login(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// LoginByEmailAndPassword login a user by his email and password
func LoginByEmailAndPassword(email, password string) (*User, error) {
	var user User
	err := user.GetFirstByEmail(email)
	if err != nil {
		return &user, err
	}

	return &user, user.Login(password)
}

func (user User) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://localhost:8080/api/user/%d", user.ID),
	}
}


func (user User) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	return nil
}

// JSONAPIMeta implements the Metable interface for a blog
func (user User) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"detail": "extra details regarding the post",
	}
}

// JSONAPIRelationshipMeta implements the RelationshipMetable interface for a blog
func (user User) JSONAPIRelationshipMeta(relation string) *jsonapi.Meta {
	return nil
}