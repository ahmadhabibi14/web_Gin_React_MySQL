package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"

// // const DB_PASSWORD = ""
// const DB_NAME = "crud_db"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	db, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/crud_db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		return nil
	}
	return db
}
