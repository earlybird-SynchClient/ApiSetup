package supplier

import (
	"github.com/jinzhu/gorm"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
)

type Supply struct {
	gorm.Model

	AuthCode string
	URL string

	VendorCodeID uint
	VendorCode *basic.VendorCode		`json:"vendorcode"`
	SupplierID uint
	Supplier *basic.Supplier			`json:"supplier"`
	SupplierEmailID uint
	SupplierEmail *basic.SupplierEmail	`json:"email"`
	MarketID uint
	Market *market.Market				`json:"invoicestatus"`
}

func UpdateEntry(i *Supply) bool{
	var _elm1 basic.VendorCode
	database.Conn.Where("id = ?", i.VendorCodeID).Find(&_elm1)
	i.VendorCode = &_elm1

	var _elm2 basic.Supplier
	database.Conn.Where("id = ?", i.SupplierID).Find(&_elm2)
	i.Supplier = &_elm2

	var _elm3 basic.SupplierEmail
	database.Conn.Where("id = ?", i.SupplierEmailID).Find(&_elm3)
	i.SupplierEmail = &_elm3

	var _elm4 market.Market
	database.Conn.Where("id = ?", i.MarketID).Find(&_elm4)
	i.Market = &_elm4

	return true;
}

func GetSupplyByID(id uint) (*Supply, error) {
	var _invoce Supply
	err := database.Conn.
		Where("id = ?", id).
		Preload("VendorCode").
		Preload("Supplier").
		Preload("SupplierEmail").
		Preload("Market").
		Find(&_invoce).
		Error

	return &_invoce, err
}

func GetSupplyByToken(token string) (*Supply, error) {
	var _supply Supply
	err := database.Conn.
		Where("auth_code = ?", token).
		Preload("VendorCode").
		Preload("Supplier").
		Preload("SupplierEmail").
		Preload("Market").
		Find(&_supply).
		Error

	return &_supply, err
}

func (i *Supply) BeforeCreate(scope *gorm.Scope) error {
	UpdateEntry(i)

	return nil
}

func (i *Supply) BeforeUpdate(scope *gorm.Scope) error {
	UpdateEntry(i)

	return nil
}