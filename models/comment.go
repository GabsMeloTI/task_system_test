package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	gorm.Model  `swaggerignore:"true"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	PublishedAt time.Time `gorm:"type:timestamptz;not null" json:"published_at"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`

	UserID int  `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignkey:UserID" json:"user"`

	TaskID int  `gorm:"not null" json:"task_id"`
	Task   Task `gorm:"foreignkey:TaskID" json:"task"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.PublishedAt = time.Now() // Define a data e hora atual para o campo PublishedAt
	return
}
