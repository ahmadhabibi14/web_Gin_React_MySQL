package main

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/films", controllers.GetAllFilms)
	route.POST("/films/add", controllers.CreateFilm)
	route.Run("localhost:8080")
}
