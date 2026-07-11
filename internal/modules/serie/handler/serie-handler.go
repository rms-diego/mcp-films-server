package seriehandler

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	commondto "github.com/rms-diego/mcp-films-server/internal/common/dto"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	serieservice "github.com/rms-diego/mcp-films-server/internal/modules/serie/service"
)

type serieHandler struct {
	s serieservice.ISeriesService
}

type ISerieHandler interface {
	FindSeriesByName(ctx context.Context, req *mcp.CallToolRequest, input commondto.FindByNameInput) (
		*mcp.CallToolResult,
		*commondto.FindByNameOutput[model.Serie],
		error,
	)
}

func NewSerieHandler(s serieservice.ISeriesService) ISerieHandler {
	return &serieHandler{s}
}

func (h *serieHandler) FindSeriesByName(ctx context.Context, req *mcp.CallToolRequest, input commondto.FindByNameInput) (
	*mcp.CallToolResult,
	*commondto.FindByNameOutput[model.Serie],
	error,
) {

	r, err := h.s.FindSeriesByName(ctx, input)
	if err != nil {
		return nil, nil, err
	}

	return nil, &commondto.FindByNameOutput[model.Serie]{Data: r}, nil
}
