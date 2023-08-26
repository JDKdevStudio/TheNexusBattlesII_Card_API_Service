package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ArmaduraGet godoc
//
// @Router /armaduras/{id} [get]
// @Tags Armaduras
// @Summary trae un documento tipo Armadura
// @Description Este método devuelve un solo documento tipo Armadura según el id proporcionado para la búsqueda
// @Param id path string true "ID de Armadura"
// @Accept json
// @Produce json
// @Success 200 {object} models.ArmadurasModel "Documento tipo Armadura"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetArmaduraController(c echo.Context) error {
	armadura_id := c.Param("id") //Obtener parámetro {id} de la URI
	armadura_ObjectID, err := primitive.ObjectIDFromHex(armadura_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.ArmadurasModel
	err = result.GetSingleObject(armadura_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// ArmadurasGet godoc
//
// @Router /armaduras/ [get]
// @Tags Armaduras
// @Summary trae una colección de documentos tipo Armaduras
// @Description Este método devuelve una colección de documento tipo Armaduras según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.ArmadurasModel "Documento tipo Armadura"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetArmadurasController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.ArmadurasModel
	var result []models.ArmadurasModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}
