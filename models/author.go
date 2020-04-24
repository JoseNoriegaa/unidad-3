package models

import "time"

// Author model
type Author struct {
	FullName string `json:"fullname" binding:"required" gorm:"primary_key;column:nombre_completo;type:varchar(200);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

// TableName sets the table name
func (Author) TableName() string {
  return "autores"
}
