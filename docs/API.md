# Coffee Cups System API Documentation

## Overview

The Coffee Cups System provides a REST API for managing coffee consumption tracking and cost distribution among colleagues.

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

Currently, the API does not require authentication. In a production environment, you should implement proper authentication.

## Endpoints

### Users

#### GET /users
Get all active users.

**Response:**
```json
[
  {
    "id": 1,
    "telegram_id": 123456789,
    "username": "john_doe",
    "first_name": "John",
    "last_name": "Doe",
    "is_active": true,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

#### GET /users/{id}
Get a specific user by Telegram ID.

**Parameters:**
- `id` (path): Telegram ID of the user

**Response:**
```json
{
  "id": 1,
  "telegram_id": 123456789,
  "username": "john_doe",
  "first_name": "John",
  "last_name": "Doe",
  "is_active": true,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Boxes

#### GET /boxes
Get all active coffee boxes.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Premium Coffee Blend",
    "total_cups": 20,
    "price": 15.99,
    "is_active": true,
    "created_by": 1,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

#### POST /boxes
Create a new coffee box.

**Request Body:**
```json
{
  "name": "Premium Coffee Blend",
  "total_cups": 20,
  "price": 15.99,
  "created_by": 1
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Premium Coffee Blend",
  "total_cups": 20,
  "price": 15.99,
  "is_active": true,
  "created_by": 1,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

#### GET /boxes/{id}
Get a specific box by ID.

**Parameters:**
- `id` (path): Box ID

### Coffee Logs

#### GET /coffee-logs
Get coffee logs for a user.

**Query Parameters:**
- `user_id` (required): User ID

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "box_id": 1,
    "logged_at": "2023-01-01T10:00:00Z",
    "created_at": "2023-01-01T10:00:00Z",
    "updated_at": "2023-01-01T10:00:00Z"
  }
]
```

#### POST /coffee-logs
Log a coffee consumption.

**Request Body:**
```json
{
  "user_id": 1,
  "box_id": 1
}
```

**Response:**
```json
{
  "id": 1,
  "user_id": 1,
  "box_id": 1,
  "logged_at": "2023-01-01T10:00:00Z",
  "created_at": "2023-01-01T10:00:00Z",
  "updated_at": "2023-01-01T10:00:00Z"
}
```

### Payments

#### GET /payments
Get payments for a user or box.

**Query Parameters:**
- `user_id`: User ID (optional)
- `box_id`: Box ID (optional)

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "box_id": 1,
    "amount": 2.50,
    "is_paid": false,
    "paid_at": null,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

## Health Check

#### GET /health
Check if the service is running.

**Response:**
```
OK
```

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE"
}
```

**Common HTTP Status Codes:**
- `200` - Success`
- `201` - Created`
- `400` - Bad Request`
- `404` - Not Found`
- `500` - Internal Server Error`
