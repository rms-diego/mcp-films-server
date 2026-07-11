package movieservice

import (
	"context"

	commondto "github.com/rms-diego/mcp-films-server/internal/common/dto"
	"github.com/rms-diego/mcp-films-server/internal/common/model"
	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
)

type movieService struct {
	g tmdbgateway.ITMDBGateway
}

type IMovieService interface {
	FindMoviesByName(ctx context.Context, input commondto.FindByNameInput) ([]model.Movie, error)
}

func NewMovieService(g tmdbgateway.ITMDBGateway) IMovieService {
	return &movieService{g}
}

func (s *movieService) FindMoviesByName(ctx context.Context, input commondto.FindByNameInput) ([]model.Movie, error) {
	r, err := s.g.FindMoviesByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return r.Results, nil
}
