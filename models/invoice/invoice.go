package invoice

import (
	"github.com/jinzhu/gorm"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"math/rand"
	"fmt"
)

type Invoice struct {
	gorm.Model

	AuthCode string
	URL string

	StatusInvoiceID uint
	StatusInvoice *basic.StatusInvoice	`json:"invoicestatus"`
	VendorCodeID uint
	VendorCode *basic.VendorCode		`json:"vendorcode"`
	DateInvoiceID uint
	DateInvoice *basic.DateInvoice		`json:"invoicedate"`
	DatePayID uint
	DatePay *basic.DatePay				`json:"estpaydate"`
	NumberInvoiceID uint
	NumberInvoice *basic.NumberInvoice	`json:"invoiceno`
	AmountInvoiceID uint
	AmountInvoice *basic.AmountInvoice	`json:"invoiceamount"`
	MarketID uint
	Market *market.Market				`json:"invoicestatus"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetInvoiceByToken(token string) (*Invoice, error) {
	var _invoce Invoice
	err := database.Conn.
		Where("auth_code = ?", token).
		Find(&_invoce).
		Error

	return &_invoce, err
}

func GetInvoiceByID(id uint) (*Invoice, error) {
	var _invoce Invoice
	err := database.Conn.
		Where("id = ?", id).
		Preload("VendorCode").
		Preload("StatusInvoice").
		Preload("DateInvoice").
		Preload("DatePay").
		Preload("NumberInvoice").
		Preload("AmountInvoice").
		Preload("Market").
		Find(&_invoce).
		Error

	return &_invoce, err
}

func UpdateEntry(i *Invoice) bool{
	var _elm1 basic.StatusInvoice
	database.Conn.Where("id = ?", i.StatusInvoiceID).Find(&_elm1)
	i.StatusInvoice = &_elm1

	var _elm2 basic.VendorCode
	database.Conn.Where("id = ?", i.VendorCodeID).Find(&_elm2)
	i.VendorCode = &_elm2

	var _elm3 basic.NumberInvoice
	database.Conn.Where("id = ?", i.NumberInvoiceID).Find(&_elm3)
	i.NumberInvoice = &_elm3

	var _elm4 basic.AmountInvoice
	database.Conn.Where("id = ?", i.AmountInvoiceID).Find(&_elm4)
	i.AmountInvoice = &_elm4

	var _elm5 market.Market
	database.Conn.Where("id = ?", i.MarketID).Find(&_elm5)
	i.Market = &_elm5

	var _elm6 basic.DatePay
	database.Conn.Where("id = ?", i.DatePayID).Find(&_elm6)
	i.DatePay = &_elm6

	var _elm7 basic.DateInvoice
	database.Conn.Where("id = ?", i.DateInvoiceID).Find(&_elm7)
	i.DateInvoice = &_elm7

	return true;
}

func GetToken() string{
	var retry = 10

	for i := 0; i < retry; i++ {
		_frist := rand.Intn(9999)
		_second := rand.Intn(999)
		_token := fmt.Sprintf("%4d-1-%3d", _frist, _second)
		_invoce, _ := GetInvoiceByToken(_token)
		if _invoce.ID == 0{
			return _token;
		}
	}
	return "NONE"
}

var urlLen = 20
func (i *Invoice) BeforeCreate(scope *gorm.Scope) error {
	url := RandSeq(urlLen)
	i.URL = url
	i.AuthCode = GetToken()
	UpdateEntry(i)

	return nil
}

func (i *Invoice) BeforeUpdate(scope *gorm.Scope) error {
	UpdateEntry(i)

	return nil
}