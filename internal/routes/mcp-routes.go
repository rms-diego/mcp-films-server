package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	moviemodule "github.com/rms-diego/mcp-films-server/internal/modules/movie"
)

func InitMCPRoutes(s *gin.Engine, mcps *mcp.Server) {
	h := mcpHandler(mcps)
	g := s.Group("/mcp")

	g.GET("", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
	g.POST("", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	g.DELETE("", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	moviemodule.Init(mcps)
}

func mcpHandler(ms *mcp.Server) http.Handler {
	return mcp.NewStreamableHTTPHandler(
		func(req *http.Request) *mcp.Server {
			return ms
		},
		nil,
	)
}
