package models // importamos la estructura de datos models

import (
	// paquete para manejar fechas
	"time"

	// ORM para bases de datos, o sea mapea objetos go a tablas de bases de datos
	// crea, lee, actualiza y borra datos sin escribir SQL, maneja relaciones entre tablas, soporta tanto PostgreSQL, MySQL, SQLite, ...
	"gorm.io/gorm"
)

// Definimos que datos FIJOS tiene un Departamento
type Depto struct {
	// Heredamos los campos ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model

	Titulo       string  `gorm:"size:256;not null" json:"titulo"`          //basicamente le decimos que el campo titulo tenga un maximo de 256 caracteres (256 porque wolovick me traumo) y "not null" le deci que el campo es obligatorio
	UrlOriginal  string  `gorm:"uniqueIndex;not null" json:"url_original"` // uniqueIndex es igual a unique (evita que haya URLs duplicadas) pero uniqueIndex es mas rapida en la busqueda de tablas
	Ubicacion    string  `json:"ubicacion"`
	Habitaciones int     `json:"habitaciones"`
	Baños        int     `json:"baños"`
	Metros       float64 `json:"metros"`

	// Relacion: Un departamento tiene muchos precios
	// creamos un array de HistorialPrecio (tipo de dato)
	// le decimos a gorm que la clave foranea es DeptoID, una clave foranea es una clave que apunta a otra tabla
	HistorialPrecios []HistorialPrecio `gorm:"foreignKey:DeptoID"`
}

type HistorialPrecio struct {
	gorm.Model

	DeptoID uint    `gorm:"not null"`
	Precio  float64 `gorm:"not null"`
	Fecha   time.Time
	Moneda  string `gorm:"size:10;not null" json:"moneda"`
}
