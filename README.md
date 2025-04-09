![Go](https://img.shields.io/github/languages/top/chapa-ai/click-counter)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/chapa-ai/click-counter)
![GitHub Repository Size](https://img.shields.io/github/repo-size/chapa-ai/click-counter)
![Github Open Issues](https://img.shields.io/github/issues/chapa-ai/click-counter)
![GitHub last commit](https://img.shields.io/github/last-commit/chapa-ai/click-counter)

# Click Counter API

API server for counting clicks on banners and collecting click statistics.

## Solution Notes

- :book: Standard Go project layout
- :cd: Dockerized with `docker-compose`
- :cloud: PostgreSQL for storing clicks statistics

## HOWTO

- :hammer_and_wrench: Clean up dependencies with `make tidy`
- :running_man: Run app in Docker Compose with `make run`
- :elephant: Launch test  with `go test ./test/click_test.go`
- :computer: API Endpoints:
    - `/counter/<bannerID>` (GET): Increments the click count for the given banner ID.
    - `/stats/<bannerID>` (POST): Retrieves click statistics for a given banner ID in a specified time range (`tsFrom`, `tsTo`).

## TODO
- Unit tests for the API endpoints

## Example

1. **Increment Click (GET /counter/<bannerID>)**
    - Example request:
   ```bash
   curl -X GET http://localhost:9999/counter/1
