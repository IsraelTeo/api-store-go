package db

import (
	"github.com/IsraelTeo/api-store-go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsnString = "root:@tcp(localhost:3306)/api_store_go"
var GDB *gorm.DB

func Connection() error {
	var err error
	GDB, err = gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDB() error {
	err := gdb.AutoMigrate(&model.Product{}, &model.Customer{}, &model.Sale{})
	if err != nil {
		return err
	}
	return nil
}