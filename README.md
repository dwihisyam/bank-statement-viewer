# Bank Statement Viewer -- Fullstack Web Application

Bank Statement Viewer is a full-stack web application consisting of a frontend, backend,
and complete Docker-based orchestration. The project follows a modular
and scalable architecture designed for modern development workflows.

## 📘 About the Project

Bank Statement Viewer is designed with a fully decoupled frontend and backend
architecture. This separation ensures scalability, clean
maintainability, and easy deployment. The project includes Docker
support for unified builds and simplified distribution.

## 🔧 Technologies Used

### Frontend

- Next.js

### Backend

- Go

### Environment / DevOps

- Docker & Docker Compose
- GitHub Actions CI/CD
- Environment variables via .env

## 📁 Folder Structure

    bank_statement_viewer/
    ├── backend/
    │   ├── cmd/
    │   │   └── app/
    │   │       └── main.go
    │   ├── internal/
    │   │   ├── handler/
    │   │   │   ├── transaction_handler_tes.go
    │   │   │   └── transaction_handler.go
    │   │   ├── model/
    │   │   │   └── transaction.go
    │   │   ├── repository/
    │   │   │   ├── transaction_repo_mock.go
    │   │   │   └── transaction_repo.go
    │   │   ├── service/
    │   │   │   ├── transaction_service_mock.go
    │   │   │   ├── transaction_service_test.go
    │   │   │   └── transaction_service.go
    │   │   └── utils/
    │   │       └── helper.go
    │   ├── pkg/
    │   │   └── response/
    │   │       └── response.go
    │   ├── go.mod
    │   ├── go.sum
    │   └── Dockerfile
    │
    ├── frontend/
    │   ├── app/
    │   │   ├── layout.tsx
    │   │   └── page.tsx
    │   ├── components/
    │   │   ├── BalanceCard.tsx
    │   │   ├── FileUploader.tsx
    │   │   ├── Snackbar.tsx
    │   │   └── Table.tsx
    │   ├── utils/
    │   │   └── api.ts
    │   ├── package.json
    │   ├── .env
    │   └── Dockerfile
    │
    ├── docker-compose.yml
    ├── .gitignore
    └── README.md

# 🛠️ Setup Instructions

## 1. Clone the Repository

    git clone <your repo url>
    cd bank-statement-viewer

## 2. Ensure Required Tools are Installed

    Go
    Node.js
    Npm
    Docker
    Docker Compose
    Git

## 3. Install Dependencies

### Backend

    cd backend/cmd/app
    go get

### Frontend

    cd frontend
    npm install

## 4. Create Environment Files

### Backend .env

    PORT=8080

### Frontend .env

    NEXT_PUBLIC_API_URL=http://localhost:8080

# 🐳 Running the Project with Docker

    docker compose up --build
    docker compose down

# 🖥️ Running the Project Locally

Backend:

    go run cmd/app/main.go

Frontend:

    npm run dev

# 🏛️ Architecture Overview

Frontend \<-\> Backend

# 🧠 Architecture Decisions

- Separated frontend/backend architecture
- Docker-based development workflow
- REST API chosen for simplicity
- Environment-based configuration
- Modular file structure
- Optional CI/CD integration

# 📡 API Overview

Example:

    POST /upload
    GET /balance
    GET /issues

# 🚀 Deployment

Deploy using Docker Compose.
