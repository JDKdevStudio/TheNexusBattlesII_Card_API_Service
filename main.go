package main

import (
	"TheNexusBattlesII_Card_API_Service/config"
	_ "TheNexusBattlesII_Card_API_Service/docs"
	"TheNexusBattlesII_Card_API_Service/routes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

// @title           The Nexus Battles II: Card API Service
// @version         1.0
// @description     Servicio API para acceder a las cartas del negocio
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   https://opensource.org/license/mit/

// @host      localhost
// @BasePath  /api

// @externalDocs.description  Repository
// @externalDocs.url          https://github.com/JDKdevStudio/TheNexusBattlesII_Card_API_Service

func main() {
	config.Init()        //Iniciar configuración preliminar
	server := echo.New() //Iniciar instancia del servidor
	server.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                               // Permitir todas las IPs
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch}, //Métodos de la API
	}))
	routes.InitRoutes(server) //Registrar rutas del servicio
	go redirect()
	if err := server.StartAutoTLS(":443"); err != nil {
		log.Fatal(err)
	}
}

func redirect() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Start(":80")
}
