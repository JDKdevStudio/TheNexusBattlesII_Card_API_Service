package models

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"os"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HeroesModel struct {
	Id         primitive.ObjectID `form:"id"         bson:"_id,omitempty"`
	UrlImagen  *string            `form:"urlImagen"  bson:"urlImagen,omitempty"  validate:"required,omitempty"`
	Clase      *string            `form:"clase"      bson:"clase,omitempty"      validate:"required,omitempty"`
	Tipo       *string            `form:"tipo"       bson:"tipo,omitempty"       validate:"required,omitempty"`
	Poder      *int               `form:"poder"      bson:"poder,omitempty"      validate:"required,omitempty"`
	Vida       *int               `form:"vida"       bson:"vida,omitempty"       validate:"required,omitempty"`
	Defensa    *int               `form:"defensa"    bson:"defensa,omitempty"    validate:"required,omitempty"`
	AtaqueBase *int               `form:"ataqueBase" bson:"ataqueBase,omitempty" validate:"required,omitempty"`
	AtaqueDado *int               `form:"ataqueDado" bson:"ataqueDado,omitempty" validate:"required,omitempty"`
	DanoMax    *int               `form:"danoMax"    bson:"danoMax,omitempty"    validate:"required,omitempty"`
	Activo     *bool              `form:"activo"     bson:"activo,omitempty"`
	Desc       *string            `form:"desc"       bson:"desc,omitempty"       validate:"required,omitempty"`
}

func (h *HeroesModel) Validate() bool {
	validate := validator.New()
	if err := validate.Struct(h); err != nil {
		return false
	}

	return true
}

func (h *HeroesModel) GetsingeObject(heroe_ObjectID primitive.ObjectID) error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("heroes")
	return collection.FindOne(context.Background(), bson.M{"_id": heroe_ObjectID}).Decode(&h)
}

func (HeroesModel) GetMultipleObjects(q QueryPagination) ([]HeroesModel, error) {
	result := make([]HeroesModel, 0)
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("heroes")
	cursor, err := collection.Find(context.Background(), bson.D{}, options.Find().SetSkip(int64((q.Page_number-1)*q.Page_size)).SetLimit(int64(q.Page_size)))
	if err != nil {
		return result, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var document HeroesModel
		if err = cursor.Decode(&document); err != nil {
			return result, err
		}
		result = append(result, document)
	}
	return result, err
}

func (h HeroesModel) PostSingleObject() error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("heroes")
	_, err := collection.InsertOne(context.Background(), bson.M{"urlImagen": h.UrlImagen, "clase": h.Clase, "tipo": h.Tipo, "poder": h.Poder, "vida": h.Vida, "defensa": h.Defensa, "ataqueBase": h.AtaqueBase, "ataqueDado": h.AtaqueDado, "danoMax": h.DanoMax, "activo": h.Activo, "desc": h.Desc})
	return err
}

func (h HeroesModel) UpdateSingleObject(heroe_ObjectID primitive.ObjectID) error {
	db := database.GetMongoClient()
	collection := db.Database(os.Getenv("MONGO_DB")).Collection("heroes")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": heroe_ObjectID}, bson.M{"$set": h})
	return err
}
