package migrations

import (
	"github.com/qor/auth/auth_identity"

	"github.com/jinzhu/gorm"
	"github.com/earlybird-SynchClient/ApiSetup/models/admins"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
	"github.com/earlybird-SynchClient/ApiSetup/models/invoice"
	"github.com/earlybird-SynchClient/ApiSetup/models/supplier"
	"time"
	"math/rand"
)

func Migrate(db *gorm.DB) {
	autoMigrate(&auth_identity.AuthIdentity{})
	autoMigrate(&admins.Admin{})

	autoMigrate(&basic.VendorCode{})
	autoMigrate(&basic.AmountInvoice{})
	autoMigrate(&basic.NumberInvoice{})
	autoMigrate(&basic.StatusInvoice{})
	autoMigrate(&basic.Supplier{})
	autoMigrate(&basic.SupplierEmail{})
	autoMigrate(&basic.DateInvoice{})
	autoMigrate(&basic.DatePay{})
	autoMigrate(&market.Market{})
	autoMigrate(&invoice.Invoice{})
	autoMigrate(&supplier.Supply{})

	rand.Seed(time.Now().UnixNano())
}

// autoMigrate runs automigrate on provided objects
func autoMigrate(values ...interface{}) {
	for _, value := range values {
		database.Conn.AutoMigrate(value)
	}
}
