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
// @Success 200 {object} models.HabilidadesModel "Documento tipo Héroe"
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
