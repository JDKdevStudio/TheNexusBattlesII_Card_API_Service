package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HeroeGet godoc
//
// @Router /heroes/{id} [get]
// @Tags Heroes
// @Summary trae un documento tipo heroes
// @Description Este método devuelve un solo documento tipo héroe según el id proporcionado para la búsqueda
// @Param id path string true "ID del heroe"
// @Accept json
// @Produce json
// @Success 200 {object} models.HeroesModel "Documento tipo Héroe"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetheroeController(c echo.Context) error {
	heroe_id := c.Param("id") //Obtener parámetro {id} de la URI
	heroe_ObjectID, err := primitive.ObjectIDFromHex(heroe_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.HeroesModel
	err = result.GetsingeObject(heroe_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// HeroesGet godoc
//
// @Router /heroes/ [get]
// @Tags Heroes
// @Summary trae una colección de documentos tipo heroes
// @Description Este método devuelve una colección de documento tipo héroe según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.HeroesModel "Documento tipo Héroe"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetheroesController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.HeroesModel
	var result []models.HeroesModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}

// HeroePost godoc
//
// @Router /heroes/ [post]
// @Tags Heroes
// @Summary crea un documento en la colección tipo heroe
// @Description Este método crea un documento en la colección de heroes
// @Param urlImagen formData string true "URL de la imagen de la carta"
// @Param clase formData string true "clase del heroe"
// @Param tipo formData string true "tipo de la carta"
// @Param poder formData int true "poder del heroe"
// @Param vida formData int true "vida del heroe"
// @Param defensa formData int true "defensa del heroe"
// @Param ataqueBase formData int true "ataque base del heroe"
// @Param ataqueDado formData int true "probabilidad de ataque del heroe"
// @Param danoMax formData int true "daño máximo del heroe"
// @Param activo formData bool false "estado de la carta"
// @Param desc formData string true "descripción de la carta"
// @Accept x-www-form-urlencoded
// @Produce plain
// @Success 200 {object} string     "Documento tipo Héroe"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func CreateHeroeController(c echo.Context) error {
	var formData models.HeroesModel
	err := c.Bind(&formData)
	if err != nil || !formData.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	err = formData.PostSingleObject()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")

	}
	return c.String(http.StatusOK, "OK")
}

// HeroePatch godoc
//
// @Router /heroes/{id} [patch]
// @Tags Heroes
// @Summary actualiza un documento en la colección tipo heroe
// @Description Este método crea un documento en la colección de heroes
// @Param id path string true "Id de la carta"
// @Param urlImagen formData string false "URL de la imagen de la carta"
// @Param clase formData string false "clase del heroe"
// @Param tipo formData string false "tipo de la carta"
// @Param poder formData int false "poder del heroe"
// @Param vida formData int false "vida del heroe"
// @Param defensa formData int false "defensa del heroe"
// @Param ataqueBase formData int false "ataque base del heroe"
// @Param ataqueDado formData int false "probabilidad de ataque del heroe"
// @Param danoMax formData int false "daño máximo del heroe"
// @Param activo formData bool false "estado de la carta"
// @Param desc formData string false "descripción de la carta"
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} models.HeroesModel "Documento tipo Héroe"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func UpdateHeroeController(c echo.Context) error {
	heroe_id := c.Param("id")
	heroe_ObjectID, err := primitive.ObjectIDFromHex(heroe_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.HeroesModel
	err = c.Bind(&result)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	err = result.UpdateSingleObject(heroe_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")

	}
	return c.String(http.StatusOK, "OK")
}
