package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	Email      string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password   string `gorm:"type:varchar(100);not null" json:"password"`
	Avatar     string `gorm:"type:varchar(255)" json:"avatar"`

	Projects []Project `gorm:"foreignkey:UserID" json:"projects" swaggerignore:"true"`
	Sections []Section `gorm:"foreignkey:UserID" json:"sections" swaggerignore:"true"`
	Tasks    []Task    `gorm:"foreignkey:UserID" json:"tasks" swaggerignore:"true"`
	Comments []Comment `gorm:"foreignkey:UserID" json:"comments" swaggerignore:"true"`
}
