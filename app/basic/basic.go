package basic

import (
	"github.com/qor/admin"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
	"github.com/earlybird-SynchClient/ApiSetup/app/application"
)

type Module struct {
	Config *Config
}

func New(config *Config) *Module {
	return &Module{Config: config}
}

type Config struct{}

// Configure blog module
func (m Module) Configure(application *application.Application) {
	if application.Admin != nil {
		m.ConfigureAdmin(application.Admin)
	}
}

func (m Module) ConfigureAdmin(adm *admin.Admin) {
	adm.AddResource(&market.Market{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.VendorCode{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.NumberInvoice{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.AmountInvoice{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.StatusInvoice{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.Supplier{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	adm.AddResource(&basic.SupplierEmail{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	invoice := adm.AddResource(&basic.DateInvoice{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	invoice.IndexAttrs("ID", "Date")
	invoice.EditAttrs("ID", "Date")
	invoice.NewAttrs("ID", "Date")
	pay := adm.AddResource(&basic.DatePay{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	pay.IndexAttrs("ID", "Date")
	pay.EditAttrs("ID", "Date")
	pay.NewAttrs("ID", "Date")
}

