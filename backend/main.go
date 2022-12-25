package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Film struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Year  int    `json:"release_year"`
	Genre string `json:"genre"`
}

func Connect() (*gorm.DB, error) {
	dsn := "root@tcp(127.0.0.1:3306)/movie_web"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error DB :::::::: !!!")
	}

	if err = db.AutoMigrate(&Film{}); err != nil {
		fmt.Println("AutoMigrate ERROR :::::::: !!!")
	}
	return db, err
}

func getAllFilms(c *gin.Context) {
	var Films []Film
	db, err := Connect()
	if err != nil {
		fmt.Println("ERROR GET ::::::::::::: !!!")
	}

	if err := db.Find(&Films).Error; err != nil {
		fmt.Println("ERROR QUERY :::::::::: !!!")
	}
	c.JSON(http.StatusOK, Films)
}

func main() {
	route := gin.Default()
	route.GET("/films", getAllFilms)
	route.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
