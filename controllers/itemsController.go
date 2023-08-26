package controllers

import (
	"TheNexusBattlesII_Card_API_Service/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ItemGet godoc
//
// @Router /items/{id} [get]
// @Tags Items
// @Summary trae un documento tipo item
// @Description Este método devuelve un solo documento tipo item según el id proporcionado para la búsqueda
// @Param id path string true "ID de"
// @Accept json
// @Produce json
// @Success 200 {object} models.ItemsModel "Documento tipo item"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetItemController(c echo.Context) error {
	item_id := c.Param("id") //Obtener parámetro {id} de la URI
	item_ObjectID, err := primitive.ObjectIDFromHex(item_id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameter")
	}
	var result models.ItemsModel
	err = result.GetSingleObject(item_ObjectID)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}
	return c.JSON(http.StatusOK, result)
}

// ItemsGet godoc
//
// @Router /items/ [get]
// @Tags Items
// @Summary trae una colección de documentos tipo Items
// @Description Este método devuelve una colección de documento tipo Items según la paginación definida
// @Param page_size query int true "Tamaño de la colección"
// @Param page_number query int true "páginación de la colección"
// @Accept json
// @Produce json
// @Success 200 {object} []models.ItemsModel "Documento tipo Items"
// @Failure 400 {object} string "Id de búsqueda inválido"
// @Failure 404 {object} string "Documento no existente en la base de datos"
func GetItemssController(c echo.Context) error {
	var query models.QueryPagination
	err := c.Bind(&query)
	if err != nil || !query.Validate() {
		return c.String(http.StatusBadRequest, "Bad Request: Invalid parameters")
	}
	var function models.ItemsModel
	var result []models.ItemsModel
	result, err = function.GetMultipleObjects(query)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Errors")
	}
	return c.JSON(http.StatusOK, result)
}
