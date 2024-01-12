package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marceljsh/grab2vec-api/config"
	"github.com/marceljsh/grab2vec-api/models"
)

func GetPointById(ctx *gin.Context) {
	pointId := ctx.Param("point_id")

	var point models.Point
	result := config.DB.Where("point_id = ?", pointId).Find(&point)
	if result.Error != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	if point.PointID == 0 {
		ctx.IndentedJSON(
			http.StatusNotFound,
			gin.H{"error": "point not found"},
		)
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"point": point},
	)
}

func GetPointsByTrjId(ctx *gin.Context) {
	trjId := ctx.Query("trj_id")
	if trjId == "" {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": "trj_id is not provided"},
		)
		return
	}

	var points []models.Point
	result := config.DB.Where("trj_id = ?", trjId).Find(&points)
	if result.Error != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"error": result.Error.Error()},
		)
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"points": points},
	)
}
