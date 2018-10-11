package invoice

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/earlybird-SynchClient/ApiSetup/models/invoice"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/app/api"
)

type InvoiceData struct {
	StatusInvoice 	string	`json:"invoicestatus"`
	VendorCode 		string	`json:"vendorcode"`
	DateInvoice 	string	`json:"invoicedate"`
	DatePay 		string	`json:"estpaydate"`
	NumberInvoice 	string	`json:"invoiceno`
	AmountInvoice 	string	`json:"invoiceamount"`
}

func invoiceGET(w http.ResponseWriter, r *http.Request) {

	_index := mux.Vars(r)["index"]

	var _invoice invoice.Invoice
	if err := database.Conn.Where("url = ?", _index).
		Preload("VendorCode").
		Preload("StatusInvoice").
		Preload("DateInvoice").
		Preload("DatePay").
		Preload("NumberInvoice").
		Preload("AmountInvoice").
		Preload("Market").
		Find(&_invoice).Error; err != nil {
		api.BadRequest(w)
		return
	}
	if (_invoice.ID == 0){
		api.BadRequest(w)
		return
	}

	response_fail := struct {
		Success bool       	`json:"success"`
		Message string     	`json:"message"`
	}{
		Success: false,
		Message: "token is invalid",
	}

	keys, ok := r.URL.Query()["Token"]
	if !ok || len(keys[0]) < 1 || keys[0] != _invoice.AuthCode{
		api.ServeJSON(w, response_fail)
		return
	}

	var ret InvoiceData
	ret.DatePay = _invoice.DatePay.Name
	ret.VendorCode = _invoice.VendorCode.Name
	ret.StatusInvoice = _invoice.StatusInvoice.Name
	ret.AmountInvoice = _invoice.AmountInvoice.Name
	ret.NumberInvoice = _invoice.NumberInvoice.Name
	ret.DateInvoice = _invoice.DateInvoice.Name

	response := struct {
		Success bool       	`json:"success"`
		Message string     	`json:"message"`
		Data    InvoiceData 		`json:"data"`
	}{
		Success: true,
		Message: "get successfully",
		Data:    ret,
	}

	api.ServeJSON(w, response)
}
