## Feature Toggle Management Application API

### Overview
This is the backend for the Feature Toggle Management application, built with Golang. It provides a RESTful API for managing feature toggles with JWT-based authentication and uses PostgreSQL as the database. The application provides JWT-based authentication and supports CRUD operations for managing feature toggles.

### Prerequisites
* Go (Golang) 1.16 or later
* Docker and Docker Compose (for running PostgreSQL database)
* Git (for version control)

### Setup
#### 1. Clone the Repository
```
git clone <repository-url>
cd <repository-directory>
```

#### 2. Configure Environment Variables
Create a `.env` file in the backend directory and set the following environment variables:
```
JWT_SECRET=your_jwt_secret
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
```

#### 3. Run PostgreSQL with Docker
Create a docker-compose.yml file to spin up a PostgreSQL instance:
```
version: '3'
services:
  db:
    image: postgres:13
    restart: always
    environment:
      POSTGRES_USER: your_db_user
      POSTGRES_PASSWORD: your_db_password
      POSTGRES_DB: your_db_name
    ports:
      - "5432:5432"
    volumes:
      - db_data:/var/lib/postgresql/data

volumes:
  db_data:
```
Run the following command to start the database:
```
docker-compose up -d
```

#### 4. Install Dependencies
```
go mod tidy
```

#### 5. Run the Backend Application
```
go run main.go
``` 

### Backend Libraries and Purpose
* `github.com/dgrijalva/jwt-go`: Library for creating and verifying JSON Web Tokens (JWTs), used for authentication.
* `github.com/go-sql-driver/mysql`: Go MySQL driver for interfacing with MySQL databases (you can replace this with pq for PostgreSQL).
* `github.com/gorilla/mux`: A powerful HTTP router and URL matcher for building Go web applications.
* `github.com/joho/godotenv`: Used for loading environment variables from a .env file.
* `database/sql`: Standard library for SQL database interactions in Go, providing a generic interface for SQL databases.

### Running the Backend 
Ensure the backend is running on `http://localhost:8080`.

### Troubleshooting
* **CORS Issues**: Ensure CORS headers are correctly configured on the backend.
* **Database Connection Errors**: Check your database connection details in the .env file and ensure the database is running.