package models

import (
	"gorm.io/gorm"
)

type CompanyType struct {
	gorm.Model

	Companies        []Company
	Name             string `gorm:"type:text"`
	IsUsedInSales    bool
	IsUsedInPurchase bool
}
