package controllers

import (
	"backend/databases"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFilms(c *gin.Context) {
	var Films []models.Film
	db, err := databases.Connect()
	if err != nil {
		fmt.Println("ERROR GET ::::::::::::: !!!")
	}

	if err := db.Find(&Films).Error; err != nil {
		fmt.Println("ERROR QUERY :::::::::: !!!")
	}
	c.JSON(http.StatusOK, Films)
}

func CreateFilm(c *gin.Context) {
	var Films *[]models.Film
	c.BindJSON(&Films)
	db, err := databases.Connect()
	if err != nil {
		fmt.Println("ERROR WEEYYY :::::::::::: !!!")
	}

	if err := db.Create(&Films).Error; err != nil {
		fmt.Println("ERROR CREATE ::::::::::: !!!")
	}
	c.JSON(http.StatusOK, Films)
}
