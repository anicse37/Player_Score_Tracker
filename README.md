# 🏏 Player Score Tracker (Go Web Server)

This is a basic HTTP server written in Go that returns a hardcoded score for a player. It's designed as a starting point to learn how to build and test web servers in Go.



## 🚀 How It Works

The server responds to HTTP GET requests. Currently, it returns a static score for any request made to `/players/<playerName>`.

### Example Request

```http
GET /players/Player-1 HTTP/1.1


