package moviemodule

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
	moviehandler "github.com/rms-diego/mcp-films-server/internal/modules/movie/handler"
	movieservice "github.com/rms-diego/mcp-films-server/internal/modules/movie/service"
)

func Init(mcps *mcp.Server, tmdbg tmdbgateway.ITMDBGateway[model.Movie]) {
	s := movieservice.NewMovieService(tmdbg)
	h := moviehandler.NewMovieHandler(s)

	mcp.AddTool(
		mcps,
		&mcp.Tool{
			Name:        "Find films by name",
			Description: "A tool to find movies searching in TMDB api",
		},
		h.FindMoviesByName,
	)
}
