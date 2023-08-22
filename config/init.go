package initConfig

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	//Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("No se encontró el archivo .env") //Se detienen los procesos si no se encuentra el archivo .env
	}
}
