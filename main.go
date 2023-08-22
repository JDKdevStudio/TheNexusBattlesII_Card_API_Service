package main

import (
	initConfig "TheNexusBattlesII_Card_API_Service/config"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Init Config.
	initConfig.InitConfig()

	// Create a connection string.
	connectionString := os.Getenv("MONGO_CONN")

	// Create a MongoClient.
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the connection.
	if err := client.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to MongoDB")

}
