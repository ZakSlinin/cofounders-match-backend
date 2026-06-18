# CoFounders Match

A platform for finding co-founders, teammates, and collaborators for startups.

## What is this?

CoFounders Match helps founders, developers, and designers find each other based on skills, goals, and project ideas.

## Architecture

```
cofounders-match-backend/
  user-service/        — auth + profiles + file uploads
  match-service/       — swipes and matches (coming soon)
  recommendation-service/ — feed and scoring (coming soon)
  chat-service/        — messaging (coming soon)
  docker-compose.yml   — runs everything
```

Microservices communicate via REST. Each service has its own database.

## Services

| Service | Status | Description |
|---------|--------|-------------|
| user-service | ✅ Active | Auth, profiles, avatars |
| match-service | 🔜 Planned | Swipes, matches |
| recommendation-service | 🔜 Planned | Feed, scoring algorithm |
| chat-service | 🔜 Planned | Real-time messaging |

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **Cache:** Redis
- **Storage:** Yandex Cloud Object Storage
- **Infrastructure:** Docker, Docker Compose

## Getting Started

```bash
git clone https://github.com/ZakSlinin/cofounders-match-backend.git
cd cofounders-match-backend
docker-compose up --build -d
```

See individual service READMEs for setup details:
- [user-service](./user-service/README.md)

## API Docs

Each service includes an `openapi.yaml`. Open it in [Swagger Editor](https://editor.swagger.io) for interactive docs.

## Author

Built by [@ZakSlinin](https://github.com/ZakSlinin)
