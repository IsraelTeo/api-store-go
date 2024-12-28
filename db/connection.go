package db

import (
	"fmt"
	"os"

	"github.com/IsraelTeo/api-store-go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func Connection() error {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Printf("%s", DSN)
	var err error
	if GDB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{}); err != nil {
		return err
	}

	return nil
}

func MigrateDB() error {
	err := GDB.AutoMigrate(
		&model.Product{},
		&model.Customer{},
		&model.Sale{},
		&model.User{},
		&model.SaleProduct{},
	)
	if err != nil {
		return err
	}
	return nil
}
