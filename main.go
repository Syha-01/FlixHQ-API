package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demonkingswarn/movie-api/handlers"
)

func main() {
	h := handlers.NewHandler()

	// Set up routes
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/health", h.HealthCheck)
	http.HandleFunc("/providers", h.ListProviders)
	http.HandleFunc("/search", h.Search)
	http.HandleFunc("/mediaId", h.GetMediaID)
	http.HandleFunc("/seasons", h.GetSeasons)
	http.HandleFunc("/episodes", h.GetEpisodes)
	http.HandleFunc("/servers", h.GetServers)
	http.HandleFunc("/link", h.GetLink)
	http.HandleFunc("/stream", h.GetStream)
	http.HandleFunc("/home", h.GetHome)

	// Get port from environment (Cloud Run requirement)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Movie API server starting on port %s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET /health              - Health check")
	fmt.Println("  GET /providers           - List available providers")
	fmt.Println("  GET /search?q=&provider= - Search movies/TV shows")
	fmt.Println("  GET /mediaId?url=&provider= - Get media ID from URL")
	fmt.Println("  GET /seasons?mediaId=&provider= - Get seasons")
	fmt.Println("  GET /episodes?id=&isSeason=&provider= - Get episodes")
	fmt.Println("  GET /servers?episodeId=&provider= - Get servers")
	fmt.Println("  GET /link?serverId=&provider= - Get embed link")
	fmt.Println("  GET /stream?serverId=&provider= - Get decrypted stream URL")
	fmt.Println("  GET /home?provider= - Get home page content (trending, latest)")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
