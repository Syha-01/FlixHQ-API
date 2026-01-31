package core

type Provider interface {
	Name() string
	Search(query string) ([]SearchResult, error)
	GetMediaID(url string) (string, error)
	GetSeasons(mediaID string) ([]Season, error)
	GetEpisodes(id string, isSeason bool) ([]Episode, error)
	GetServers(episodeID string) ([]Server, error)
	GetLink(serverID string) (string, error)
}
