package model

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `json:"email"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}

	UserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=6,max=16"`
	}
)
