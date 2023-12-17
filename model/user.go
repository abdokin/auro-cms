package model

import (
	"auro-cms/utils"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `json:"email"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password,omitempty"`
	}

	RegisterRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Name     string `json:"name" form:"name" validate:"required,min=2,max=32"`
		Phone    string `json:"phone" form:"phone" validate:"required,e164"`
		Password string `json:"password" form:"password" validate:"required,min=6,max=16"`
	}
	LoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=6,max=16"`
	}
)

func (u *User) TableName() string {
	return "users"
}

func NewUser(req *RegisterRequest) *User {
	return &User{
		Email:    req.Email,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.EncryptPassword([]byte(u.Password))
	return
}
