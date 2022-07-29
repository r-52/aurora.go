package models

import (
	"gorm.io/gorm"
)

type CompanyAddress struct {
	gorm.Model

	Company   Company
	CompanyID int

	Street1    string
	Street2    string
	Zip        string
	City       string
	Name1      string
	Name2      string
	County     string
	Department string

	Country   Country
	CountryID int
}
