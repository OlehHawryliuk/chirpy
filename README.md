# Chirpy - REST API Server

Chirpy is a robust, production-ready backend service (a Twitter clone) written in **Go**. The project implements a complete REST API for user management, secure authentication, short text post creation (Chirps), content moderation, and payment integration via webhooks.

## Tech Stack
- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **SQL Generator:** SQLC (or standard library database/sql)
- **Authentication:** JWT (JSON Web Tokens) with Access & Refresh Token rotation
- **Security:** bcrypt (for secure password hashing)

## Features
- **Users:** Registration, secure login, and profile updates.
- **Chirps:** Create, retrieve (with author filtering and sorting), and delete posts.
- **Content Moderation:** Automatic profanity filter that replaces restricted words with `****`.
- **Security:** Route protection via custom JWT middleware and secure password hashing.
- **Webhooks (Polka):** Handle external payment events to upgrade users to "Chirpy Red" status.
- **Metrics & Admin:** Built-in middleware to track API hits and an admin dashboard to view/reset server state.

## Getting Started

### 1. Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/) (version 1.20+)
- [PostgreSQL](https://www.postgresql.org/)

### 2. Clone the Repository
```bash
git clone https://github.com/OlehHawryliuk/chirpy.git
cd chirpy
```
### 3. Setup Environment Variables
Create your local configuration file by copying the template:

```bash
cp .env.example .env
```

Note: After running this command, open the newly created .env file and replace the placeholder values with your actual PostgreSQL connection details and JWT secrets.

### 4. Run the Server
```bash
go run main.go
```

The server will start listening on port :8080.