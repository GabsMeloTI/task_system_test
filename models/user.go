package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`

	Projects []Project `gorm:"foreignkey:UserID" json:"projects"`
	Sections []Section `gorm:"foreignkey:UserID" json:"sections"`
	Tasks    []Task    `gorm:"foreignkey:UserID" json:"tasks"`
	Comments []Comment `gorm:"foreignkey:UserID" json:"comments"`
}
