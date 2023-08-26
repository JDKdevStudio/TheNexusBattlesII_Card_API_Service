package config

import (
	"TheNexusBattlesII_Card_API_Service/database"
	"TheNexusBattlesII_Card_API_Service/utils"
	"context"
	"fmt"
)

func Init() {
	//Load .env file
	utils.LoadEnv()

	var pingdb = database.GetMongoClient()

	if err := pingdb.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to MongoDB")
}
