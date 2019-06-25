// Author: James Mallon <jamesmallondev@gmail.com>
// dao package - package contains the DAOs of the system
package dao

import (
	"axeman/libs/databases"
	log "axeman/libs/logger"
	"axeman/models"
	"github.com/jinzhu/gorm"
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

//
func (ud *userDAO) InsertUser(user *models.User) (res bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		conn.Create(&user)
		res = conn.NewRecord(user)
		defer wg.Done()
	}()
	wg.Wait()
	return
}
