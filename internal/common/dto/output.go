package commondto

type FindByNameOutput[T any] struct {
	Data []T `json:"data" jsonschema:"list of movies or tv shows found in TMDB API"`
}
