package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string
	Body   string
	IsDone bool
	// add your field here...
}

// Override the table name
func (Todo) TableName() string {
	return "todos"
}
