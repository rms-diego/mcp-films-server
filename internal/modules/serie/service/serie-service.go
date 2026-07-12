package serieservice

import (
	"context"

	commondto "github.com/rms-diego/mcp-films-server/internal/common/dto"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
)

type seriesService struct {
	g tmdbgateway.ITMDBGateway[model.Serie]
}

type ISeriesService interface {
	FindSeriesByName(ctx context.Context, input commondto.FindByNameInput) ([]model.Serie, error)
}

func NewSeriesService(g tmdbgateway.ITMDBGateway[model.Serie]) ISeriesService {
	return &seriesService{g}
}

func (s *seriesService) FindSeriesByName(ctx context.Context, input commondto.FindByNameInput) ([]model.Serie, error) {
	r, err := s.g.FindManyByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return r.Results, nil
}
