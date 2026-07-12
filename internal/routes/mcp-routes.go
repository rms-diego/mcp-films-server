package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
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

	moviemodule.Init(mcps, tmdbgateway.NewTMDBGateway[model.Movie]())
	seriemodule.Init(mcps, tmdbgateway.NewTMDBGateway[model.Serie]())
}

func mcpHandler(ms *mcp.Server) http.Handler {
	return mcp.NewStreamableHTTPHandler(
		func(req *http.Request) *mcp.Server {
			return ms
		},
		nil,
	)
}
