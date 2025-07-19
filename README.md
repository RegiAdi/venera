# Venera

[![unit-test](https://github.com/RegiAdi/venera/actions/workflows/unit_test.yml/badge.svg)](https://github.com/RegiAdi/venera/actions/workflows/unit_test.yml)
[![codecov](https://codecov.io/gh/RegiAdi/venera/graph/badge.svg?token=M5SJRT8ZSF)](https://codecov.io/gh/RegiAdi/venera)
[![build](https://github.com/RegiAdi/venera/actions/workflows/build.yml/badge.svg)](https://github.com/RegiAdi/venera/actions/workflows/build.yml)

A modern, high-performance HTTP server built with Go Fiber and MongoDB, featuring robust authentication, product management, and scalable architecture.

> 🤖 **AI Collaboration Notice**: This project was developed with the assistance of AI tools, including GitHub Copilot, to enhance code quality, documentation, and development efficiency.

## Features

- 🚀 Built with [Go Fiber](https://gofiber.io/) for high performance
- 📦 MongoDB integration for flexible data storage
- 🔐 Authentication system
- 🛍️ Product management system
- ✅ Comprehensive test coverage
- 🔄 Graceful shutdown handling
- 🎯 Clean architecture with clear separation of concerns

## Prerequisites

- Go 1.21 or higher
- MongoDB
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/RegiAdi/venera.git
   cd venera
   ```

2. Set up your environment variables by copying the example:
   ```bash
   cp .env.example .env
   ```
   Then edit the `.env` file with your configuration.

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the server:
   ```bash
   go run main.go
   ```
   Or use Make:
   ```bash
   make run
   ```

## Project Structure

```
.
├── config/         # Application configuration
├── handlers/       # HTTP request handlers
├── helpers/        # Utility functions
├── kernel/         # Core application setup
├── middleware/     # HTTP middleware
├── models/         # Data models
├── repositories/   # Data access layer
├── responses/      # HTTP response structures
├── routes/         # Route definitions
└── services/       # Business logic layer
```

## API Documentation

The API provides the following main endpoints:

### Authentication
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration

### Products
- `GET /api/products` - List all products
- `POST /api/products` - Create a new product
- `GET /api/products/:id` - Get a specific product
- `PUT /api/products/:id` - Update a product
- `DELETE /api/products/:id` - Delete a product

## Testing

Run the test suite:
```bash
go test ./...
```

Or with coverage:
```bash
make test
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
