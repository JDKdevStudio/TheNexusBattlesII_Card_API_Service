package models

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArmasModel struct {
	Id        primitive.ObjectID     `bson:"_id,omitempty"`
	UrlImagen string                 `bson:"urlImagen,omitempty"`
	Nombre    string                 `bson:"nombre,omitempty"`
	TipoHeroe string                 `bson:"tipoHeroe,omitempty"`
	Efecto    map[string]interface{} `bson:"efecto,omitempty"`
	Activo    bool                   `bson:"activo,omitempty"`
	Desc      string                 `bson:"desc,omitempty"`
}

func (a *ArmasModel) GetSingleObject(arma_ObjectID primitive.ObjectID) error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("armas")
	return collection.FindOne(context.Background(), bson.M{"_id": arma_ObjectID}).Decode(&a)
}

func (ArmasModel) GetMultipleObjects(q QueryPagination) ([]ArmasModel, error) {
	result := make([]ArmasModel, 0)
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("armas")
	cursor, err := collection.Find(context.Background(), bson.D{}, options.Find().SetSkip(int64((q.Page_number-1)*q.Page_size)).SetLimit(int64(q.Page_size)))
	if err != nil {
		return result, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var document ArmasModel
		if err = cursor.Decode(&document); err != nil {
			return result, err
		}
		result = append(result, document)
	}
	return result, err
}
