package main

import (
	"TheNexusBattlesII_Card_API_Service/config"
	_ "TheNexusBattlesII_Card_API_Service/docs"
	"TheNexusBattlesII_Card_API_Service/routes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title           The Nexus Battles II: Card API Service
// @version         1.0
// @description     Servicio API para acceder a las cartas del negocio
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   https://opensource.org/license/mit/

// @host      prime.bucaramanga.upb.edu.co
// @BasePath  /api

// @externalDocs.description  Repository
// @externalDocs.url          https://github.com/JDKdevStudio/TheNexusBattlesII_Card_API_Service

func main() {
	config.Init()        //Iniciar configuración preliminar
	server := echo.New() //Iniciar instancia del servidor
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Permitir todas las IPs
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch},
	}))
	routes.InitRoutes(server) //Registrar rutas del servicio
	if err := server.Start(":80"); err != nil {
		log.Fatal(err)
	}
}
