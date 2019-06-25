// Author: James Mallon <jamesmallondev@gmail.com>
// models package - package contains all the models of the system
package models

import (
	"github.com/jinzhu/gorm"
)

// Struct type Comment - Comment model
type Comment struct {
	gorm.Model
	Title   string
	Content string
	Author  string
	Date    string
}
