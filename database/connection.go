package database

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongo_db     *mongo.Client   //Variable de acceso global para el cliente mongoDB
	syncInstance = &sync.Mutex{} //controlador para obtener una sola instancia del cliente mongoDB
)

// Obtener el cliente de mongoDB
func GetMongoClient() *mongo.Client {
	if mongo_db == nil {
		syncInstance.Lock()
		defer syncInstance.Unlock()
		//doble verificación de la instancia
		if mongo_db == nil {
			mongo_db = connectMongoClient()
		} else {
			return mongo_db
		}

	} else {
		return mongo_db
	}
	return mongo_db
}

// Realizar la conexión del cliente a mongoDB
func connectMongoClient() *mongo.Client {
	var monbo_client *mongo.Client
	monbo_client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_CONN")))
	if err != nil {
		log.Fatal(err)
	}
	return monbo_client
}
