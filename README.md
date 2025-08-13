# Go-Tweets API (`register` branch)

This document provides detailed information about the Go-Tweets API, specifically for the `register` branch. This branch focuses on user registration and email verification functionalities.

## Table of Contents

1.  [Overview](#overview)
2.  [Features](#features)
3.  [Prerequisites](#prerequisites)
4.  [Getting Started](#getting-started)
    - [Environment Configuration](#environment-configuration)
    - [Database Migration](#database-migration)
5.  [Makefile Usage](#makefile-usage)
6.  [Docker Usage](#docker-usage)
7.  [API Specification](#api-specification)
    - [POST /api/users/register](#post-apiusersregister)
    - [GET /api/users/verify-email](#get-apiusersverify-email)
8.  [API Usage Examples](#api-usage-examples)
    - [Register a New User](#register-a-new-user)
    - [Verify Email Address](#verify-email-address)

## Overview

Go-Tweets is a backend service built with Go that aims to provide core functionalities for a Twitter-like application. This `register` branch contains the initial setup for user creation, including API endpoints for registration and a mechanism for verifying user emails via a token-based system.

The application is containerized using Docker and managed with Docker Compose for easy setup and development. Database migrations are handled by `dbmate`.

## Features

- **User Registration:** Allows new users to create an account.
- **Email Verification:** Sends a verification link to the user's email address and provides an endpoint to confirm it.
- **Containerized Environment:** Fully configured to run within Docker containers.
- **Database Migration Management:** Uses `dbmate` for structured database schema changes.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/doc/install) (for local development outside Docker)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- `make` command-line utility

## Getting Started

Follow these steps to get the application running.

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/kucingscript/go-tweets/tree/register
    cd go-tweets
    ```

2.  **Configure Environment Variables:**
    Create a `.env` file in the root of the project by copying the example below. These variables are essential for connecting to the database, JWT signing, and email (SMTP) services.

    ```env
    # Application
    APP_PORT=8080

    # Database (PostgreSQL)
    DB_HOST=db
    DB_PORT=5432
    POSTGRES_USER=your_db_user
    POSTGRES_PASSWORD=your_db_password
    POSTGRES_DB=go_tweets_dev
    DATABASE_URL="postgres://your_db_user:your_db_password@db:5432/go_tweets_dev?sslmode=disable"

    # JWT
    JWT_SECRET=your_jwt_secret_key

    # SMTP for Mailer
    SMTP_HOST=your_smtp_host
    SMTP_PORT=587
    SMTP_USER=your_smtp_username
    SMTP_PASS=your_smtp_password
    SMTP_SENDER="Your Name <no-reply@example.com>"
    ```

    **Note:** Ensure the `DATABASE_URL` credentials match the `POSTGRES_USER` and `POSTGRES_PASSWORD` variables.

3.  **Run Database Migrations:**
    With the Docker daemon running, use the Makefile to set up the database schema.

    ```sh
    make db-up
    ```

4.  **Start the Application:**
    Use Docker Compose to build and start all services (Go app, PostgreSQL DB, Adminer).
    ```sh
    make up
    ```
    The API will be accessible at `http://localhost:8080`.

## Makefile Usage

The `Makefile` provides several commands to simplify development and management tasks.

| Command             | Description                                                   |
| ------------------- | ------------------------------------------------------------- |
| `build`             | Builds the Docker images for the services.                    |
| `up`                | Starts all services in detached mode and builds if necessary. |
| `down`              | Stops all services and removes the data volumes.              |
| `logs`              | Tails the logs for all running services.                      |
| `logs-app`          | Tails the logs specifically for the Go application container. |
| `restart-app`       | Rebuilds and restarts only the Go application service.        |
| `db-up`             | Runs all pending database migrations.                         |
| `db-down`           | Rolls back the most recent database migration.                |
| `db-new name=<...>` | Creates a new migration file. `name` is a required argument.  |
| `db-status`         | Shows the status of all database migrations.                  |

## Docker Usage

The application is managed via `docker-compose.yml`.

- **Services:**

  - `app`: The main Go application.
  - `db`: The PostgreSQL database.
  - `adminer`: A web-based database management tool accessible at `http://localhost:8081/adminer`.
  - `dbmate`: A service to run database migrations.

- **Start all services:**

  ```sh
  docker-compose up -d
  ```

- \*\*Stop all services:

  ```sh
  docker-compose down
  ```

- **Build the application image:**
  ```sh
  docker-compose build app
  ```

## API Specification

### POST /api/users/register

Registers a new user in the system. It validates the input, checks for existing users with the same email or username, and sends a verification email upon successful registration.

- **Method:** `POST`
- **Endpoint:** `/api/users/register`
- **Headers:**

  - `Content-Type: application/json`

- **Request Body:**

  ```json
  {
    "email": "user@example.com",
    "username": "newuser",
    "password": "password123",
    "password_confirm": "password123"
  }
  ```

- **Success Response (201 Created):**

  ```json
  {
    "id": 1,
    "email": "user@example.com",
    "username": "newuser",
    "is_verified": false,
    "created_at": "2025-08-13T10:00:00Z"
  }
  ```

- **Error Responses:**
  - `400 Bad Request`: Invalid request body, validation errors (e.g., password mismatch, weak password), or missing fields.
  - `409 Conflict`: A user with the provided email or username already exists.

### GET /api/users/verify-email

Verifies a user's email address using the token sent to them after registration.

- **Method:** `GET`
- **Endpoint:** `/api/users/verify-email`
- **Query Parameters:**

  - `token` (string, required): The verification token from the email link.

- **Success Response (200 OK):**

  ```json
  {
    "message": "Email verified successfully"
  }
  ```

- **Error Responses:**
  - `400 Bad Request`: The `token` query parameter is missing.
  - `404 Not Found`: The verification token is invalid or has expired.

## API Usage Examples

### Register a New User

```sh
curl -X POST http://localhost:8080/api/users/register \
-H "Content-Type: application/json" \
-d '{
  "email": "test.user@example.com",
  "username": "testuser",
  "password": "a-strong-password",
  "password_confirm": "a-strong-password"
}'
```

### Verify Email Address

After registration, a link is sent to the user's email. The link will look something like this: `http://yourapp.com/verify-email?token=SOME_VERIFICATION_TOKEN`.

To simulate this verification, use the token with the endpoint:

```sh
curl -X GET http://localhost:8080/api/users/verify-email?token=YOUR_VERIFICATION_TOKEN
```

```

```
