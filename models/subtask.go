package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Subtask struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Status      string    `gorm:"type:varchar(50)" json:"status"`

	TaskID int  `gorm:"not null" json:"task_id"`
	Task   Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"task"`
}
