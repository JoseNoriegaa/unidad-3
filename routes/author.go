package routes

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	models "github.com/josenoriegaa/unidad-3/models"
)

// FindAuthor finds an author
func FindAuthor(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var author models.Author

	// Query
	err := db.Where("fullname LIKE ?", c.Param("fullname") + "%").First(&author).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
	} else {
		c.JSON(http.StatusOK, author)
	}
}

// FindAllAuthors finds all authors
func FindAllAuthors(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var author []models.Author

	// Query
	db.Order("created_at", true).Find(&author)

	c.JSON(http.StatusOK, author)
}

// CreateAuthor creates a new author into the database
func CreateAuthor(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Query
	db.Create(&author)

	message := "The author has been registered successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": author})
}

// UpdateAuthor updates an author
func UpdateAuthor(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var author models.Author

	// Query
	err := db.Where("fullname LIKE ?", c.Param("fullname") + "%").First(&author).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
		return
	}

	// Model Input
	var replaceAuthor models.Author
	if err := c.ShouldBindJSON(&replaceAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Query
	db.Model(&author).Updates(replaceAuthor)

	message := "The author has been updated successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": author})
}

// DeleteAuthor deletes an author from the database
func DeleteAuthor(c *gin.Context) {
	// Db instance
	db := c.MustGet("db").(*gorm.DB)

	// Model
	var author models.Author

	// Query
	err := db.Where("fullname LIKE ?", c.Param("fullname") + "%").First(&author).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource not found."})
		return
	}

	// Query
	db.Delete(&author)

	message := "The author has been deleted successfully"

	c.JSON(http.StatusOK, gin.H{"message": message, "data": author})
}
