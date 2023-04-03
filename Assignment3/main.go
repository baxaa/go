package main

import (
	"github.com/baxaa/e-commerce/Router"
	"github.com/baxaa/e-commerce/db"
	"github.com/baxaa/e-commerce/model"
)

func main() {
	db.DatabaseConnection()
	db.DB.AutoMigrate(model.Book{})
	Router.Routes()
}
