// @APIVersion 1.0.0
// @APITitle SIGMA API
// @APIDescription API for SIGMA app.
// @BasePath http://host:port/api
package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/qor/admin"
	"github.com/qor/location"
	"github.com/qor/media"
	"github.com/earlybird-SynchClient/ApiSetup/paola/auth"
	"github.com/earlybird-SynchClient/ApiSetup/paola/database"
	"github.com/earlybird-SynchClient/ApiSetup/paola/jsonconfig"
	"github.com/earlybird-SynchClient/ApiSetup/paola/server"

	"github.com/earlybird-SynchClient/ApiSetup/models/migrations"
	"github.com/earlybird-SynchClient/ApiSetup/models/admins"

	adminapp "github.com/earlybird-SynchClient/ApiSetup/app/admin"
	"github.com/earlybird-SynchClient/ApiSetup/app/application"
	"github.com/earlybird-SynchClient/ApiSetup/app/static"
	"github.com/earlybird-SynchClient/ApiSetup/app/basic"
	"github.com/earlybird-SynchClient/ApiSetup/app/invoice"
	"github.com/earlybird-SynchClient/ApiSetup/app/supplier"
)

var (
	config = &Configuration{}
)

type Configuration struct {
	Server   server.Server           `json:"server"`
	Database database.DatabaseConfig `json:"database"`
}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadComponents parses the JSON file and sets up components
func LoadComponents() *Configuration {
	configPath := "config.json"
	jsonconfig.Load(configPath, config)

	database.Configure(config.Database)
	// Configure database rds or local test
	host := os.Getenv("RDS_HOSTNAME")
	port := os.Getenv("RDS_PORT")
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	db := os.Getenv("RDS_DB_NAME")

	if host != "" {
		configDatabase := database.DatabaseConfig{}
		portInt, _ := strconv.Atoi(port)

		configDatabase.Host = host
		configDatabase.Port = portInt
		configDatabase.User = username
		configDatabase.Password = password
		configDatabase.Database = db

		database.Configure(configDatabase)
	}
	database.Connect()

	media.RegisterCallbacks(database.Conn)

	// Update the database
	migrations.Migrate(database.Conn)

	location.GoogleAPIKey = os.Getenv("GOOGLE_API_KEY")

	return config
}

func main() {
	LoadComponents()

	Admin := admin.New(&admin.AdminConfig{
		SiteName: "MAPPING API",
		DB:       database.Conn,
		Auth:     auth.AdminAuth{},
	})

	app := application.New(&application.Config{
		Admin: Admin,
		Auth:  auth.NewAuth(),
	})

	// check if there are admins, if not create one
	adminUsers, _ := admins.GetAdminUsers()
	if len(adminUsers) == 0 {
		adm := admins.Admin{
			Name:     "Admin",
			Email:    "admin@admin.com",
			Password: "admin123!",
			Role:     "Admin",
		}
		database.Conn.Create(&adm)
	}

	app.Use(adminapp.New(&adminapp.Config{}))
	app.Use(basic.New(&basic.Config{}))
	app.Use(invoice.New(&invoice.Config{}))
	app.Use(supplier.New(&supplier.Config{}))
	app.Use(static.New(&static.Config{}))

	server.Run(app.GetMux(), config.Server)
}
