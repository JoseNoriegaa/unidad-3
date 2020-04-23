package routes

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	models "github.com/josenoriegaa/unidad-3/models"
)

// FindBook finds a book
func FindBook(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var book models.Book

	// Query
	err := db.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

// FindAllBooks finds all the books
func FindAllBooks(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var books []models.Book

	// Query
	db.Order("created_at", true).Find(&books)

	c.JSON(http.StatusOK, books)
}

// CreateBook creates a new book into the database
func CreateBook(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the author exists
	var author models.Author
	// Query
	err := db.Where("fullname LIKE ?", book.Author + "%").First(&author).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The author is not registered."})
		return
	}

	// Query
	db.Create(&book)

	message := "The book has been registered successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": book})
}

// UpdateBook updates a book
func UpdateBook(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var book models.Book

	// Query
	err := db.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
		return
	}

	// Model Input
	var replaceBook models.Book
	if err := c.ShouldBindJSON(&replaceBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Query
	db.Model(&book).Updates(replaceBook)

	message := "The book has been updated successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": book})
}

// DeleteBook deletes a book from the database
func DeleteBook(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var book models.Book

	// Query
	err := db.Where("id = ?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
		return
	}

	// Query
	db.Delete(&book)

	message := "The book has been deleted successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": book})
}
