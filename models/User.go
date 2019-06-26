// Author: James Mallon <jamesmallondev@gmail.com>
// models package - package contains all the models of the system
package models

import (
	"github.com/jinzhu/gorm"
)

// Struct type User - User model
type User struct {
	gorm.Model
	Email        string
	Password     string
	SessionToken string
}
