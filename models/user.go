package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"` // create primary key
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitializeDB(db *gorm.DB) {

	// Auto migrate the schema (creates the table if it doesn't exist)
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
