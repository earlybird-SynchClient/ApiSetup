package admin

import (
	"html/template"

	"github.com/qor/admin"
)

func initFuncMap(Admin *admin.Admin) {
	Admin.RegisterFuncMap("render_market_list", render_market_list)
}

func render_market_list(context *admin.Context) template.HTML {
	var orderContext = context.NewResourceContext("Market")
	orderContext.Searcher.Pagination.PerPage = 5
	// orderContext.SetDB(orderContext.GetDB().Where("state in (?)", []string{"paid"}))

	if orders, err := orderContext.FindMany(); err == nil {
		return orderContext.Render("index/table", orders)
	}
	return template.HTML("")
}