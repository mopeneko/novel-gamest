package infrastructure

import (
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql" // For using *gorm.DB
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db := connect()

	jwtSecretTable := jwtSecret{}
	db.AutoMigrate(&jwtSecretTable)

	db.First(&jwtSecretTable)
	if len(jwtSecretTable.secret) <= 0 {
		jwtSecretTable = newJWTSecret()
		db.Create(&jwtSecretTable)
	}
}

func connect() *gorm.DB {
	gormDB, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=loc=%s",
			"mopeneko",
			"mopepass",
			"db",
			"novelgamest",
			url.QueryEscape("Asia/Tokyo"),
		),
	)
	if err != nil {
		log.Println("Failed to connect to DB. Waiting for 3 seconds...")
		time.Sleep(time.Second * 3)
		return connect()
	}

	return gormDB
}
