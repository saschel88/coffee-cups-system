# Deployment Guide

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Docker and Docker Compose (optional)

## Local Development Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd coffee-cups-system
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Database

Create a PostgreSQL database:

```sql
CREATE DATABASE coffee_cups;
CREATE USER coffee_user WITH PASSWORD 'coffee_password';
GRANT ALL PRIVILEGES ON DATABASE coffee_cups TO coffee_user;
```

### 4. Configure Environment

Copy the example configuration and update it:

```bash
cp configs/config.yaml.example configs/config.yaml
```

Update the configuration with your database credentials and Telegram bot token.

### 5. Run Database Migrations

```bash
go run ./cmd/migrate
```

### 6. Start the Application

```bash
go run ./cmd/server
```

## Docker Deployment

### 1. Using Docker Compose

```bash
# Set your Telegram bot token
export TELEGRAM_BOT_TOKEN=your_bot_token_here

# Start all services
docker-compose up -d
```

### 2. Using Docker

```bash
# Build the image
docker build -t coffee-cups-system .

# Run with environment variables
docker run -d \
  --name coffee-cups-system \
  -p 8080:8080 \
  -e TELEGRAM_BOT_TOKEN=your_bot_token_here \
  -e DATABASE_HOST=your_db_host \
  -e DATABASE_PORT=5432 \
  -e DATABASE_USER=coffee_user \
  -e DATABASE_PASSWORD=coffee_password \
  -e DATABASE_DBNAME=coffee_cups \
  coffee-cups-system
```

## Production Deployment

### 1. Prepare the Server

1. Install Docker and Docker Compose
2. Create a dedicated user for the application
3. Set up SSL certificates (if using HTTPS)
4. Configure firewall rules

### 2. Deploy the Application

```bash
# Clone the repository
git clone <repository-url>
cd coffee-cups-system

# Set up environment variables
cp .env.example .env
# Edit .env with production values

# Start the application
docker-compose -f docker-compose.prod.yml up -d
```

### 3. Set Up Reverse Proxy (Nginx)

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 4. Set Up SSL (Let's Encrypt)

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get SSL certificate
sudo certbot --nginx -d your-domain.com
```

## Monitoring and Logging

### 1. Application Logs

Logs are written to stdout in JSON format. For production, consider using a log aggregation service.

### 2. Health Checks

The application provides a health check endpoint:

```bash
curl http://localhost:8080/health
```

### 3. Database Monitoring

Monitor your PostgreSQL database for:
- Connection count
- Query performance
- Disk usage
- Backup status

## Backup Strategy

### 1. Database Backups

```bash
# Create backup
pg_dump -h localhost -U coffee_user coffee_cups > backup_$(date +%Y%m%d_%H%M%S).sql

# Restore backup
psql -h localhost -U coffee_user coffee_cups < backup_20230101_120000.sql
```

### 2. Automated Backups

Set up a cron job for automated backups:

```bash
# Add to crontab
0 2 * * * /usr/local/bin/backup_coffee_cups.sh
```

## Security Considerations

1. **Environment Variables**: Never commit sensitive data to version control
2. **Database Security**: Use strong passwords and limit network access
3. **API Security**: Implement authentication and rate limiting
4. **SSL/TLS**: Always use HTTPS in production
5. **Updates**: Keep dependencies and system packages updated

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Check database credentials
   - Verify database is running
   - Check network connectivity

2. **Telegram Bot Not Responding**
   - Verify bot token is correct
   - Check if bot is running
   - Verify webhook configuration (if using webhooks)

3. **High Memory Usage**
   - Check for memory leaks
   - Monitor database connections
   - Review application logs

### Log Analysis

```bash
# View application logs
docker-compose logs -f app

# View database logs
docker-compose logs -f postgres
```

## Performance Optimization

1. **Database Indexing**: Ensure proper indexes on frequently queried columns
2. **Connection Pooling**: Configure appropriate connection pool settings
3. **Caching**: Implement caching for frequently accessed data
4. **Load Balancing**: Use multiple application instances behind a load balancer
