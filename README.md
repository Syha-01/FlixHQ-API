# Movie API

A REST API for scraping movie and TV show data from various providers. Extracted from the [Luffy CLI](https://github.com/DemonKingSwarn/luffy).

## Features

- **Multiple Providers**: Support for FlixHQ, Brocoflix, XPrime, Sflix, Braflix, HDRezka, Movies4u, and YouTube.
- **Unified Interface**: Common response format across all providers.
- **Stream Extraction**: Automatically decrypts and extracts stream URLs (m3u8) from embed links.
- **Cloud Run Ready**: Dockerized and configured for Google Cloud Run.

## Installation & Usage

### Local Development

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd movie-API
    ```

2.  **Run the server:**
    ```bash
    go run main.go
    ```
    The server will start on port 8080 (or the port specified by `$PORT`).

### Docker

1.  **Build the image:**
    ```bash
    docker build -t movie-api .
    ```

2.  **Run the container:**
    ```bash
    docker run -p 8080:8080 movie-api
    ```

## API Endpoints

### General

-   `GET /health`: Health check. Returns `{"status": "ok"}`.
-   `GET /providers`: List available provider names.

### Search & Discovery

-   `GET /search`: Search for movies or TV shows.
    -   Query Params:
        -   `q`: Search query (equired).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON array of search results.

### Media Information

-   `GET /seasons`: Get seasons for a specific series.
    -   Query Params:
        -   `mediaId`: Media ID returned from search (required).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON array of seasons.

-   `GET /episodes`: Get episodes for a season (or movie servers).
    -   Query Params:
        -   `id`: Season ID or Media ID (required).
        -   `isSeason`: `true` if fetching episodes for a season, `false` for movie (required).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON array of episodes or servers.

-   `GET /servers`: Get available servers for an episode.
    -   Query Params:
        -   `episodeId`: Episode ID (required).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON array of servers.

### Streaming

-   `GET /link`: Get the raw embed link.
    -   Query Params:
        -   `serverId`: Server ID (required).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON object `{"link": "..."}`.

-   `GET /stream`: Get the decrypted stream URL (m3u8).
    -   Query Params:
        -   `serverId`: Server ID (required).
        -   `provider`: Provider name (default: `flixhq`).
    -   Response: JSON object `{"stream_url": "...", "subtitles": [...], "referer": "..."}`.

## Deployment to Google Cloud Run

```bash
gcloud builds submit --tag gcr.io/PROJECT_ID/movie-api
gcloud run deploy movie-api --image gcr.io/PROJECT_ID/movie-api --platform managed --allow-unauthenticated
```
