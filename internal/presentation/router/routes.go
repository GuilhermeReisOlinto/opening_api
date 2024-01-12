package router

import (
	"opening_api/internal/domain/handler"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpening)
		v1.POST("/opening", handler.CreateOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.PATCH("/opening", handler.UpdateOpening)
		v1.GET("/openings", handler.ListOpening)
	}
}
