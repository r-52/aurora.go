package database


import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )


func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
