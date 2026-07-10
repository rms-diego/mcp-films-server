package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(s *gin.Engine) {
	s.GET("/heath-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running"})
	})
}
