# CoFounders Match — User Service

Auth + Profile API.

## Stack

- Go + Gin
- PostgreSQL
- JWT (access 15 min, refresh 30 days)
- Yandex Cloud Object Storage (avatars)

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Setup

**1. Clone the repo**

```bash
git clone https://github.com/ZakSlinin/cofounders-match-backend.git
cd cofounders-match-backend
```

**2. Create `user-service/.env.docker`**

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

YC_ACCESS_KEY=your_yc_access_key
YC_SECRET_KEY=your_yc_secret_key
YC_BUCKET=cofounders-match-avatars
```

**3. Build and start**

```bash
docker-compose up --build -d
```

**4. Run migrations**

```bash
cd user-service
make migrate-up
```

Service is available at `http://localhost:8080`

---

## API

### Auth

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /auth/register | Register |
| POST | /auth/login | Login |
| POST | /auth/refresh | Refresh access token |

### Profile

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /profiles | ✓ | Create profile |
| POST | /profiles/avatar | ✓ | Upload avatar |

---

## Authorization

Protected endpoints require:

```
Authorization: Bearer <access_token>
```

Access token expires in **15 minutes**. On `401` use refresh token to get a new one.

---

## Full API Docs

Open `openapi.yaml` in [Swagger Editor](https://editor.swagger.io).