package seriemodule

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
	seriehandler "github.com/rms-diego/mcp-films-server/internal/modules/serie/handler"
	serieservice "github.com/rms-diego/mcp-films-server/internal/modules/serie/service"
)

func Init(mcps *mcp.Server) {
	g := tmdbgateway.NewTMDBGateway()
	s := serieservice.NewSeriesService(g)
	h := seriehandler.NewSerieHandler(s)

	mcp.AddTool(
		mcps,
		&mcp.Tool{
			Name:        "Find series by name",
			Description: "A tool to find series searching in TMDB api",
		},
		h.FindSeriesByName,
	)
}
