package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Init(r *gin.Engine, ms *mcp.Server) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running"})
	})

	h := mcp.NewStreamableHTTPHandler(
		func(req *http.Request) *mcp.Server {
			return ms
		},
		nil,
	)

	r.Any("/mcp", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}
