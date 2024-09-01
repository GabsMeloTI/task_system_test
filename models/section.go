package models

import (
	"github.com/jinzhu/gorm"
)

type Section struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:varchar(250);not null" json:"description"`

	UserID int  `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`

	ProjectID int     `gorm:"not null" json:"project_id"`
	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"project"`

	Tasks []Task `gorm:"foreignkey:SectionID" json:"tasks"`
}
