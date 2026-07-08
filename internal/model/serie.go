package model

type Serie struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float32  `json:"vote_average"`
	VoteCount        float32  `json:"vote_count"`
}

// {
//   "adult": false,
//   "backdrop_path": "/bsNm9z2TJfe0WO3RedPGWQ8mG1X.jpg",
//   "genre_ids": [
//     18,
//     80
//   ],
//   "id": 1396,
//   "origin_country": [
//     "US"
//   ],
//   "original_language": "en",
//   "original_name": "Breaking Bad",
//   "overview": "When Walter White, a New Mexico chemistry teacher, is diagnosed with Stage III cancer and given a prognosis of only two years left to live. He becomes filled with a sense of fearlessness and an unrelenting desire to secure his family's financial future at any cost as he enters the dangerous world of drugs and crime.",
//   "popularity": 298.884,
//   "poster_path": "/ggFHVNu6YYI5L9pCfOacjizRGt.jpg",
//   "first_air_date": "2008-01-20",
//   "name": "Breaking Bad",
//   "vote_average": 8.879,
//   "vote_count": 11536
// }
