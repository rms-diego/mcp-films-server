package tmdbgateway

import (
	"context"
	"fmt"

	"github.com/rms-diego/mcp-films-server/internal/common/model"
	"github.com/rms-diego/mcp-films-server/internal/config"
	"github.com/rms-diego/mcp-films-server/internal/consts"
	"github.com/rms-diego/mcp-films-server/internal/utils"
)

type tMDBGateway struct {
	headers map[string]string
}

type ITMDBGateway interface {
	FindMoviesByName(ctx context.Context, name string) (*model.SearchResult[model.Movie], error)
	FindSeriesByName(ctx context.Context, name string) (*model.SearchResult[model.Serie], error)
}

func NewTMDBGateway() ITMDBGateway {
	return &tMDBGateway{
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", config.Env.TMDBConfig.ReadAccessToken),
		},
	}
}

func (s *tMDBGateway) FindMoviesByName(ctx context.Context, name string) (*model.SearchResult[model.Movie], error) {
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

func (s *tMDBGateway) FindSeriesByName(ctx context.Context, name string) (*model.SearchResult[model.Serie], error) {
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
