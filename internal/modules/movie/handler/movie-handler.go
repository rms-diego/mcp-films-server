package moviehandler

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	commondto "github.com/rms-diego/mcp-films-server/internal/common/dto"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	movieservice "github.com/rms-diego/mcp-films-server/internal/modules/movie/service"
)

type movieHandler struct {
	s movieservice.IMovieService
}

type IMovieHandler interface {
	FindMoviesByName(ctx context.Context, req *mcp.CallToolRequest, input commondto.FindByNameInput) (
		*mcp.CallToolResult,
		*commondto.FindByNameOutput[model.Movie],
		error,
	)
}

func NewMovieHandler(s movieservice.IMovieService) IMovieHandler {
	return &movieHandler{s}
}

func (h *movieHandler) FindMoviesByName(ctx context.Context, req *mcp.CallToolRequest, input commondto.FindByNameInput) (
	*mcp.CallToolResult,
	*commondto.FindByNameOutput[model.Movie],
	error,
) {
	r, err := h.s.FindMovieByName(ctx, input)
	if err != nil {
		return nil, nil, err
	}

	return nil, &commondto.FindByNameOutput[model.Movie]{Data: r}, nil
}
