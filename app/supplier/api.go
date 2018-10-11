package supplier

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/app/api"
	"github.com/earlybird-SynchClient/ApiSetup/models/supplier"
)

type InvoiceData struct {
	VendorCode 		string	`json:"vendorcode"`
	Supplier 		string	`json:"supplier"`
	Email 			string	`json:"email"`
}

func supplierGET(w http.ResponseWriter, r *http.Request) {

	_index := mux.Vars(r)["index"]

	var _supply supplier.Supply
	if err := database.Conn.Where("url = ?", _index).
		Preload("VendorCode").
		Preload("Supplier").
		Preload("SupplierEmail").
		Preload("Market").
		Find(&_supply).Error; err != nil {
		api.BadRequest(w)
		return
	}
	if (_supply.ID == 0){
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
	if !ok || len(keys[0]) < 1 || keys[0] != _supply.AuthCode{
		api.ServeJSON(w, response_fail)
		return
	}

	var ret InvoiceData
	ret.VendorCode = _supply.VendorCode.Name
	ret.Supplier = _supply.Supplier.Name
	ret.Email = _supply.SupplierEmail.Name

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
