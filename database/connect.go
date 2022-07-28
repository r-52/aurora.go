package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"path"

	"github.com/r-52/aurora/models"
	"github.com/r-52/aurora/util"
)

func ConnectDB() {
	p := GetDatabasePath()
	var err error
	DB, err = gorm.Open(sqlite.Open(p), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Customer{})
}

func GetDatabasePath() string {
	p := util.GetCurrentExecPath()
	res := path.Join(p, "db", "aurora.db")
	return res
}
