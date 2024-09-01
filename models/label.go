package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Color       string `gorm:"type:varchar(50)" json:"color"`
	Description string `gorm:"type:text" json:"description"`

	Tasks []Task `gorm:"many2many:task_labels;association_jointable_foreignkey:task_id;join_table_foreignkey:label_id" json:"tasks"`
}
