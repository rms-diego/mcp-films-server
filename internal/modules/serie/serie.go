package seriemodule

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
	seriehandler "github.com/rms-diego/mcp-films-server/internal/modules/serie/handler"
	serieservice "github.com/rms-diego/mcp-films-server/internal/modules/serie/service"
)

func Init(mcps *mcp.Server, tmdbg tmdbgateway.ITMDBGateway[model.Serie]) {
	s := serieservice.NewSeriesService(tmdbg)
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
