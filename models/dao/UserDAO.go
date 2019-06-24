// Author: James Mallon <jamesmallondev@gmail.com>
// dao package - package contains the DAOs of the system
package dao

import (
	"axeman/libs/databases"
	"axeman/models"
	"sync"
)

// userDAO -
type userDAO struct {
	//conn *databases.SqlConn
}

//
var conn *databases.SqlConn

// init function - data and process initialization
func init() {

	conn, _ = databases.SQLConn()
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
		conn.DB.Create(&user)
		res = conn.DB.NewRecord(user)
		defer wg.Done()
	}()
	wg.Wait()
	return
}
