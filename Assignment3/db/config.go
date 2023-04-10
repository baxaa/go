package db

import (
	"fmt"
	"github.com/baxaa/e-commerce/Helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "uhupup667"
	dbName   = "books"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	Helper.Check(err)
}
