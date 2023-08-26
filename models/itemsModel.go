package models

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemsModel struct {
	Id        primitive.ObjectID     `bson:"_id,omitempty"`
	UrlImagen string                 `bson:"urlImagen,omitempty"`
	Heroe     string                 `bson:"heroe,omitempty"`
	Nombre    string                 `bson:"nombre,omitempty"`
	Efecto    map[string]interface{} `bson:"efecto,omitempty"`
	Activo    bool                   `bson:"activo,omitempty"`
	Desc      string                 `bson:"desc,omitempty"`
}

func (i *ItemsModel) GetSingleObject(item_ObjectID primitive.ObjectID) error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("items")
	return collection.FindOne(context.Background(), bson.M{"_id": item_ObjectID}).Decode(&i)
}

func (ItemsModel) GetMultipleObjects(q QueryPagination) ([]ItemsModel, error) {
	result := make([]ItemsModel, 0)
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("items")
	cursor, err := collection.Find(context.Background(), bson.D{}, options.Find().SetSkip(int64((q.Page_number-1)*q.Page_size)).SetLimit(int64(q.Page_size)))
	if err != nil {
		return result, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var document ItemsModel
		if err = cursor.Decode(&document); err != nil {
			return result, err
		}
		result = append(result, document)
	}
	return result, err
}
