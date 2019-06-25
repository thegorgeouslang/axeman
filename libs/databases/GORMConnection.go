// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

import (
	log "axeman/libs/logger"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
	e    error
)

// init function - data and process initialization
// Connection created as a singleton class along with GORM to add flexeibilty in
// terms of SQL flavor preferences
func init() {
	e = godotenv.Load()
	if e != nil {
		log.It.WriteLog("error", e.Error(), log.It.GetTraceMsg())
	}
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_user"),
		os.Getenv("db_name"),
		os.Getenv("db_pass"),
	)
	// fmt.Println(uri)
	conn, e := gorm.Open(os.Getenv("db_database"), uri)
	if e != nil {
		log.It.WriteLog("error", e.Error(), log.It.GetTraceMsg())
	}
	db = conn
}

// SQLConn function - retrieves the connection object
func SQLConn() (*gorm.DB, error) {
	return db, e
}

//  Migrate method - create tables based on models
func Migrate(models ...interface{}) {
	db.Debug().AutoMigrate(models...)
}
