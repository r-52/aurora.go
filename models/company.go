package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	CompanyTypeId int
	CompanyType   CompanyType

	DeliverAddresses []CompanyAddress
	InvoiceAddresses []CompanyAddress
}
