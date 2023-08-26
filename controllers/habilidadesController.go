package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HabilidadeGet godoc
//
// @Router /habilidades/{id} [get]
// @Tags Habilidades
// @Summary trae un documento tipo habilidad
// @Description Este método devuelve un solo documento tipo héroe según el id proporcionado para la búsqueda
// @Param id path string true "ID de"
// @Accept json
// @Produce json
// @Success 200 {object} models.HabilidadesModel "Documento tipo habilidad"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetHabilidadController(c echo.Context) error {
	heroe_id := c.Param("id") //Obtener parámetro {id} de la URI
	habilidad_ObjectID, err := primitive.ObjectIDFromHex(heroe_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.HabilidadesModel
	err = result.GetSingleObject(habilidad_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// HabilidadesGet godoc
//
// @Router /habilidades/ [get]
// @Tags Habilidades
// @Summary trae una colección de documentos tipo Habilidades
// @Description Este método devuelve una colección de documento tipo Habilidades según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.HabilidadesModel "Documento tipo Habilidades"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetHabilidadesController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.HabilidadesModel
	var result []models.HabilidadesModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}
