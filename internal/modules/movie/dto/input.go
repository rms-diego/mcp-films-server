package moviedto

type FindMovieByNameInput struct {
	Name string `json:"name" jsonschema:"the name of the movie to make a search in TMDB API"`
}
