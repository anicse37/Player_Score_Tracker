# Player Score Tracker

A simple RESTful API written in Go that tracks player scores, records wins, and displays a league leaderboard. It includes an in-memory store, file-based store, and full test coverage using the Go testing package.

## Features

- Record wins for players via HTTP POST
- Retrieve player scores via HTTP GET
- Return the current league table as JSON
- In-memory and file-backed data stores
- Modular project structure
- Dockerized for easy deployment
- Fully tested using `net/http/httptest`

## Project Structure

<details>
  <summary><strong>📁 Application</strong></summary>

  <details>
    <summary>📂 main</summary>

    - main.go
  </details>

  <details>
    <summary>📂 Models</summary>

    - players.go
  </details>

  <details>
    <summary>📂 Servers</summary>

    - server.go  
    - server_test.go  
    - server_integration_test.go
  </details>

  - functions.go  
  - Dockerfile  
  - docker-compose.yml  
  - go.mod
</details>


markdown
Copy code

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18+
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized deployment)

### Running Locally

```bash
go run main/main.go
```

### Running Tests

```bash
go test ./...
```

### Using Docker
Build and run the container using:


```bash
docker build -t player-score-tracker .
docker run -p 8080:8080 player-score-tracker
```

### Or use Docker Compose:

```bash
docker-compose up
```

### API Endpoints
GET /players/{name} - Get score for a specific player

POST /players/{name} - Record a win for a player

GET /league - Get league table as JSON
