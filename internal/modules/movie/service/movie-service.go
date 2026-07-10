package movieservice

import (
	"context"

	tmdbgateway "github.com/rms-diego/mcp-films-server/internal/gateway/tmdb"
	"github.com/rms-diego/mcp-films-server/internal/model"
	moviedto "github.com/rms-diego/mcp-films-server/internal/modules/movie/dto"
)

type movieService struct {
	g tmdbgateway.ITMDBGateway
}

type IMovieService interface {
	FindMovieByName(ctx context.Context, i moviedto.FindMovieByNameInput) ([]model.Movie, error)
}

func NewMovieService(g tmdbgateway.ITMDBGateway) IMovieService {
	return &movieService{g}
}

func (s *movieService) FindMovieByName(ctx context.Context, input moviedto.FindMovieByNameInput) ([]model.Movie, error) {
	r, err := s.g.FindMovieByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return r.Results, nil
}
