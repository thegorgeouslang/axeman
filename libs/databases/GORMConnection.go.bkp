// Author: James Mallon <jamesmallondev@gmail.com>
// databases package -
package databases

import (
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
)

func GORMConn() *gorm.DB {
	once.Do(func() {
		// initialize godotenv
		e := godotenv.Load()
		if e != nil {
			fmt.Print(e)
		}
		uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("db_host"),
			os.Getenv("db_port"),
			os.Getenv("db_user"),
			os.Getenv("db_name"),
			os.Getenv("db_pass"),
		)
		fmt.Println(uri)
		conn, err := gorm.Open(os.Getenv("db_database"), uri)
		if err != nil {
			fmt.Print(err)
		}
		db = conn
	})
	return db
}

//  -
func Migrate(models ...interface{}) {
	db.Debug().AutoMigrate(models...)
}
