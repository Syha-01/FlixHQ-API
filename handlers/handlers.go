package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/demonkingswarn/movie-api/core"
	"github.com/demonkingswarn/movie-api/providers"
)

type Handler struct {
	providers map[string]core.Provider
	client    *http.Client
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProvidersResponse struct {
	Providers []string `json:"providers"`
}

func NewHandler() *Handler {
	client := core.NewClient()
	h := &Handler{
		providers: make(map[string]core.Provider),
		client:    client,
	}

	// Register all providers
	h.providers["flixhq"] = providers.NewFlixHQ(client)
	h.providers["brocoflix"] = providers.NewBrocoflix(client)
	h.providers["xprime"] = providers.NewXPrime(client)
	h.providers["sflix"] = providers.NewSflix(client)
	h.providers["braflix"] = providers.NewBraflix(client)
	h.providers["hdrezka"] = providers.NewHDRezka(client)
	h.providers["movies4u"] = providers.NewMovies4u(client)
	h.providers["youtube"] = providers.NewYouTube(client)

	return h
}

func (h *Handler) getProvider(name string) core.Provider {
	if name == "" {
		name = "flixhq" // default provider
	}
	return h.providers[strings.ToLower(name)]
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

// Health check endpoint for Cloud Run
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// List available providers
func (h *Handler) ListProviders(w http.ResponseWriter, r *http.Request) {
	var names []string
	for name := range h.providers {
		names = append(names, name)
	}
	writeJSON(w, http.StatusOK, ProvidersResponse{Providers: names})
}

// Search for movies/TV shows
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	providerName := r.URL.Query().Get("provider")

	if query == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'q' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	results, err := provider.Search(query)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, results)
}

// Get media ID from URL
func (h *Handler) GetMediaID(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	providerName := r.URL.Query().Get("provider")

	if url == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'url' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	mediaID, err := provider.GetMediaID(url)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"media_id": mediaID})
}

// Get seasons for a TV show
func (h *Handler) GetSeasons(w http.ResponseWriter, r *http.Request) {
	mediaID := r.URL.Query().Get("mediaId")
	providerName := r.URL.Query().Get("provider")

	if mediaID == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'mediaId' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	seasons, err := provider.GetSeasons(mediaID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, seasons)
}

// Get episodes for a season
func (h *Handler) GetEpisodes(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	providerName := r.URL.Query().Get("provider")
	isSeason := r.URL.Query().Get("isSeason") == "true"

	if id == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'id' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	episodes, err := provider.GetEpisodes(id, isSeason)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, episodes)
}

// Get available servers for an episode
func (h *Handler) GetServers(w http.ResponseWriter, r *http.Request) {
	episodeID := r.URL.Query().Get("episodeId")
	providerName := r.URL.Query().Get("provider")

	if episodeID == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'episodeId' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	servers, err := provider.GetServers(episodeID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, servers)
}

// Get embed link for a server
func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("serverId")
	providerName := r.URL.Query().Get("provider")

	if serverID == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'serverId' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	link, err := provider.GetLink(serverID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"link": link})
}

// Get decrypted stream URL
func (h *Handler) GetStream(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("serverId")
	providerName := r.URL.Query().Get("provider")

	if serverID == "" {
		writeError(w, http.StatusBadRequest, "query parameter 'serverId' is required")
		return
	}

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	// First get the embed link
	link, err := provider.GetLink(serverID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Then decrypt to get stream URL
	streamURL, subtitles, referer, err := core.DecryptStream(link, h.client)
	if err != nil {
		// If decryption fails, return the embed link
		writeJSON(w, http.StatusOK, core.StreamResult{
			StreamURL: link,
			Referer:   "",
		})
		return
	}

	writeJSON(w, http.StatusOK, core.StreamResult{
		StreamURL: streamURL,
		Subtitles: subtitles,
		Referer:   referer,
	})
}

// Get home page content (trending, latest, coming soon)
func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	providerName := r.URL.Query().Get("provider")

	provider := h.getProvider(providerName)
	if provider == nil {
		writeError(w, http.StatusBadRequest, "invalid provider")
		return
	}

	result, err := provider.GetHome()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}
