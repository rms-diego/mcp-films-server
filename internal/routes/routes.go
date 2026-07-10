package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Init(r *gin.Engine, mcps *mcp.Server) {
	mcphandler := mcpHandler(mcps)

	mcproutes := r.Group("/mcp")
	mcproutes.GET("", func(c *gin.Context) {
		mcphandler.ServeHTTP(c.Writer, c.Request)
	})
	mcproutes.POST("", func(c *gin.Context) {
		mcphandler.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/heath-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running"})
	})
}

func mcpHandler(ms *mcp.Server) http.Handler {
	return mcp.NewStreamableHTTPHandler(
		func(req *http.Request) *mcp.Server {
			return ms
		},
		nil,
	)
}
