package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
	moviemodule "github.com/rms-diego/mcp-films-server/internal/modules/movie"
	seriemodule "github.com/rms-diego/mcp-films-server/internal/modules/serie"
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

	tmdbg := tmdbgateway.NewTMDBGateway()
	moviemodule.Init(mcps, tmdbg)
	seriemodule.Init(mcps, tmdbg)
}

func mcpHandler(ms *mcp.Server) http.Handler {
	return mcp.NewStreamableHTTPHandler(
		func(req *http.Request) *mcp.Server {
			return ms
		},
		nil,
	)
}
