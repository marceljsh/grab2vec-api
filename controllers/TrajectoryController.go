package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marceljsh/grab2vec-api/config"
	"github.com/marceljsh/grab2vec-api/models"
	"gorm.io/gorm"
)

func GetTrajectoryById(ctx *gin.Context) {
	trjId := ctx.Param("trj_id")

	var trajectory models.Trajectory
	var result *gorm.DB

	preload := ctx.Query("preload")
	if preload == "true" {
		result = config.DB.Preload("Points").Where("trj_id = ?", trjId).Find(&trajectory)
	} else {
		result = config.DB.Where("trj_id = ?", trjId).Find(&trajectory)
	}

	if result.Error != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	if trajectory.TrjID == 0 {
		ctx.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "trajectory not found"},
		)
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"trajectory": trajectory},
	)
}
