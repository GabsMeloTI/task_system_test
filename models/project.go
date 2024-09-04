package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model  `swaggerignore:"true"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"type:varchar(50)" json:"status"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignkey:UserID" json:"user"`

	Sections []Section `gorm:"foreignkey:ProjectID" json:"sections"`
}
