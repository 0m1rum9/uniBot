# uniBot

Telegram bot interface for interacting with the UNI-X online lesson platform.

The bot stores a user's platform credentials, lists available modules, courses, and topics, then can automate lesson watch events and quiz completion through the platform API.

> Use this project only with accounts, platforms, and automation flows you are authorized to access. Automating lesson progress or quizzes may violate university, platform, or course policies.

## Features

- Telegram bot built with `telebot.v4`
- User registration with UNI-X login and password
- PostgreSQL storage for Telegram users and platform credentials
- Redis client setup for repository/cache support
- Module, course, and topic navigation through inline Telegram buttons
- Lesson watch event automation with progress updates
- Quiz start/check API calls after lessons
- SQL migrations managed with Goose
- SQL query code generated with sqlc
- Local PostgreSQL and Redis services through Docker Compose

## Tech Stack

- Go `1.25.5`
- PostgreSQL
- Redis
- Docker Compose
- Goose migrations
- sqlc
- Telegram Bot API via `gopkg.in/telebot.v4`

## Project Structure

```text
.
├── main.go                         # Application entry point
├── docker-compose.yml              # Local PostgreSQL and Redis services
├── Makefile                        # Run, migration, and Docker helpers
├── sqlc.yaml                       # sqlc configuration
├── internal/
│   ├── db/                         # sqlc queries, models, migrations
│   ├── handlers/                   # Telegram command/callback handlers
│   ├── middlewares/                # Registration middleware
│   ├── repository/                 # User repository layer
│   ├── cache/                      # Cache helpers
│   └── services/                   # Notification service
└── pkg/adapters/unix/              # UNI-X API adapter
```

## Requirements

- Go installed
- Docker and Docker Compose
- Telegram bot token from BotFather
- `goose` CLI for migrations
- `sqlc` CLI if you need to regenerate database code

Install common tooling:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Configuration

Create a `.env` file in the project root:

```env
BOT_TOKEN=your_telegram_bot_token

POSTGRES_USER=unibot
POSTGRES_PASSWORD=unibot_password
POSTGRES_DB=unibot
POSTGRES_HOST=localhost
POSTGRES_PORT=5432

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=redis_password
REDIS_USER=unibot
REDIS_USER_PASSWORD=redis_user_password
```

The application reads `POSTGRES_DB` at runtime. The current `Makefile` migration URL uses `POSTGRES_NAME`, so either add this alias to `.env`:

```env
POSTGRES_NAME=unibot
```

or update the `Makefile` to use `POSTGRES_DB`.

## Running Locally

Start PostgreSQL and Redis:

```bash
make docker-up
```

Run database migrations:

```bash
make migrate-up
```

Start the bot:

```bash
make run
```

## Bot Usage

1. Open the bot in Telegram.
2. Send `/start`.
3. Register UNI-X credentials:

   ```text
   /register <login> <password>
   ```

4. After registration, send any text message to open the module menu.
5. Select a module.
6. Select a course.
7. Select an available topic.
8. The bot will process lessons in order and update the Telegram message with progress.

Topics can only be processed when previous required topics are already passed.

## Database

The project has one migration that creates a `users` table:

```sql
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username TEXT,
    chat_id BIGINT,
    login TEXT,
    password TEXT
);
```

Generated sqlc files are stored in `internal/db`.

Regenerate database code after changing SQL queries or schema:

```bash
sqlc generate
```

## Make Targets

```bash
make docker-up    # Start PostgreSQL and Redis containers
make migrate-up   # Apply Goose migrations
make run          # Run the bot locally
```


