# Golang Starter Project with Gin, MongoDB, and Redis

This repository is a starter template for building web applications using Go (Golang) with the Gin framework, MongoDB as the database, and Redis for caching. It provides a basic structure and essential configurations to help you get started quickly.

## Features

- **Gin Framework**: A fast HTTP web framework for Go.
- **MongoDB**: NoSQL database for storing application data.
- **Redis**: In-memory data structure store for caching.
- **Environment Configuration**: Simple setup using environment variables.
- **Basic REST API**: Example routes to demonstrate functionality.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (1.16 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) (local or cloud instance)
- [Redis](https://redis.io/download) (local or cloud instance)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/masDimas98/baseProject.git
cd golang-starter-gin-mongo-redis

go mod tidy

go run main.go
