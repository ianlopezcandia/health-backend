package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar .env: %v", err)
	}

	//add mongo

	gin := gin.Default()

	if err := gin.Run(os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Error al iniciar %v",err)
	}

}
