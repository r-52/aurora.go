package models

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model

	Name     string
	Template string `gorm:"type:text"`
	IsActive bool
}
