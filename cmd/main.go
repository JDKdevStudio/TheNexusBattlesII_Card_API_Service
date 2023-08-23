package main

import (
	initConfig "TheNexusBattlesII_Card_API_Service/config"
	"TheNexusBattlesII_Card_API_Service/database"
	"context"
	"fmt"
)

func main() {
	//Inicializar congifuración
	initConfig.Init()

	//Get db
	var pingdb = database.GetMongoClient()

	if err := pingdb.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to MongoDB")
}
