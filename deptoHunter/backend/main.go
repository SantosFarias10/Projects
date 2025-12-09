package main // le decimos a go que este archivo se puede ejecutar

// Importamos las librerias que vamos a usar
import (
	// libreria estandar de go, maneja las peticiones http
	// define codigos de estado http (200 ok, 404 not fount, ...), maneja las requests y responses, permite crear servidores web basicos
	"net/http"

	// framework web
	// facilita la creciacion de APIs REST, maneja endpoints, procesa JSON, incluye middleware (autenticacion, logging, ...)
	"github.com/gin-gonic/gin"

	// Paquete que contiene la conexion a la base de datos
	"github.com/SantosFarias10/deptoHunter/config"
)

func main() {
	// Conectamos a la base de datos
	config.ConnectDB()

	// Inicializamos el server con Logger + Recovery
	// Basicamente inicializamos el router con middleware, si ocurre un error lo captura y evita que el server se caiga
	router := gin.Default()

	// Definimos una ruta GET para la URL "/ping"
	// Cuando entres a la URL, responderá un JSON.
	/* De esta forma gin.H crea diccionarios en Go que se convierten automaticamente a JSON, otra forma de hacer es la de abajo
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Hello World!",
		})
	})
	*/

	type Response struct {
		Mensaje string `json:"mensaje"`
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Mensaje: "Base de Datos conectada!",
		})
	})

	// Arrancamos el server en el puerto 8080
	router.Run(":8080")
}
