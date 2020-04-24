package models

import "time"

// BookModel model
type Book struct {
	ID uint `json:"id" gorm:"primary_key;column:id;not null" gorm:`
	Title string `json:"title" binding:"required" gorm:"column:titulo;type:varchar(100);not null"`
	Description string `json:"description" binding:"required" gorm:"column:descripcion;type:varchar(450);not null"`
	Author string `json:"author" binding:"required" gorm:"column:autor;type:varchar(200);not null"`
	Editorial string `json:"editorial" binding:"required" gorm:"column:editorial;type:varchar(200);not null"`
	PublicationDate time.Time `json:"publicationDate" binding:"required" gorm:"column:fecha_publicacion;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

// TableName sets the table name
func (Book) TableName() string {
  return "libros"
}
