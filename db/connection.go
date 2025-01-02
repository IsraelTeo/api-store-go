package db

import (
	"fmt"

	"github.com/IsraelTeo/api-store-go/config"
	"github.com/IsraelTeo/api-store-go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declara una variable llamada GDB de tipo puntero a gorm.DB.
// DB: Es la estructura principal que maneja la conexión a la base de datos en GORM.
// Esta estructura es responsable de realizar consultas y transacciones en la base de datos.
var GDB *gorm.DB

func Connection(cfg *config.Config) error {

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	fmt.Printf("Connecting to DB with DSN: %s\n", DSN)

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
