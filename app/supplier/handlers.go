package supplier

import (
	"github.com/qor/render"
	"net/http"
	"github.com/earlybird-SynchClient/ApiSetup/utils"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/models/supplier"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
)

type Controller struct {
	View *render.Render
}

func (ctrl Controller) Index(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	marketID, _ := strconv.Atoi(_marketID)

	supplies := make([]supplier.Supply, 0)
	database.Conn.
		Where("market_id = ?", uint(marketID)).
		Preload("VendorCode").
		Preload("Supplier").
		Preload("SupplierEmail").
		Preload("Market").
		Find(&supplies)

	ctrl.View.Execute("supplier", map[string]interface{}{
		"Supplies": supplies,
	}, r, w)
}

func (ctrl Controller) Edit(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	_index := mux.Vars(r)["index"]
	index, _ := strconv.Atoi(_index)
	_supply, err := supplier.GetSupplyByID(uint(index))
	if err != nil {
		http.Redirect(w, r, "/"+_marketID+"/supplier/", http.StatusSeeOther)
		return
	}

	codes := make([]basic.VendorCode, 0)
	utils.GetDB(r).Find(&codes)
	suppliers := make([]basic.Supplier, 0)
	utils.GetDB(r).Find(&suppliers)
	emails := make([]basic.SupplierEmail, 0)
	utils.GetDB(r).Find(&emails)

	ctrl.View.Execute("edit", map[string]interface{}{
		"Supply": _supply,
		"VendorCodes": codes,
		"Suppliers": suppliers,
		"SupplierEmails": emails,
	}, r, w)
}

func GetValue(r *http.Request, key string) uint {
	_value := r.FormValue(key)
	value, _ := strconv.Atoi(_value)
	return uint(value)
}

func (ctrl Controller) Register_edit(w http.ResponseWriter, r *http.Request) {
	_marketID := mux.Vars(r)["id"]
	codeID := GetValue(r,"VendorCode")
	supplierID := GetValue(r,"Supplier")
	emailID := GetValue(r,"Email")
	id := GetValue(r,"id")

	if id != 0 {
		_supplier, err := supplier.GetSupplyByID(id)
		if err != nil {
			http.Redirect(w, r, "/"+_marketID+"/supplier/", http.StatusSeeOther)
			return
		}

		_supplier.SupplierEmailID = emailID
		_supplier.SupplierID = supplierID
		_supplier.VendorCodeID = codeID
		err = database.Conn.Save(&_supplier).Error
	}

	http.Redirect(w, r, "/"+_marketID+"/supplier/", http.StatusSeeOther)
}