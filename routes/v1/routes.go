package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marceljsh/grab2vec-api/controllers"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ping received"})
		})

		v1_points := v1.Group("/points")
		{
			v1_points.GET("/", controllers.GetPointsByTrjId)
			v1_points.GET("/:point_id", controllers.GetPointById)
		}

		v1_trajectories := v1.Group("/trajectories")
		{
			v1_trajectories.GET("/:trj_id", controllers.GetTrajectoryById)
		}
	}
}
