package controllers

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// AllGet godoc
//
// @Router /all/ [get]
// @Tags Root
// @Summary obtener todos los documentos en un solo objeto
// @Description Este método devuelve un json que contiene todos los documentos de todas las colecciones existentes
// @Accept json
// @Produce json
// @Success 200 {object} []interface{}
func GetAllController(c echo.Context) error {
	//Definición de contexto y call al cliente de mongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.GetMongoClient()
	//Listar todas las colecciones de la base de datos
	collections, err := db.Database("nexusBattlesII").ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	//Crear un objeto para almacenar los documentos por colección
	var documents []interface{}
	//Recorrer las colecciones y obtener todos los documentos
	for _, collectionName := range collections {
		collection := db.Database(os.Getenv("MONGO_DB")).Collection(collectionName)
		cur, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return err
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var document map[string]interface{}
			err = cur.Decode(&document)
			if err != nil {
				return err
			}
			documents = append(documents, document)
		}
		if err = cur.Err(); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, documents)
}
