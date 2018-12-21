package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")

	dbUri := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, name)
	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatalf("cannot open mysql connection: %s", err)
	}
	defer conn.Close()

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
