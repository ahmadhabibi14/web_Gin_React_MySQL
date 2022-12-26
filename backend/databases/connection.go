package databases

import (
	"backend/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "root@tcp(127.0.0.1:3306)/movie_web"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error DB :::::::: !!!")
	}

	if err = db.AutoMigrate(&models.Film{}); err != nil {
		fmt.Println("AutoMigrate ERROR :::::::: !!!")
	}
	return db, err
}
