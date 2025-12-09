package config

import (
	// Libreria para registrar mensajes en la consola
	// registra errores, eventos importantes, detiene el programa cuando hay errores criticos
	"log"

	// Paquete que contiene la estructura de datos Depto
	"github.com/SantosFarias10/deptoHunter/models"

	// driver que permite a gorm conectarse y trabajar con bases de datos sqlite
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

// Variable global que almacena la conexion a la base de datos
var DB *gorm.DB

func ConnectDB() {
	var err error

	configuration := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// conectamos a la base de datos, gorm.Open() abre una conexion a la base de datos y retorna si todo sale bien (conexion, nil), en caso contrario (nil, error)
	DB, err = gorm.Open(sqlite.Open("base_de_datos.db"), configuration)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	// AutoMigrate crea las tablas si no existe, si ya existen, no hace nada
	err = DB.AutoMigrate(&models.Depto{}, &models.HistorialPrecio{})
	if err != nil {
		log.Fatal("Error migrando la base de datos: ", err)
	}

	log.Println("Base de datos conectada y tablas creadas exitosamente")
}
