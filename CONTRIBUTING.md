# Contributing to OSArk Server

Thank you for your interest in contributing to OSArk Server! This document provides guidelines and information about the project's structure and workflow.

## Project Structure

The project follows a clean architecture pattern with the following main components:

```
osark-server/
├── cmd/               # Application entry points
├── internal/          # Private application code
│   ├── config/        # Configuration management
│   ├── models/        # Data models and structures
│   ├── repository/    # Data access layer
│   ├── server/        # HTTP server and handlers
│   └── service/       # Business logic
├── views/             # HTML templates
└── ...
```

### Internal Package

The `internal` package contains the core application code, organized by responsibility:

#### Config (`internal/config`)

Configuration management for the application:

- `config.go`: Defines environment variable parsing and configuration structures.
  - Uses `env.Parse()` to load configurations from environment variables
  - Provides accessor methods for different configuration parameters
  - Handles default values for certain configurations (e.g., `VIEW_PATH`)

#### Models (`internal/models`)

Data structures used throughout the application:

- `sysinfo.go`: System information data structures
- `apps.go`: Application data structures
- `events.go`: Event data structures
- `query.go`: Query parameters for API requests
- `data.go`: Generic data structures

#### Repository (`internal/repository`)

Data access layer that handles database operations:

- `sysinfo.go`: Database operations for system information
- `apps.go`: Database operations for applications
- `processes.go`: Database operations for processes

#### Server (`internal/server`)

HTTP server setup and request handlers:

- `server.go`: Server initialization and configuration (using Fiber framework)
- `handler.go`: Base handler structure and dependencies
- `routes.go`: Route definitions and endpoint mapping
- `systems.go`: System-related request handlers
- `views.go`: View-related request handlers
- `events.go`: Event handling

#### Service (`internal/service`)

Business logic layer:

- `data.go`: Core data processing logic
- `event.go`: Event processing and handling logic

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Access to a PostgreSQL database (can be configured via environment variables)

### Environment Setup

The application uses environment variables for configuration. Key variables include:

```
PORT=8080                  # Server port
DB_HOST=localhost          # Database host
DB_PORT=5432               # Database port
DB_USER=postgres           # Database user
DB_PASSWORD=password       # Database password
DB_NAME=osark             # Database name
VIEW_PATH=/path/to/views   # Path to HTML templates
```

### Running Locally

1. Clone the repository:

   ```
   git clone https://github.com/your-username/osark-server.git
   cd osark-server
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Run the application:

   ```
   go run cmd/api/main.go
   ```

### Docker

The project includes a Dockerfile for containerized deployment:

```
docker build -t osark-server .
docker run -p 8080:8080 osark-server
```

## Development Guidelines

### Code Style

- Follow Go standard formatting using `gofmt`
- Use meaningful variable and function names
- Write comments for exported functions and complex logic

### Testing

- Write unit tests for new functionality
- Ensure existing tests pass before submitting a pull request
- Use Go's built-in testing framework

### Commit Messages

Follow conventional commit format:

- `feat: add new feature`
- `fix: resolve issue with X`
- `docs: update documentation`
- `refactor: improve code structure`
- `test: add tests for feature Y`

### Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes with descriptive messages
4. Push your branch to your fork
5. Submit a pull request to the main repository

### Project Workflow

1. Issues are created for bugs or feature requests
2. Contributors pick issues to work on
3. Pull requests are reviewed by maintainers
4. Approved changes are merged into the main branch

## Views and Templates

The project uses the Fiber HTML template engine for rendering views. Template files are located in the `views/` directory and organized by functionality.

When contributing to templates:

- Maintain consistent styling and layout
- Test templates across different device sizes
- Follow the established pattern for template organization

## License

By contributing to OSArk Server, you agree that your contributions will be licensed under the project's license.

## Contact

If you have questions or need help, please reach out to the maintainers through GitHub issues.
