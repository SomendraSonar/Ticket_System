```md
# Ticket System API

## Run Locally

go run .

## Health Check

GET /health

## Authentication

POST /auth/register

POST /auth/login

## Tickets

POST /tickets

GET /tickets

GET /tickets/:id

PATCH /tickets/:id/status

## Authorization

Authorization: Bearer <token>

## Docker

docker build -t ticket-system .

docker run -p 8080:8080 ticket-system
```
