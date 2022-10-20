package main

import (
	"assignment/app/routers"
	"assignment/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

// @title Users API
// @description Sample API Spec for Users
// @version v1.0
// @termsOfService https://9gag.com
// @BasePath /
// @host localhost:1337
// @contact.name Hugos
// @contact.email hu@go.com
func main() {
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.InitRouter()
}
