package main

// Author: Jose Noriega
// Email: josenoriega723@gmail.com
// Last Update: 2020-04-23

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	routes "github.com/josenoriegaa/unidad-3/routes"
)

func main() {
	r := gin.Default()


	DB_USER := os.Getenv("DB_USER")
	DB_PWD := os.Getenv("DB_PWD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("PORT")

	connectionString := DB_USER + ":" + DB_PWD + "@(" + DB_HOST + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Middleware to provide access to all the endpoints to the db object
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Endpoints
	r.GET("/api/v1/authors/:fullname", routes.FindAuthor)
	r.GET("/api/v1/authors/", routes.FindAllAuthors)
	r.POST("/api/v1/authors/", routes.CreateAuthor)
	r.PUT("/api/v1/authors/:fullname", routes.UpdateAuthor)
	r.DELETE("/api/v1/authors/:fullname", routes.DeleteAuthor)
	
	r.GET("/api/v1/books/:id", routes.FindBook)
	r.GET("/api/v1/books/", routes.FindAllBooks)
	r.POST("/api/v1/books/", routes.CreateBook)
	r.PUT("/api/v1/books/:id", routes.UpdateBook)
	r.DELETE("/api/v1/books/:id", routes.DeleteBook)

	SERVER_PORT := ":" + PORT
	fmt.Println("Server running on port " + SERVER_PORT)
	r.Run(SERVER_PORT)
}
