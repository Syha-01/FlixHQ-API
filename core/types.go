package core

type Action string
type MediaType string

const (
	ActionPlay     Action = "play"
	ActionDownload Action = "download"

	Movie  MediaType = "movie"
	Series MediaType = "series"
)

const (
	FLIXHQ_BASE_URL   = "https://flixhq.to"
	FLIXHQ_SEARCH_URL = FLIXHQ_BASE_URL + "/search"
	FLIXHQ_AJAX_URL   = FLIXHQ_BASE_URL + "/ajax"
	DECODER           = "https://dec.eatmynerds.live"
	TMDB_API_KEY      = "653bb8af90162bd98fc7ee32bcbbfb3d"
	TMDB_BASE_URL     = "https://api.themoviedb.org/3"
)

type TmdbSearchResult struct {
	Results []struct {
		ID           int    `json:"id"`
		MediaType    string `json:"media_type"`
		Title        string `json:"title"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		ReleaseDate  string `json:"release_date"`
		FirstAirDate string `json:"first_air_date"`
	} `json:"results"`
}

type TmdbSeason struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SeasonNumber int    `json:"season_number"`
}

type TmdbShowDetails struct {
	Seasons []TmdbSeason `json:"seasons"`
}

type TmdbEpisode struct {
	ID            int    `json:"id"`
	EpisodeNumber int    `json:"episode_number"`
	Name          string `json:"name"`
}

type TmdbSeasonDetails struct {
	Episodes []TmdbEpisode `json:"episodes"`
}

type SearchResult struct {
	Title  string    `json:"title"`
	URL    string    `json:"url"`
	Type   MediaType `json:"type"`
	Poster string    `json:"poster"`
	Year   string    `json:"year,omitempty"`
}

type Season struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Episode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StreamResult contains the decrypted stream information
type StreamResult struct {
	StreamURL string   `json:"stream_url"`
	Subtitles []string `json:"subtitles,omitempty"`
	Referer   string   `json:"referer,omitempty"`
}

// HomeItem represents a single item on the home page
type HomeItem struct {
	Title  string    `json:"title"`
	URL    string    `json:"url"`
	Type   MediaType `json:"type"`
	Poster string    `json:"poster"`
}

// HomeResult contains all sections from the home page
type HomeResult struct {
	Trending     []HomeItem `json:"trending"`
	LatestMovies []HomeItem `json:"latest_movies"`
	LatestShows  []HomeItem `json:"latest_tv_shows"`
	ComingSoon   []HomeItem `json:"coming_soon"`
}
