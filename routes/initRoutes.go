package routes

import (
	"TheNexusBattlesII_Card_API_Service/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(server *echo.Echo) {
	//Root
	server.GET("/", controllers.RootController)           //Status del servidor
	server.GET("/api/all/", controllers.GetAllController) // Endpoint privado para el Ecommerce
	server.GET("/api/", controllers.IndexController)      //Index del servidor
	server.GET("/docs/*", echoSwagger.WrapHandler)        //Iniciar el servicio de Swagger API
	//Armaduras collection
	server.GET("/api/armaduras/", controllers.GetArmadurasController)   //Obtener colección armaduras por paginación
	server.GET("/api/armaduras/:id", controllers.GetArmaduraController) //Obtener documento armadura por id
	//Armas collection
	server.GET("/api/armas/", controllers.GetArmasController)
	server.GET("/api/armas/:id", controllers.GetArmaController)
	//Epicas collection
	server.GET("/api/epicas/", controllers.GetEpicasController)
	server.GET("/api/epicas/:id", controllers.GetEpicaController)
	//Heroes collection
	server.GET("/api/heroes/", controllers.GetheroesController)
	server.GET("/api/heroes/:id", controllers.GetheroeController)
	server.POST("/api/heroes/", controllers.CreateHeroeController)
	server.PATCH("/api/heroes/:id", controllers.UpdateHeroeController)
	//Items collection
	server.GET("/api/items/", controllers.GetItemssController)
	server.GET("/api/items/:id", controllers.GetItemController)
}
