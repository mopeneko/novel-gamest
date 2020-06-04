package infrastructure

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/mopeneko/novel-gamest/api/domain"

	_ "github.com/go-sql-driver/mysql" // For using *gorm.DB
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB initials db variable and migrate
func InitDB() {
	db = connect()

	jwtSecretTable := jwtSecret{}
	db.AutoMigrate(&jwtSecretTable, &domain.User{}, &domain.Game{}, &domain.Post{})

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
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
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
