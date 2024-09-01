package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title              string    `gorm:"type:varchar(100);not null" json:"title"`
	Description        string    `gorm:"type:text" json:"description"`
	ExpectedCompletion time.Time `gorm:"type:timestamp" json:"expected_completion"`
	Priority           Priority  `gorm:"type:varchar(50)" json:"priority"`
	Status             string    `gorm:"type:varchar(50)" json:"status"`

	UserID int  `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`

	SectionID int     `gorm:"not null" json:"section_id"`
	Section   Section `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"section"`

	Comments []Comment `gorm:"foreignkey:TaskID" json:"comments"`
	Subtasks []Subtask `gorm:"foreignkey:TaskID" json:"subtasks"`

	Labels []Label `gorm:"many2many:task_labels;" json:"labels"`
}
