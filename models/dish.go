package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Thumbnail   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
