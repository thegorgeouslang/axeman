// Author: James Mallon <jamesmallondev@gmail.com>
// dao package - package contains the DAOs of the system
package dao

import (
	"axeman/libs/databases"
	log "axeman/libs/logger"
	"axeman/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// userDAO -
type userDAO struct{}

//
var conn *gorm.DB

// init function - data and process initialization
func init() {
	var e error
	conn, e = databases.SQLConn()
	if e != nil {
		log.It.WriteLog("error", e.Error(), log.It.GetTraceMsg())
	}
}

//
func UserDAO() *userDAO {
	return &userDAO{}
}

// InsertUser method - Stores a new user in the system
func (ud *userDAO) InsertUser(user *models.User) (err error) {
	var wg sync.WaitGroup
	errChannel := make(chan error) // creates a new channel
	wg.Add(1)
	go func() {
		// it creates a hashed byte slice from the user password
		hb, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err == nil {
			user.Password = string(hb)

			conn.Create(&user)
			if conn.NewRecord(user) {
				err = fmt.Errorf("The user %s wasn't created", user.Email)
			}
		}
		errChannel <- err
		defer wg.Done()
	}()
	err = <-errChannel
	defer close(errChannel)
	wg.Wait()
	return
}
