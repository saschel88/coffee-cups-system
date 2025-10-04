# Coffee Cups System ☕

A Telegram-based system for tracking coffee capsule consumption and fair cost distribution among colleagues. Built with Go, this system allows colleagues to log each capsule use via a Telegram bot, then automatically calculates and distributes the box cost based on actual usage.

## Features

- 🤖 **Telegram Bot Integration**: Easy coffee logging via Telegram commands
- 📊 **Usage Tracking**: Track individual coffee consumption
- 💰 **Fair Cost Distribution**: Automatically calculate each person's share
- 📱 **REST API**: Full API for web/mobile integration
- 🗄️ **PostgreSQL Database**: Reliable data storage
- 🐳 **Docker Support**: Easy deployment with Docker Compose
- 📈 **Usage Analytics**: Track consumption patterns and costs

## Project Structure

```
coffee-cups-system/
├── cmd/                    # Application entry points
│   ├── server/            # Main application server
│   └── migrate/           # Database migration tool
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── database/         # Database connection and setup
│   ├── handlers/         # HTTP request handlers
│   ├── logger/           # Logging utilities
│   ├── models/           # Data models
│   ├── services/         # Business logic services
│   ├── server/           # HTTP server setup
│   └── telegram/         # Telegram bot implementation
├── configs/              # Configuration files
├── docs/                 # Documentation
├── scripts/              # Setup and utility scripts
├── tests/                # Test files
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile           # Docker image definition
├── Makefile            # Build and development commands
└── go.mod              # Go module definition
```

## Quick Start

### Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Docker and Docker Compose (optional)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd coffee-cups-system

# Install dependencies
go mod download
```

### 2. Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE coffee_cups;
CREATE USER coffee_user WITH PASSWORD 'coffee_password';
GRANT ALL PRIVILEGES ON DATABASE coffee_cups TO coffee_user;
```

### 3. Configuration

Copy and update the configuration:

```bash
cp configs/config.yaml.example configs/config.yaml
```

Update the configuration with your database credentials and Telegram bot token.

### 4. Run Database Migrations

```bash
go run ./cmd/migrate
```

### 5. Start the Application

```bash
go run ./cmd/server
```

## Docker Deployment

### Using Docker Compose

```bash
# Set your Telegram bot token
export TELEGRAM_BOT_TOKEN=your_bot_token_here

# Start all services
docker-compose up -d
```

This will start:
- The Coffee Cups System application
- PostgreSQL database
- All necessary services

## Telegram Bot Usage

Once the bot is running, users can interact with it using these commands:

- `/start` - Start using the bot
- `/coffee <box_id>` - Log a coffee consumption
- `/status` - View your recent coffee logs
- `/boxes` - View available coffee boxes
- `/help` - Show help message

### Example Workflow

1. Admin creates a coffee box: "Premium Blend - 20 cups - $15.99"
2. Users log coffee: `/coffee 1` (where 1 is the box ID)
3. System tracks usage and calculates individual costs
4. Users can check their status: `/status`
5. System generates payment records for fair cost distribution

## API Endpoints

The system provides a REST API for integration:

- `GET /api/v1/users` - Get all users
- `GET /api/v1/boxes` - Get all coffee boxes
- `POST /api/v1/boxes` - Create a new box
- `GET /api/v1/coffee-logs` - Get coffee logs
- `POST /api/v1/coffee-logs` - Log coffee consumption
- `GET /api/v1/payments` - Get payment information

See [API Documentation](docs/API.md) for detailed endpoint information.

## Development

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Building

```bash
# Build the application
make build

# Run the application
make run
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint
```

## Configuration

The application can be configured via:

1. **Configuration file**: `configs/config.yaml`
2. **Environment variables**: Set in `.env` file
3. **Command line flags**: (future enhancement)

### Key Configuration Options

- **Database**: PostgreSQL connection settings
- **Telegram**: Bot token and debug mode
- **Server**: HTTP server host and port
- **Logging**: Log level and format

## Deployment

See [Deployment Guide](docs/DEPLOYMENT.md) for detailed deployment instructions including:

- Local development setup
- Docker deployment
- Production deployment
- Monitoring and logging
- Backup strategies

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions:

1. Check the [documentation](docs/)
2. Review the [API documentation](docs/API.md)
3. Open an issue on GitHub

## Roadmap

- [ ] Web dashboard for administration
- [ ] Mobile app integration
- [ ] Advanced analytics and reporting
- [ ] Multi-language support
- [ ] Integration with payment systems
- [ ] Automated notifications
- [ ] User authentication and authorization