package service

import (
	// Libreria estandar de Go para formatear y mostrar texto en la consola
	"fmt"

	"time"

	"github.com/SantosFarias10/deptoHunter/config"
	"github.com/SantosFarias10/deptoHunter/models"
)

// ProcesarDepto recibe datos del scraper y decide que hacer con ellos
func ProcesarDepto(titulo, url, ubicacion string, habitaciones int, baños int, metros float64, precio float64, moneda string) {
	var departamento models.Depto

	// Buscamos si ya existe la propiedad por su url
	// Where busca, First intenta llenar la variable "departamento"
	resultado := config.DB.Where("url_original = ?", url).First(&departamento)

	if resultado.Error != nil {
		// Caso 1: no existe => la creamos desde cero
		newDepto := models.Depto{
			Titulo:       titulo,
			UrlOriginal:  url,
			Ubicacion:    ubicacion,
			Habitaciones: habitaciones,
			Baños:        baños,
			Metros:       metros,
			HistorialPrecios: []models.HistorialPrecio{
				{Precio: precio, Moneda: moneda, Fecha: time.Now()},
			},
		}
		config.DB.Create(&newDepto)
		fmt.Printf("Nuevo Departament: %s - %s | %.2f\n", titulo, moneda, precio)
	} else {
		// Caso 2: existe, por lo que verificamos si cambio de precio

		huboCambio := false

		if departamento.Habitaciones == 0 && habitaciones > 0 {
			departamento.Habitaciones = habitaciones
			huboCambio = true
		}
		if departamento.Baños == 0 && baños > 0 {
			departamento.Baños = baños
			huboCambio = true
		}
		if (departamento.Ubicacion == "" || departamento.Ubicacion == departamento.Titulo) && ubicacion != "" {
			departamento.Ubicacion = ubicacion
			huboCambio = true
		}

		if huboCambio {
			config.DB.Save(&departamento)
			fmt.Printf("Departamento actualizado: %s\n", departamento.Titulo)
		}

		verificarCambioDePrecio(departamento, precio, moneda)
	}
}

func verificarCambioDePrecio(d models.Depto, precioNuevo float64, moneda string) {
	// obtenemos el ultimo precio registrado
	var ultimoPrecio models.HistorialPrecio
	config.DB.Where("depto_id = ?", d.ID).Order("fecha desc").First(&ultimoPrecio)

	// comparamos
	if ultimoPrecio.Precio != precioNuevo {
		variacion := precioNuevo - ultimoPrecio.Precio

		fmt.Printf("Cambio detectado en %s: Antes %.2f ---> Ahora %.2f. Diferencia de %.2f\n", d.Titulo, ultimoPrecio.Precio, precioNuevo, variacion)

		nuevoPrecio := models.HistorialPrecio{
			DeptoID: d.ID,
			Precio:  precioNuevo,
			Moneda:  moneda,
			Fecha:   time.Now(),
		}

		config.DB.Create(&nuevoPrecio)

		// TODO: Agregar notificaciones
	} else {
		fmt.Printf("Sin cambios en %s\n", d.Titulo)
	}
}
