package mysql

import (
	"log"
	"os"
	"strconv"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	user     string
	password string
	host     string
	port     int
	database string
}

var once sync.Once
var dbSingleInstance *gorm.DB

// GetInstance The function `GetInstance` establishes a connection to a SQL Server database using the provided
// configuration and returns a single instance of the database connection.
func (fg DB) GetInstance() *gorm.DB {
	// logPrint := logger.InitLog()

	if dbSingleInstance == nil {
		once.Do(func() {
			port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 16)
			if err != nil {
				// logPrint.Error(err.Error())
				log.Fatal(err)
			}

			fg.user = os.Getenv("DB_USER")
			fg.password = os.Getenv("DB_PASSWORD")
			fg.host = os.Getenv("DB_HOST")
			fg.port = int(port)
			fg.database = os.Getenv("DB_NAME")

			var db *gorm.DB
			// Generate DSN
			//dsn := fg.user + ":" + fg.password + "@" + fg.host + ":" + strconv.Itoa(fg.port) + "?database=" + fg.database
			dsn := fg.user + ":" + fg.password + "@tcp(" + fg.host + ":" + strconv.Itoa(fg.port) + ")/" + fg.database + "?charset=utf8mb4&parseTime=True&loc=Local"
			// Connection to db
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

			if err != nil {
				log.Fatal(err)
			}

			dbSingleInstance = db
		})
	}

	return dbSingleInstance
}
