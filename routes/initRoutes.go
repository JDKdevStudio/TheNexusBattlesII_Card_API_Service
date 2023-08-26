package routes

import (
	"TheNexusBattlesII_Card_API_Service/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(server *echo.Echo) {
	server.GET("/", controllers.RootController)      //Status del servidor
	server.GET("/api/", controllers.IndexController) //Index del servidor
	server.GET("/docs/*", echoSwagger.WrapHandler)   //Iniciar el servicio de Swagger API
	//Heroes collection
	server.GET("/api/heroes/", controllers.GetheroesController)
	server.GET("/api/heroes/:id", controllers.GetheroeController)
	server.POST("/api/heroes/", controllers.CreateHeroeController)
	server.PATCH("/api/heroes/:id", controllers.UpdateHeroeController)
	//Habilidades collection
	server.GET("/api/habilidades/:id", controllers.GetHabilidadController)
}
