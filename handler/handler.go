package handler

import (
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)