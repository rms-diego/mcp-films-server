package moviedto

import "github.com/rms-diego/mcp-films-server/internal/model"

type FindMovieByNameOutput struct {
	Films []model.Movie
}
