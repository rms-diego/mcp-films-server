package tmdbgateway

import (
	"context"
	"fmt"

	"github.com/rms-diego/mcp-films-server/internal/config"
	"github.com/rms-diego/mcp-films-server/internal/consts"
	"github.com/rms-diego/mcp-films-server/internal/model"
	"github.com/rms-diego/mcp-films-server/internal/utils"
)

type ITMDBGateway interface {
	FindMovieByName(ctx context.Context, name string) (*model.SearchResult[model.Movie], error)
	FindSerieByName(ctx context.Context, name string) (*model.SearchResult[model.Serie], error)
}

type tMDBGateway struct {
	headers map[string]string
}

func NewTMDBGateway() ITMDBGateway {
	return &tMDBGateway{
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", config.Env.TMDBConfig.ReadAccessToken),
		},
	}
}

func (s tMDBGateway) FindMovieByName(ctx context.Context, name string) (*model.SearchResult[model.Movie], error) {
	p := utils.Payload{
		Url:     consts.TMDB_API_URL + "/search/movie?query=" + name,
		Headers: s.headers,
		Method:  utils.GET,
	}

	r, err := utils.Fetch[model.SearchResult[model.Movie]](ctx, p)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s tMDBGateway) FindSerieByName(ctx context.Context, name string) (*model.SearchResult[model.Serie], error) {
	p := utils.Payload{
		Url:     consts.TMDB_API_URL + "/search/tv?query=" + name,
		Headers: s.headers,
		Method:  utils.GET,
	}

	r, err := utils.Fetch[model.SearchResult[model.Serie]](ctx, p)
	if err != nil {
		return nil, err
	}

	return r, nil
}
