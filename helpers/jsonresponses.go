package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  false,
		"message": message,
	})
}

func JsonResponse(c *gin.Context, status bool, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  status,
		"message": message,
	})
}

func NotFoundResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data Tidak Ditemukan",
			})
}

func InternalServerResponse(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": err,
			})
}

func UnAuthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  false,
		"message": message,
	})
}