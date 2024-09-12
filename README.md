# Trocup User Service

## Description

This is a user management microservice built with Go, Fiber, and MongoDB.

## Setup

1. Install Go
2. Clone the repository
3. Install the dependencies: `go mod tidy`
4. Set up MongoDB Atlas and update the URI in `config/config.go` - ℹ️ Refer to ticket `BACK-3`
5. Refer to `.env.example`file to create your `.env`
6. Run the service's server: `go run main.go`
