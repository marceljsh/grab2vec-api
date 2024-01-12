package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/marceljsh/grab2vec-api/routes/v1"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1.SetupRoutes(router)

	return router
}