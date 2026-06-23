# Chirpy - REST API Server

Chirpy is a robust, production-ready backend service written in **Go**. The project implements a complete REST API for user management, secure authentication, short text post creation (Chirps), content moderation, and payment integration via webhooks.

## Motivation
The motivation behind this project is to build a scalable, secure, and performant REST API from scratch using Go's standard library. By avoiding heavy third-party frameworks, this project demonstrates a deep understanding of HTTP protocol fundamentals, database integration with PostgreSQL, secure password hashing, and clean architectural design in Go.

## Quick Start

### 1. Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/) (version 1.20+)
- [PostgreSQL](https://www.postgresql.org/)

### 2. Clone and Setup
```bash
git clone https://github.com/OlehHawryliuk/chirpy.git
cd chirpy
cp .env.example .env
```

### 3. Run the Server
```bash
go run main.go
```
The server will start listening on port :8080.

### Usage
Once the server is running, you can interact with the API endpoints. Below are the key available routes:

Method	Endpoint	         Description	Auth Required
GET	    /api/healthz	     Check server readiness	Public
POST	/api/users	         Register a new user	Public
POST	/api/login	         User login (returns JWT)	Public
POST	/api/chirps	         Create a new Chirp	Private (JWT)
GET	    /api/chirps	         Retrieve all Chirps	Public
DELETE	/api/chirps/{id}     Delete a Chirp	Private (Owner only)
POST	/api/polka/webhooks	 Handle Polka payment webhooks	Private (Polka Key)
You can use curl, Postman, or any HTTP client to send requests to http://localhost:8080.

### Contributing
Contributions are welcome! If you would like to contribute to this project, please follow these steps:

Fork the repository.
Create a new branch for your feature or bug fix (git checkout -b feature/your-feature-name).
Commit your changes (git commit -m 'Add some feature').
Push to the branch (git push origin feature/your-feature-name).
Open a Pull Request.