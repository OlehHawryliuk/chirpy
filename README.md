# Chirpy - Production-Ready Twitter Clone Backend

![Go](https://img.shields.io/badge/Go-1.20+-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Production%20Ready-brightgreen)

A production-grade REST API backend for a Twitter-like social platform built with Go. Chirpy demonstrates enterprise-level architecture with secure authentication, database optimization, webhook integration, and comprehensive API design.

## Features

- ✅ **User Management** - Registration, authentication, profile updates
- ✅ **JWT Authentication** - Secure token-based auth with refresh tokens
- ✅ **Chirps (Posts)** - Create, retrieve, delete text posts
- ✅ **Content Moderation** - Automated profanity filtering
- ✅ **Premium Membership** - User upgrade via webhook payments
- ✅ **Webhook Integration** - Payment processing with Polka
- ✅ **Password Security** - bcrypt hashing with salting
- ✅ **API Metrics** - Built-in metrics endpoint
- ✅ **Health Checks** - Readiness and liveness probes
- ✅ **Database Queries** - Type-safe SQL with sqlc
- ✅ **Clean Architecture** - Modular handler-based design

## Tech Stack

- **Language:** Go 1.20+
- **Database:** PostgreSQL 14+
- **Authentication:** JWT (JSON Web Tokens)
- **Password Hashing:** bcrypt
- **SQL Generation:** sqlc
- **HTTP Server:** Go standard library
- **Payment Webhook:** Polka API

## Prerequisites

- Go 1.20 or higher
- PostgreSQL 14 or higher
- curl or Postman (for API testing)
- Git

## Quick Start

### 1. Clone Repository

```bash
git clone https://github.com/OlehHawryliuk/chirpy.git
<<<<<<< HEAD
cp .env.example .env
=======
cd chirpy
>>>>>>> b892cfb (readme fix)
```

### 2. Setup Environment

```bash
cp .env.example .env
# Edit .env with your database credentials
```

### 3. Initialize Database

```bash
# Create PostgreSQL database
createdb chirpy

# Run migrations
psql -d chirpy -f sql/schema.sql
```

### 4. Run Server

```bash
go run main.go
```

Server starts at `http://localhost:8080`

## API Documentation

### Base URL

Fork the repository.
Create a new branch for your feature or bug fix (git checkout -b feature/your-feature-name).
Commit your changes (git commit -m 'Add some feature').
Push to the branch (git push origin feature/your-feature-name).
Open a Pull Request.
http://localhost:8080/api

### Authentication

For protected endpoints, include Bearer token:

```bash
Authorization: Bearer <your_jwt_token>
```

### Core Endpoints

#### Health & Readiness
- `GET /healthz` - Server health check
- `GET /ready` - Readiness probe
- `GET /metrics` - Server metrics

#### User Management
- `POST /users` - Register new user
- `POST /login` - User login (returns JWT)
- `PUT /users` - Update user profile (requires auth)

#### Chirps (Posts)
- `POST /chirps` - Create new chirp (requires auth)
- `GET /chirps` - Get all chirps (optional author filter: `?author_id=<id>`)
- `GET /chirps/{id}` - Get specific chirp
- `DELETE /chirps/{id}` - Delete chirp (owner only)

#### Premium Webhooks
- `POST /polka/webhooks` - Handle payment webhooks (Polka API key required)

## Usage Examples

### Register User

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePassword123"
  }'
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2026-07-31T12:00:00Z",
  "updated_at": "2026-07-31T12:00:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

### Login

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePassword123"
  }'
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "created_at": "2026-07-31T12:00:00Z",
  "updated_at": "2026-07-31T12:00:00Z",
  "email": "user@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "is_chirpy_red": false
}
```

### Create Chirp

```bash
curl -X POST http://localhost:8080/api/chirps \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "body": "Hello Chirpy!"
  }'
