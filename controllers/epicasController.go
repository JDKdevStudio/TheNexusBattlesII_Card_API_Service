package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EpicaeGet godoc
//
// @Router /epicas/{id} [get]
// @Tags Epicas
// @Summary trae un documento tipo Epica
// @Description Este método devuelve un solo documento tipo héroe según el id proporcionado para la búsqueda
// @Param id path string true "ID de"
// @Accept json
// @Produce json
// @Success 200 {object} models.EpicasModel "Documento tipo Epica"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetEpicaController(c echo.Context) error {
	heroe_id := c.Param("id") //Obtener parámetro {id} de la URI
	Epica_ObjectID, err := primitive.ObjectIDFromHex(heroe_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.EpicasModel
	err = result.GetSingleObject(Epica_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// EpicasGet godoc
//
// @Router /epicas/ [get]
// @Tags Epicas
// @Summary trae una colección de documentos tipo Epicas
// @Description Este método devuelve una colección de documento tipo Epicas según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.EpicasModel "Documento tipo Epicas"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetEpicasController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.EpicasModel
	var result []models.EpicasModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}
