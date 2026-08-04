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


### Setup

**1. Clone the repo**

```bash
git clone https://github.com/ZakSlinin/cofounders-match-backend.git
cd cofounders-match-backend
```

**2. Build and start**

```bash
docker-compose up --build -d
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

| Method | Endpoint         | Description       |
|--------|------------------|-------------------|
| POST   | /profiles        | Create Profile    |
| POST   | /profiles/me     | Get my profile    |
| PATCH  | /profiles/me     | Update my profile |
| POST   | /profiles/:id    | Get profile by ID |
| POST   | /profiles/avatar | Upload Avatar     |
| GET    | /feed            | Get feed          |


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