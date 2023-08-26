package models

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
