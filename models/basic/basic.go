package basic

import (
	"github.com/jinzhu/gorm"
	"time"
)

type VendorCode struct {
	gorm.Model
	Name	string		`json:"vendorcode"`
}

type NumberInvoice struct {
	gorm.Model
	Name	string		`json:"invoiceno"`
}

type AmountInvoice struct {
	gorm.Model
	Name	string		`json:"invoiceamount"`
}

type StatusInvoice struct {
	gorm.Model
	Name	string		`json:"invoicestatus"`
}

type Supplier struct {
	gorm.Model
	Name	string		`json:"supplier"`
}

type SupplierEmail struct {
	gorm.Model
	Name	string		`json:"email"`
}

type DateInvoice struct {
	gorm.Model
	Name	string
	Date	time.Time	`json:"invoicedate"`
}

type DatePay struct {
	gorm.Model
	Name	string
	Date	time.Time	`json:"estpaydate"`
}

func (d *DateInvoice) AfterCreate(scope *gorm.Scope) error {
	if (!d.Date.IsZero()) {
		d.Name = d.Date.Format("2006-01-02 15:04")
	}
	return nil
}

func (d *DateInvoice) BeforeUpdate(scope *gorm.Scope) error {
	if (!d.Date.IsZero()) {
		d.Name = d.Date.Format("2006-01-02 15:04")
	}
	return nil
}

func (d *DatePay) AfterCreate(scope *gorm.Scope) error {
	if (!d.Date.IsZero()) {
		d.Name = d.Date.Format("2006-01-02 15:04")
	}
	return nil
}

func (d *DatePay) BeforeUpdate(scope *gorm.Scope) error {
	if (!d.Date.IsZero()) {
		d.Name = d.Date.Format("2006-01-02 15:04")
	}
	return nil
}