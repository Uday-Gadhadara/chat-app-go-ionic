package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
