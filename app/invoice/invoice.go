package invoice

import (
	"github.com/qor/admin"
	"github.com/earlybird-SynchClient/ApiSetup/app/application"
	"github.com/earlybird-SynchClient/ApiSetup/models/invoice"
	"github.com/qor/qor"
	"fmt"
	"github.com/earlybird-SynchClient/ApiSetup/models/basic"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/qor/render"
	"github.com/earlybird-SynchClient/ApiSetup/utils/funcmapmaker"
	"github.com/earlybird-SynchClient/ApiSetup/app/api"
)

type Module struct {
	Config *Config
}

func New(config *Config) *Module {
	return &Module{Config: config}
}

type Config struct{}

func (m Module) Configure(application *application.Application) {
	controller := &Controller{
		View: render.New(
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("invoice")},
			"app/invoice/views",
		),
	}
	funcmapmaker.AddFuncMapMaker(controller.View)
	application.Router.Handle("/",
		application.Api.MemberMiddleware.ThenFunc(controller.Index))
	application.Router.Handle("/{id}/invoice/",
		application.Api.MemberMiddleware.ThenFunc(controller.Index))
	application.Router.Handle("/{id}/invoice/add/",
		application.Api.MemberMiddleware.ThenFunc(controller.Add))
	application.Router.Handle("/{id}/invoice/remove/{index}",
		application.Api.MemberMiddleware.ThenFunc(controller.Remove))
	application.Router.Handle("/{id}/invoice/edit/{index}",
		application.Api.MemberMiddleware.ThenFunc(controller.Edit))
	application.Router.Handle("/{id}/invoice/register_add/",
		application.Api.MemberMiddleware.ThenFunc(controller.Register_add))
	application.Router.Handle("/{id}/invoice/register_edit/",
		application.Api.MemberMiddleware.ThenFunc(controller.Register_edit))

	if application.Admin != nil {
		m.ConfigureAdmin(application.Admin)
	}
	m.ConfigureAPI(application.Api)
}

func (m Module) ConfigureAPI(api *api.API) {
	sr := api.Router.PathPrefix("/").Subrouter()

	sr.HandleFunc("/invoice/{index}", invoiceGET).Methods("GET")
}

func (m Module) ConfigureAdmin(adm *admin.Admin) {
	invoices := adm.AddResource(&invoice.Invoice{}, &admin.Config{
		Menu: []string{"Invoice"},
	})
	invoices.Meta(&admin.Meta{
		Name: "DateInvoice",
		Type: "select_one",
		Config: &admin.SelectOneConfig{
			Collection: func(value interface{}, context *qor.Context) [][]string {
				var collectionValues [][]string
				var _invoices []*basic.DateInvoice
				database.Conn.Find(&_invoices)
				for _, _invoice := range _invoices {
					collectionValues = append(collectionValues, []string{fmt.Sprint(_invoice.ID), _invoice.Name})
				}
				return collectionValues
			},
		},
	})
	invoices.Meta(&admin.Meta{
		Name: "DatePay",
		Type: "select_one",
		Config: &admin.SelectOneConfig{
			Collection: func(value interface{}, context *qor.Context) [][]string {
				var collectionValues [][]string
				var _pays []*basic.DatePay
				database.Conn.Find(&_pays)
				for _, _pay := range _pays {
					collectionValues = append(collectionValues, []string{fmt.Sprint(_pay.ID), _pay.Name})
				}
				return collectionValues
			},
		},
	})

	invoices.EditAttrs("ID", "AuthCode", "URL", "StatusInvoice", "VendorCode", "DateInvoice", "DatePay", "NumberInvoice", "AmountInvoice", "Market")
	invoices.NewAttrs("ID", "StatusInvoice", "VendorCode", "DateInvoice", "DatePay", "NumberInvoice", "AmountInvoice", "Market")
}
