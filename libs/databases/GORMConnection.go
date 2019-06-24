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
)

// Struct type gormConn - type to deal with connection resources
type SqlConn struct {
	DB *gorm.DB
}

// Connection created as a singleton class along with GORM to add flexeibilty in
// terms of SQL flavor preferences
func SQLConn() (*SqlConn, error) {
	conn := SqlConn{}
	var e error
	once.Do(func() {
		// initialize godotenv
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
		conn.DB, e = gorm.Open(os.Getenv("db_database"), uri)
		if e != nil {
			log.It.WriteLog("error", e.Error(), log.It.GetTraceMsg())
		}
	})
	return &conn, e
}

//  Migrate method - create tables based on models
func (sc *SqlConn) Migrate(models ...interface{}) {
	sc.DB.Debug().AutoMigrate(models...)
}
