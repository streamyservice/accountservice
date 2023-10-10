# User Account Management Microservice

[![Build and Deploy](https://github.com/streamyservice/accountservice/actions/workflows/go.yml/badge.svg)](https://github.com/streamyservice/accountservice/actions/workflows/go.yml)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

The User Account Management Microservice is a Go-based microservice that provides functionality for user registration,
authentication, and account management.

## Features

- User registration
- User authentication (JWT-based)
- Profile management
- Password reset
- User roles and permissions
- API documentation (Swagger/OpenAPI)

## Technologies Used

- Go
- Gin (Web Framework)
- GORM (ORM for Go)
- JWT (JSON Web Tokens)
- PostgreSQL (or your preferred database)
- Docker

## Getting Started

### Prerequisites

- Go (>= 1.21)
- PostgreSQL (or another database of your choice)
- Docker (for containerization)

### Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:streamyservice/accountservice.git
   cd accountservice

2. Set up your configuration. Copy the example configuration file and modify it as needed:

   ```bash
   cp example.env env
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Create and migrate the database:

   ```bash
   go run migrations/migrate.go
   ```

5. Build and run the microservice:

   ```bash
   go run main.go
   ```

### Usage

- Access the API documentation at `http://localhost:8080/swagger/index.html` for detailed information on available
  endpoints and how to use them.

### Configuration

You can configure the microservice by modifying the `config.yaml` file. This includes setting up your database
connection, JWT secret, and other environment-specific settings.

### Deployment

For production deployment, it's recommended to containerize the microservice using Docker and use an orchestration tool
like Kubernetes or Docker Compose for scalability and reliability.

## Contributing

Contributions are welcome! Please follow the [contribution guidelines](CONTRIBUTING.md) to get started.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - Web framework for Go
- [GORM](https://gorm.io/) - Object-relational mapping for Go
- [JWT-Go](https://github.com/dgrijalva/jwt-go) - JSON Web Tokens for Go
- [Swagger](https://swagger.io/) - API documentation tool