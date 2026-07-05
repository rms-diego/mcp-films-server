package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type TMDBConfig struct {
	APIKey          string
	ReadAccessToken string
}

type env struct {
	PORT       string
	TMDBConfig *TMDBConfig
}

var Cfg *env

func Init() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	switch {
	case os.Getenv("TMDB_API_KEY") == "":
		return fmt.Errorf("TMDB_API_KEY environment variable is not set")

	case os.Getenv("TMDB_READ_ACCESS_TOKEN") == "":
		return fmt.Errorf("TMDB_READ_ACCESS_TOKEN environment variable is not set")

	default:
		p := os.Getenv("PORT")
		if p == "" {
			p = "8080"
		}

		Cfg = &env{
			PORT: p,
			TMDBConfig: &TMDBConfig{
				APIKey:          os.Getenv("TMDB_API_KEY"),
				ReadAccessToken: os.Getenv("TMDB_READ_ACCESS_TOKEN"),
			},
		}

		return nil
	}
}