```

### Get All Chirps

```bash
curl http://localhost:8080/api/chirps
```

### Delete Chirp

```bash
curl -X DELETE http://localhost:8080/api/chirps/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer <token>"
```

## Project Structure

```text
chirpy/
├── main.go               # Application entry point
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
├── sqlc.yaml             # sqlc configuration
├── .env.example          # Environment template
│
├── internal/             # Internal packages
│   ├── auth.go           # JWT token handling & validation
│   ├── database.go       # PostgreSQL connection & queries
│   ├── models.go         # Data structures
│   └── utils.go          # Helper functions
│
├── handlers/             # HTTP request handlers
│   ├── users.go          # User registration & login
│   ├── chirps.go         # Chirp CRUD operations
│   ├── webhooks.go       # Payment webhook handling
│   ├── health.go         # Health & readiness checks
│   └── auth.go           # Token refresh & validation
│
├── sql/                  # Database schema & queries
│   ├── schema.sql        # Table definitions
│   └── queries.sql       # sqlc query definitions
│
├── assets/               # Static assets & frontend
│   └── index.html        # Frontend application
│
└── README.md             # This file
```

## Security Features

- **JWT Authentication** - Secure token-based authentication with 1-hour expiration
- **Refresh Tokens** - Long-lived tokens for seamless user experience
- **Password Hashing** - bcrypt with automatic salt generation
- **CORS Support** - Configurable cross-origin requests
- **Input Validation** - Request payload validation
- **SQL Injection Prevention** - Parameterized queries via sqlc
- **Webhook Key Verification** - API key validation for payment webhooks

## Health Checks

### Readiness Check

```bash
curl http://localhost:8080/api/ready
```

Returns `200 OK` when server is ready to accept requests.

### Metrics

```bash
curl http://localhost:8080/api/metrics
```

Returns server statistics and performance metrics.

## Database Schema

Main tables:

- **users** - User accounts (email, password_hash, created_at, updated_at, is_chirpy_red)
- **chirps** - User posts (body, user_id, created_at, updated_at)
- **refresh_tokens** - Token management (token, user_id, revoked_at, expires_at)

## Webhook Integration (Polka)

The API integrates with Polka for payment processing:

```bash
POST /api/polka/webhooks
Header: Authorization: ApiKey <polka_api_key>
Body: { "data": { "user_id": "...", "event": "user.upgraded" } }
```

Handles user premium upgrades automatically.

## Error Handling

All endpoints return standardized error responses:

```json
{
  "error": "Error message describing what went wrong"
}
```

Common HTTP Status Codes:
- `200 OK` - Success
- `201 Created` - Resource created
- `400 Bad Request` - Invalid input
- `401 Unauthorized` - Missing or invalid auth
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Testing

Test endpoints using curl, Postman, or any HTTP client:

```bash
# Test server readiness
curl http://localhost:8080/api/ready

# Test user registration and login flow
# See usage examples above
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Learning Outcomes

Building Chirpy demonstrates:

- ✅ RESTful API design principles
- ✅ JWT authentication implementation
- ✅ Database design and SQL optimization
- ✅ Go's `net/http` package mastery
- ✅ Clean code architecture
- ✅ Type-safe SQL with sqlc
- ✅ Password security best practices
- ✅ Webhook integration patterns
- ✅ Error handling strategies
- ✅ API versioning and evolution

## Known Limitations

- Single server deployment (no horizontal scaling)
- In-memory metrics (reset on restart)
- No request rate limiting
- Synchronous webhook processing

## Future Enhancements

- [ ] Implement rate limiting
- [ ] Add request logging middleware
- [ ] Database connection pooling optimization
- [ ] Caching layer (Redis)
- [ ] Advanced search and filtering
- [ ] Image upload support
- [ ] Real-time updates (WebSockets)
- [ ] API documentation (Swagger/OpenAPI)

## License

This project is open source and available under the MIT License.

## Author

**Oleh Hawryliuk**

- GitHub: [@OlehHawryliuk](https://github.com/OlehHawryliuk)
- Email: edifier373@gmail.com
- LinkedIn: [oleh-havryliuk](https://linkedin.com/in/oleh-havryliuk-2b4233419/)

##  Resources

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [JWT.io](https://jwt.io)
- [sqlc](https://sqlc.dev)
- [bcrypt](https://github.com/golang/crypto/tree/master/bcrypt)

##  Support & Questions

For issues, questions, or suggestions:
1. Open an issue on GitHub
2. Check existing documentation
3. Review similar projects for patterns

---

**Made with ❤️ by Oleh Hawryliuk**

**Stars and feedback are appreciated!** ⭐
