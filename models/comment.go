package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Content     string    `gorm:"type:text;not null" json:"content"`
	PublishedAt time.Time `gorm:"type:autoCreateTime;not null" json:"published_at"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`

	UserID int  `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignkey:UserID" json:"user"`

	TaskID int  `gorm:"not null" json:"task_id"`
	Task   Task `gorm:"foreignkey:TaskID" json:"task"`
}
