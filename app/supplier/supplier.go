package supplier

import (
	"github.com/earlybird-SynchClient/ApiSetup/app/application"
	"github.com/qor/admin"
	"github.com/earlybird-SynchClient/ApiSetup/models/supplier"
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
			&render.Config{AssetFileSystem: application.AssetFS.NameSpace("supplier")},
			"app/supplier/views",
		),
	}
	funcmapmaker.AddFuncMapMaker(controller.View)

	application.Router.Handle("/{id}/supplier/",
		application.Api.MemberMiddleware.ThenFunc(controller.Index))
	application.Router.Handle("/{id}/supplier/edit/{index}",
		application.Api.MemberMiddleware.ThenFunc(controller.Edit))
	application.Router.Handle("/{id}/supplier/register_edit/",
		application.Api.MemberMiddleware.ThenFunc(controller.Register_edit))

	if application.Admin != nil {
		m.ConfigureAdmin(application.Admin)
	}
	m.ConfigureAPI(application.Api)
}

func (m Module) ConfigureAPI(api *api.API) {
	sr := api.Router.PathPrefix("/").Subrouter()

	sr.HandleFunc("/supplier/{index}", supplierGET).Methods("GET")
}

func (m Module) ConfigureAdmin(adm *admin.Admin) {
	adm.AddResource(&supplier.Supply{}, &admin.Config{
		Menu: []string{"Supplier"},
	})
}
