package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marceljsh/grab2vec-api/config"
	"github.com/marceljsh/grab2vec-api/routes"
	log "github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("APP_PORT")

	app := routes.Init()

	log.Fatal(app.Run(":" + port))
}
