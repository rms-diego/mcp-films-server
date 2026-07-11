package model

type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// {
//   "adult": false,
//   "backdrop_path": "/hZkgoQYus5vegHoetLkCJzb17zJ.jpg",
//   "genre_ids": [
//     18,
//     53,
//     35
//   ],
//   "id": 550,
//   "original_language": "en",
//   "original_title": "Fight Club",
//   "overview": "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion.",
//   "popularity": 73.433,
//   "poster_path": "/pB8BM7pdSp6B6Ih7QZ4DrQ3PmJK.jpg",
//   "release_date": "1999-10-15",
//   "title": "Fight Club",
//   "video": false,
//   "vote_average": 8.433,
//   "vote_count": 26279
// }
