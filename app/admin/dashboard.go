package admin

import (
	"github.com/qor/admin"
)

// SetupDashboard setup dashboard
func SetupDashboard(Admin *admin.Admin) {
	// Add Dashboard link on menu
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin", Priority: 1})

	initFuncMap(Admin)
}
