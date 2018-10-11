package invoice

import (
	"github.com/qor/render"
	"net/http"
	"github.com/earlybird-SynchClient/ApiSetup/models/invoice"
	"github.com/earlybird-SynchClient/ApiSetup/utils"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/models/supplier"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
)

type Controller struct {
	View *render.Render
}

func (ctrl Controller) Index(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	marketID, _ := strconv.Atoi(_marketID)
	if _marketID == ""{
		var markets market.Market
		utils.GetDB(r).First(&markets)
		marketID = int(markets.ID)
	}

	invoices := make([]invoice.Invoice, 0)
	database.Conn.
		Where("market_id = ?", uint(marketID)).
		Preload("VendorCode").
		Preload("StatusInvoice").
		Preload("DateInvoice").
		Preload("DatePay").
		Preload("NumberInvoice").
		Preload("AmountInvoice").
		Preload("Market").
		Find(&invoices)

	ctrl.View.Execute("invoice", map[string]interface{}{
		"Invoices": invoices,
	}, r, w)
}

func (ctrl Controller) Add(w http.ResponseWriter, r *http.Request) {
	codes := make([]basic.VendorCode, 0)
	utils.GetDB(r).Find(&codes)
	numbers := make([]basic.NumberInvoice, 0)
	utils.GetDB(r).Find(&numbers)
	dates := make([]basic.DateInvoice, 0)
	utils.GetDB(r).Find(&dates)
	amounts := make([]basic.AmountInvoice, 0)
	utils.GetDB(r).Find(&amounts)
	status := make([]basic.StatusInvoice, 0)
	utils.GetDB(r).Find(&status)
	pays := make([]basic.DatePay, 0)
	utils.GetDB(r).Find(&pays)
	suppliers := make([]basic.Supplier, 0)
	utils.GetDB(r).Find(&suppliers)
	emails := make([]basic.SupplierEmail, 0)
	utils.GetDB(r).Find(&emails)

	ctrl.View.Execute("add", map[string]interface{}{
		"VendorCodes": codes,
		"InvoiceNos": numbers,
		"InvoiceDates": dates,
		"InvoiceAmounts": amounts,
		"InvoiceStatus": status,
		"EstPayDates": pays,
		"Suppliers": suppliers,
		"Emails": emails,
	}, r, w)
}

func (ctrl Controller) Edit(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	_index := mux.Vars(r)["index"]
	index, _ := strconv.Atoi(_index)
	_invoice, err := invoice.GetInvoiceByID(uint(index))
	if err != nil {
		http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
		return
	}

	codes := make([]basic.VendorCode, 0)
	utils.GetDB(r).Find(&codes)
	numbers := make([]basic.NumberInvoice, 0)
	utils.GetDB(r).Find(&numbers)
	dates := make([]basic.DateInvoice, 0)
	utils.GetDB(r).Find(&dates)
	amounts := make([]basic.AmountInvoice, 0)
	utils.GetDB(r).Find(&amounts)
	status := make([]basic.StatusInvoice, 0)
	utils.GetDB(r).Find(&status)
	pays := make([]basic.DatePay, 0)
	utils.GetDB(r).Find(&pays)

	ctrl.View.Execute("edit", map[string]interface{}{
		"Invoice": _invoice,
		"VendorCodes": codes,
		"InvoiceNos": numbers,
		"InvoiceDates": dates,
		"InvoiceAmounts": amounts,
		"InvoiceStatus": status,
		"EstPayDates": pays,
	}, r, w)
}

func (ctrl Controller) Remove(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	_index := mux.Vars(r)["index"]
	index, _ := strconv.Atoi(_index)

	elm, err := invoice.GetInvoiceByID(uint(index))
	if err != nil {
		http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
		return
	}

	elm1, err1 := supplier.GetSupplyByToken(elm.AuthCode)
	if err1 == nil && elm1.ID != 0 {
		err1 = database.Conn.Delete(&elm1).Error
	}

	err = database.Conn.Delete(&elm).Error
	http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
	return
}

func GetValue(r *http.Request, key string) uint {
	_value := r.FormValue(key)
	value, _ := strconv.Atoi(_value)
	return uint(value)
}

func (ctrl Controller) Register_add(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	marketID, _ := strconv.Atoi(_marketID)
	codeID := GetValue(r,"VendorCode")
	numberID := GetValue(r,"InvoiceNo")
	dateID := GetValue(r,"InvoiceDate")
	amountID := GetValue(r,"InvoiceAmount")
	statusID := GetValue(r,"InvoiceStatus")
	payID := GetValue(r,"EstPayDate")
	supplierID := GetValue(r,"Supplier")
	emailID := GetValue(r,"Email")

	var _invoice invoice.Invoice
	_invoice.AmountInvoiceID = amountID
	_invoice.StatusInvoiceID = statusID
	_invoice.DatePayID = payID
	_invoice.DateInvoiceID = dateID
	_invoice.NumberInvoiceID = numberID
	_invoice.VendorCodeID = codeID
	_invoice.MarketID = uint(marketID)
	err := database.Conn.Create(&_invoice).Error
	if err != nil {
		http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
		return
	}

	var _supplier supplier.Supply
	_supplier.VendorCodeID = codeID
	_supplier.SupplierEmailID = emailID
	_supplier.SupplierID = supplierID
	_supplier.AuthCode = _invoice.AuthCode
	_supplier.URL = _invoice.URL
	_supplier.MarketID = uint(marketID)
	err = database.Conn.Create(&_supplier).Error

	http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
}

func (ctrl Controller) Register_edit(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	codeID := GetValue(r,"VendorCode")
	numberID := GetValue(r,"InvoiceNo")
	dateID := GetValue(r,"InvoiceDate")
	amountID := GetValue(r,"InvoiceAmount")
	statusID := GetValue(r,"InvoiceStatus")
	payID := GetValue(r,"EstPayDate")
	id := GetValue(r,"id")

	if id != 0 {
		_invoice, err := invoice.GetInvoiceByID(id)
		if err != nil {
			http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
			return
		}

		_invoice.AmountInvoiceID = amountID
		_invoice.StatusInvoiceID = statusID
		_invoice.DatePayID = payID
		_invoice.DateInvoiceID = dateID
		_invoice.NumberInvoiceID = numberID
		_invoice.VendorCodeID = codeID
		err = database.Conn.Save(&_invoice).Error
	}

	http.Redirect(w, r, "/"+_marketID+"/invoice/", http.StatusSeeOther)
}