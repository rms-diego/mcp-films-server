package tmdbgateway

import (
	"context"
	"fmt"

	"github.com/rms-diego/mcp-films-server/internal/common/model"
	"github.com/rms-diego/mcp-films-server/internal/config"
	"github.com/rms-diego/mcp-films-server/internal/consts"
	"github.com/rms-diego/mcp-films-server/internal/utils"
)

type entertainment string

const (
	movie entertainment = "movie"
	tv    entertainment = "tv"
)

type tMDBGateway[T model.Movie | model.Serie] struct {
	headers     map[string]string
	gatewayType entertainment
}

type ITMDBGateway[T model.Movie | model.Serie] interface {
	FindManyByName(ctx context.Context, name string) (*model.SearchResult[T], error)
}

func NewTMDBGateway[T model.Movie | model.Serie]() ITMDBGateway[T] {
	var selType T
	var gatewayType entertainment

	switch any(selType).(type) {
	case model.Movie:
		gatewayType = movie
	case model.Serie:
		gatewayType = tv
	}

	return &tMDBGateway[T]{
		headers:     map[string]string{"Authorization": fmt.Sprintf("Bearer %v", config.Env.TMDBConfig.ReadAccessToken)},
		gatewayType: gatewayType,
	}
}

func (s *tMDBGateway[T]) FindManyByName(ctx context.Context, name string) (*model.SearchResult[T], error) {
	p := utils.Payload{
		Url:     fmt.Sprintf("%v/search/%v?query=%v", consts.TMDB_API_URL, string(s.gatewayType), name),
		Headers: s.headers,
		Method:  utils.GET,
	}

	r, err := utils.Fetch[model.SearchResult[T]](ctx, p)
	if err != nil {
		return nil, err
	}

	return r, nil
}
