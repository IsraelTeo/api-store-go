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

// Función para conectar a la base de datos
func Connection() error {
	// Obtener los valores de la configuración cargados en InitConfig()
	config := config.Envs

	// Crear la cadena DSN (Data Source Name) para la conexión
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBAddress,
		config.DBName)

	// Imprimir la cadena de conexión para verificar
	fmt.Printf("Connecting to DB with DSN: %s\n", DSN)

	// Conexión a la base de datos utilizando GORM
	var err error
	if GDB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{}); err != nil {
		return err // Si hay un error, se devuelve
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
