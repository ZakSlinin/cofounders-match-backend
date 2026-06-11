# CoFounders Match — User Service

Auth API: registration, login, token refresh.

## Stack

- Go + Gin
- PostgreSQL
- JWT (access 15 min, refresh 30 days)

## Getting Started

### What you need

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI

Install migrate CLI:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Setup

**1. Clone the repo**

```bash
git clone https://github.com/ZakSlinin/cofounders-match-backend.git
cd cofounders-match-backend
```

**2. Create `.env.docker` inside `user-service/`**

```env
DB_USER=postgres
DB_PASSWORD=your_password
DB_HOST=postgres
DB_PORT=5432
DB_NAME=cofounders_match
DB_URL=postgres://postgres:your_password@postgres:5432/cofounders_match?sslmode=disable

PORT=8080

JWT_SECRET=your_jwt_secret
JWT_REFRESH_SECRET=your_refresh_secret
```

**3. Start the database**

```bash
docker-compose up --build -d
```

**4. Run migrations**

```bash
cd user-service
make migrate-up
```

Service is available at `http://localhost:8080`

## Authorization

All protected endpoints require the header:

```
Authorization: Bearer <access_token>
```

Access token expires in **15 minutes**. Use refresh token to get a new one when you receive `401`.

## Full API Docs

Open `openapi.yaml` in [Swagger Editor](https://editor.swagger.io) for interactive documentation.