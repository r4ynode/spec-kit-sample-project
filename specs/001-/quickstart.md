# Quickstart

**Feature**: Simple Task Management Application

This guide provides instructions to run the application locally using Docker.

## Prerequisites
- Docker
- Docker Compose

## Setup

1.  **Create a `docker-compose.yml` file in the project root:**

    ```yaml
    version: '3.8'
    services:
      backend:
        build:
          context: ./backend
        ports:
          - "8080:8080"
        volumes:
          - ./backend:/app
          - ./data:/data

      frontend:
        build:
          context: ./frontend
        ports:
          - "3000:3000"
        volumes:
          - ./frontend/src:/app/src
        depends_on:
          - backend
    ```

2.  **Create a `Dockerfile` for the backend (`./backend/Dockerfile`):**

    ```dockerfile
    FROM golang:1.19

    WORKDIR /app

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN go build -o main .

    CMD ["/app/main"]
    ```

3.  **Create a `Dockerfile` for the frontend (`./frontend/Dockerfile`):**

    ```dockerfile
    FROM node:18

    WORKDIR /app

    COPY package.json ./
    RUN npm install

    COPY . .

    CMD ["npm", "start"]
    ```

## Running the Application

1.  **Build and start the containers:**

    ```bash
    docker-compose up --build
    ```

2.  **Access the application:**
    - The frontend will be available at `http://localhost:3000`.
    - The backend API will be available at `http://localhost:8080`.

## Validation Scenarios

1.  **Add a task:**
    - Open `http://localhost:3000`.
    - In the input field, type "My first task" and press Enter.
    - **Expected**: The task "My first task" appears in the list.

2.  **Complete a task:**
    - Click the checkbox next to "My first task".
    - **Expected**: The task is marked as complete (e.g., strikethrough).

3.  **Delete a task:**
    - Click the delete icon next to "My first task".
    - **Expected**: The task is removed from the list.
