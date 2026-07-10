package moviehandler

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	moviedto "github.com/rms-diego/mcp-films-server/internal/modules/movie/dto"
	movieservice "github.com/rms-diego/mcp-films-server/internal/modules/movie/service"
)

type movieHandler struct {
	s movieservice.IMovieService
}

type IMovieHandler interface {
	FindMovieByName(ctx context.Context, req *mcp.CallToolRequest, input moviedto.FindMovieByNameInput) (
		*mcp.CallToolResult,
		*moviedto.FindMovieByNameOutput,
		error,
	)
}

func NewMovieHandler(s movieservice.IMovieService) IMovieHandler {
	return &movieHandler{s: s}
}

func (h *movieHandler) FindMovieByName(ctx context.Context, req *mcp.CallToolRequest, input moviedto.FindMovieByNameInput) (
	*mcp.CallToolResult,
	*moviedto.FindMovieByNameOutput,
	error,
) {
	result, err := h.s.FindMovieByName(ctx, input)
	if err != nil {
		return nil, nil, err
	}

	return nil, &moviedto.FindMovieByNameOutput{Films: result}, nil
}
