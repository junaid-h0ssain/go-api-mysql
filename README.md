# Go MySQL API

This project is a REST API for managing books in a MySQL database. It provides endpoints for creating, reading, updating, and deleting book records.

## Stack

- **Language**: Go
- **Framework**: Gorilla Mux for routing
- **ORM**: GORM for database interactions
- **Database**: MySQL (using MariaDB image in Docker)
- **Containerization**: Docker Compose for database setup

## Prerequisites

- Go 1.26 or later
- Docker and Docker Compose

## Running the Application

1. Start the MySQL database using Docker Compose:
   ```
   docker-compose up -d
   ```
   This will start a MariaDB container with the database configured.

2. Run the Go application:
   ```
   go run main.go
   ```
   The server will start on port 8080.

## API Endpoints

- `POST /book/` - Create a new book
- `GET /book/{id}` - Get a book by ID
- `GET /books/` - Get all books
- `PUT /book/{id}` - Update a book by ID
- `DELETE /book/{id}` - Delete a book by ID

## Database Configuration

The application connects to a MySQL database. The connection details are configured in the code (check `app.go` or related files for specifics). The Docker Compose file sets up a MariaDB instance with:
- Database: testdb
- User: apiuser
- Password: apipass
- Root Password: secret123