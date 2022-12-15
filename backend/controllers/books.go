package controllers

import (
	"backend/database"
	"backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

// ini untuk menyalin yang ada di db ke struct
func New() *BookRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Book{})
	return &BookRepo{Db: db}
}

// Create book
func (repository *BookRepo) CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.CreateBook(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Get all Book
func (repository *BookRepo) GetBooks(c *gin.Context) {
	var book []models.Book
	err := models.GetBooks(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Get book by ID
func (repository *BookRepo) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Update book
func (repository *BookRepo) UpdateBook(c *gin.Context) {
	var book models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&book)
	err = models.UpdateBook(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// delete book
func (repository *BookRepo) DeleteBook(c *gin.Context) {
	var book models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteBook(repository.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
