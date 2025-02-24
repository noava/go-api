package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize the database with gorm
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/app.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

    if err := DB.AutoMigrate(&Plant{}); err != nil {
        panic("Failed to migrate database: " + err.Error())
    }
}


type Plant struct {
	Name      string `gorm:"uniqueIndex"`
	StartDay  uint
	EndDay    uint
	Type 			string
}