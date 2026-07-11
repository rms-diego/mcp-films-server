package commondto

type FindByNameInput struct {
	Name string `json:"name" jsonschema:"Name of the movie or tv show to make a search in TMDB API"`
}
