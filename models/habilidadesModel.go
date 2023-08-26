package models

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HabilidadesModel struct {
	Id           primitive.ObjectID     `form:"id" bson:"_id,omitempty"`
	UrlImagen    string                 `form:"urlImagen" bson:"urlImagen,omitempty" validate:"required,omitempty"`
	Heroe        string                 `bson:"heroe,omitempty"`
	Nombre       string                 `bson:"nombre,omitempty"`
	EfectoGlobal map[string]interface{} `bson:"efectoGlobal,omitempty"`
	EfectoHeroe  map[string]interface{} `bson:"efectoHeroe,omitempty"`
	Activo       bool                   `bson:"activo,omitempty"`
	Desc         string                 `bson:"desc,omitempty"`
}

func (h *HabilidadesModel) GetSingleObject(habilidad_ObjectID primitive.ObjectID) error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("habilidades")
	return collection.FindOne(context.Background(), bson.M{"_id": habilidad_ObjectID}).Decode(&h)
}

func (HabilidadesModel) GetMultipleObjects(q QueryPagination) ([]HabilidadesModel, error) {
	result := make([]HabilidadesModel, 0)
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("habilidades")
	cursor, err := collection.Find(context.Background(), bson.D{}, options.Find().SetSkip(int64((q.Page_number-1)*q.Page_size)).SetLimit(int64(q.Page_size)))
	if err != nil {
		return result, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var document HabilidadesModel
		if err = cursor.Decode(&document); err != nil {
			return result, err
		}
		result = append(result, document)
	}
	return result, err
}
