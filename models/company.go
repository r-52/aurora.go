package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type CompanyPurchaseSetting struct {
	IsSupplierRatingEnabled bool  `gorm:"default:false"`
	SupplierRating          int16 `gorm:"default:0"`
}

type CompanySalesSetting struct {
	IsMaxOpenOrdersActive            bool `gorm:"default:false"`
	IsMaxOpenOrderOnTotalValueActive bool `gorm:"default:false"`
	MaxOpenOrderTotalValue           sql.NullFloat64
}

type Company struct {
	gorm.Model

	CompanyTypeId int
	CompanyType   CompanyType

	Meta Meta `gorm:"embedded"`

	CompanySalesSetting    CompanySalesSetting    `gorm:"embedded;embeddedPrefix:sales_"`
	CompanyPurchaseSetting CompanyPurchaseSetting `gorm:"embedded;embeddedPrefix:purchase_"`

	DeliverAddresses []CompanyAddress
	InvoiceAddresses []CompanyAddress
}
