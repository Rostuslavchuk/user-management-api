# User Management API

A robust RESTful API service for user management built with Go, featuring PostgreSQL integration, comprehensive CRUD operations, and secure database connectivity.

## 🎯 Overview

This project provides a complete backend solution for user data management with:
- **RESTful Architecture**: Clean API design following REST principles
- **Database Integration**: PostgreSQL with connection pooling and error handling
- **Environment Configuration**: Secure configuration management with .env files
- **CRUD Operations**: Complete Create, Read, Update, Delete functionality
- **Data Validation**: Input validation and sanitization
- **Error Handling**: Comprehensive error management and logging

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **HTTP Router**: Go's built-in `http.ServeMux`
- **Configuration**: Godotenv for environment variable management
- **Database Driver**: PostgreSQL driver (`lib/pq`)
- **Data Format**: JSON for request/response handling

## 🚀 Features

### API Endpoints

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|----------------|
| `GET` | `/users` | Retrieve all users | None |
| `GET` | `/users/{id}` | Retrieve user by ID | None |
| `POST` | `/users` | Create new user | None |
| `DELETE` | `/users/{id}` | Delete user by ID | None |

### Data Model
```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

### Database Schema
```sql
CREATE TABLE IF NOT EXISTS users_test(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(55) NOT NULL,
    age SMALLINT NOT NULL check (age >= 18) default 18
);
```

## 📋 Prerequisites

- Go 1.21 or higher installed
- PostgreSQL server running
- Database created for the application
- Git for cloning the repository

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up the database**
   ```sql
   CREATE DATABASE your_database_name;
   ```

4. **Configure environment variables**
   Create a `.env` file in the project root:
   ```env
   USER=your_postgres_user
   PASSWORD=your_postgres_password
   DBNAME=your_database_name
   PORT=5432
   ```

## 🚀 Usage

### Starting the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### API Usage Examples

#### 1. Create a New User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 25}'
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "age": 25
}
```

#### 2. Get All Users
```bash
curl http://localhost:8080/users
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 25
  },
  {
    "id": 2,
    "name": "Jane Smith", 
    "age": 30
  }
]
```

#### 3. Get User by ID
```bash
curl http://localhost:8080/users/1
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "age": 25
}
```

#### 4. Delete User
```bash
curl -X DELETE http://localhost:8080/users/1
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "age": 25
}
```

## 🏗️ Project Structure

```
api/
├── main.go              # Main application entry point
├── go.mod               # Go module file
├── go.sum               # Go module checksums
├── .env                 # Environment variables (create this)
└── README.md            # This file
```

## 🧩 Core Components

### Database Connection
```go
connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' port=%d sslmode=disable",
    user, password, dbname, port)
db, err := sql.Open("postgres", connStr)
```

### Route Handlers
- `getUsers()`: Retrieves all users from database
- `getByID()`: Retrieves a specific user by ID
- `createUser()`: Creates a new user record
- `deleteUser()`: Deletes a user by ID

### Error Handling
- Database connection errors
- Query execution errors
- JSON parsing errors
- HTTP status code management

## 🔧 Configuration

### Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `USER` | PostgreSQL username | Yes | - |
| `PASSWORD` | PostgreSQL password | Yes | - |
| `DBNAME` | Database name | Yes | - |
| `PORT` | PostgreSQL port | Yes | 5432 |

### Server Configuration
- **Port**: 8080 (hardcoded)
- **Database**: PostgreSQL with SSL disabled
- **Connection Pooling**: Default Go SQL settings

## 🎯 Performance Considerations

- **Connection Pooling**: Utilizes Go's built-in SQL connection pool
- **Query Optimization**: Simple, indexed queries
- **Memory Efficiency**: Streaming results for large datasets
- **Error Handling**: Fast fail on database connection issues

## 🧪 Testing

### Manual Testing
Use the provided curl examples or API testing tools like:
- Postman
- Insomnia
- curl commands

### Database Testing
```sql
-- Verify table creation
SELECT * FROM users_test;

-- Check data integrity
SELECT COUNT(*) FROM users_test;
```

## 🚧 Development Notes

### Extending the API
- Add authentication middleware
- Implement pagination for large datasets
- Add input validation middleware
- Implement rate limiting
- Add logging and monitoring

### Security Considerations
- Currently no authentication (add JWT or OAuth)
- No input validation (add validation middleware)
- No rate limiting (add middleware)
- SQL injection protection via parameterized queries

### Known Limitations
- No authentication/authorization
- Limited error responses
- No pagination
- No input validation
- No API documentation (Swagger)

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📞 Contact

For questions or support regarding this API, please open an issue in the repository.

---

**Built with ❤️ using Go and PostgreSQL**
