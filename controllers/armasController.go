package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ArmaGet godoc
//
// @Router /armas/{id} [get]
// @Tags Armas
// @Summary trae un documento tipo arma
// @Description Este método devuelve un solo documento tipo arma según el id proporcionado para la búsqueda
// @Param id path string true "ID de arma"
// @Accept json
// @Produce json
// @Success 200 {object} models.ArmasModel "Documento tipo arma"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetArmaController(c echo.Context) error {
	arma_id := c.Param("id") //Obtener parámetro {id} de la URI
	arma_ObjectID, err := primitive.ObjectIDFromHex(arma_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.ArmasModel
	err = result.GetSingleObject(arma_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// ArmasGet godoc
//
// @Router /armas/ [get]
// @Tags Armas
// @Summary trae una colección de documentos tipo Armas
// @Description Este método devuelve una colección de documento tipo Armas según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.ArmasModel "Documento tipo Armas"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetArmasController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.ArmasModel
	var result []models.ArmasModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}
