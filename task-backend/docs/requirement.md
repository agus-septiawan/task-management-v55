# Task Management API - Requirements Document

## 📋 Project Overview

**Project Name**: Task Management API  
**Purpose**: Backend web API untuk pembelajaran dengan fitur lengkap  
**Target**: Menunjang pembelajaran backend development dengan Go  
**Complexity Level**: Beginner to Intermediate

## 🎯 Project Goals

- Implementasi backend API modern dengan Go tanpa framework
- Pembelajaran best practices Go development
- Implementasi security features (JWT + OAuth)
- Database operations tanpa ORM
- Testing dan documentation yang baik

## 🏗️ Technical Requirements

### Backend Technology Stack

| Component           | Technology                        | Purpose                       |
| ------------------- | --------------------------------- | ----------------------------- |
| **Language**        | Go (Golang) latest version        | Main programming language     |
| **HTTP Router**     | Gorilla Mux                       | HTTP request routing          |
| **Database**        | MySQL                             | Data persistence              |
| **Database Access** | Native SQL (no ORM)               | Direct SQL operations         |
| **Validation**      | go-playground/validator           | Request validation            |
| **Configuration**   | Viper                             | Configuration management      |
| **Migration**       | golang-migrate                    | Database schema management    |
| **Authentication**  | JWT + OAuth                       | Security implementation       |
| **CORS**            | Gorilla CORS                      | Cross-origin resource sharing |
| **Logging**         | Standard log + structured logging | Application logging           |
| **Testing**         | Go native testing                 | Unit and integration tests    |
| **Documentation**   | OpenAPI/Swagger                   | API documentation             |

### Development Principles

- ✅ **No Framework**: Menggunakan standard library Go + packages minimal
- ✅ **Best Practices**: Mengikuti convention dan best practices Go terbaru
- ✅ **Clean Architecture**: Separation of concerns yang jelas
- ✅ **Security First**: Implementasi security yang proper
- ✅ **Testing**: Unit testing untuk semua komponen kritis
- ✅ **Documentation**: Kode terdokumentasi dalam bahasa Indonesia

## 👥 User Management Requirements

### User Roles

1. **User** (Regular user)

   - Dapat mengelola task milik sendiri
   - Akses terbatas sesuai ownership

2. **Admin** (Administrator)
   - Dapat melihat semua user dan task
   - Akses penuh untuk management

### User Features

- ✅ User registration dengan validasi
- ✅ User login dengan email/password
- ✅ Profile management
- ✅ Role-based access control

## 🔐 Authentication & Authorization Requirements

### JWT Implementation

- **Access Token**:

  - Short-lived (15-30 minutes)
  - Berisi user information dan roles
  - Digunakan untuk API authentication

- **Refresh Token**:
  - Long-lived (7-30 days)
  - Disimpan dalam HTTP-only secure cookie
  - Digunakan untuk mendapatkan access token baru

### OAuth Integration

- **Google OAuth 2.0**: Login dengan akun Google
- **GitHub OAuth** (optional): Login dengan akun GitHub
- **Security**: Proper state parameter dan PKCE implementation

### Security Features

- ✅ Password hashing (bcrypt)
- ✅ JWT token validation
- ✅ Secure cookie for refresh token
- ✅ CORS configuration
- ✅ Rate limiting (optional)
- ✅ Input validation dan sanitization

## 📝 Task Management Requirements

### Core CRUD Operations

1. **Create Task**

   - Title (required, max 255 chars)
   - Description (optional)
   - Status (pending, in_progress, completed)
   - Auto-assign user_id dari JWT token

2. **Read Tasks**

   - List tasks dengan pagination
   - Filter by status
   - Search by title/description
   - User: hanya task sendiri
   - Admin: semua task

3. **Update Task**

   - Update title, description, status
   - Authorization check berdasarkan ownership
   - Admin bisa update semua task

4. **Delete Task**
   - Soft delete atau hard delete
   - Authorization check berdasarkan ownership
   - Admin bisa delete semua task

### Task Features

- ✅ Pagination support
- ✅ Search and filter functionality
- ✅ Status management
- ✅ Timestamp tracking (created_at, updated_at)
- ✅ Authorization per task

## 🗄️ Database Requirements

### Database Schema

#### Users Table

```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255), -- NULL untuk OAuth users
    role ENUM('user', 'admin') DEFAULT 'user',
    oauth_provider VARCHAR(50), -- google, github, null
    oauth_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_email (email),
    INDEX idx_oauth (oauth_provider, oauth_id)
);
```

#### Tasks Table

```sql
CREATE TABLE tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);
```

#### Refresh Tokens Table (optional, untuk token blacklisting)

```sql
CREATE TABLE refresh_tokens (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_token_hash (token_hash),
    INDEX idx_expires_at (expires_at)
);
```

## 🏛️ Project Structure Requirements

```
task-manager-api/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/                       # Private application code
│   ├── handler/                    # HTTP handlers (controllers)
│   │   ├── auth_handler.go
│   │   ├── task_handler.go
│   │   └── admin_handler.go
│   ├── service/                    # Business logic
│   │   ├── auth_service.go
│   │   ├── task_service.go
│   │   └── user_service.go
│   ├── repository/                 # Data access layer
│   │   ├── user_repository.go
│   │   └── task_repository.go
│   ├── middleware/                 # HTTP middleware
│   │   ├── auth_middleware.go
│   │   ├── cors_middleware.go
│   │   └── logging_middleware.go
│   ├── model/                      # Data models
│   │   ├── user.go
│   │   ├── task.go
│   │   └── auth.go
│   └── config/                     # Configuration
│       └── config.go
├── pkg/                            # Public library code
│   ├── jwt/                        # JWT utilities
│   ├── oauth/                      # OAuth utilities
│   ├── validator/                  # Validation utilities
│   └── response/                   # Response utilities
├── migrations/                     # Database migrations
│   ├── 001_create_users_table.up.sql
│   ├── 001_create_users_table.down.sql
│   ├── 002_create_tasks_table.up.sql
│   └── 002_create_tasks_table.down.sql
├── configs/                        # Configuration files
│   ├── config.yaml
│   └── config.example.yaml
├── docs/                          # Documentation
│   └── swagger.yaml               # OpenAPI specification
├── tests/                         # Test files
│   ├── integration/
│   └── unit/
├── scripts/                       # Build and deployment scripts
├── .env.example                   # Environment variables example
├── .gitignore                     # Git ignore file
├── go.mod                         # Go modules
├── go.sum                         # Go modules checksum
├── Makefile                       # Build automation
└── README.md                      # Project documentation
```

## 📡 API Endpoints Requirements

### Authentication Endpoints

- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/auth/me` - Get current user profile

### OAuth Endpoints

- `GET /api/v1/auth/oauth/google` - Redirect to Google OAuth
- `GET /api/v1/auth/oauth/google/callback` - Google OAuth callback

### Task Endpoints

- `GET /api/v1/tasks` - Get tasks (with pagination, search, filter)
- `POST /api/v1/tasks` - Create new task
- `GET /api/v1/tasks/{id}` - Get task by ID
- `PUT /api/v1/tasks/{id}` - Update task by ID
- `DELETE /api/v1/tasks/{id}` - Delete task by ID

### Admin Endpoints

- `GET /api/v1/admin/users` - Get all users (admin only)

## ✅ Validation Requirements

### Request Validation

- **Email format validation**
- **Password strength validation** (minimal 6 karakter)
- **Required field validation**
- **String length validation**
- **Enum validation** untuk status dan role
- **Custom validation** untuk business rules

### Response Validation

- **Consistent error format**
- **Proper HTTP status codes**
- **Detailed validation error messages**
- **Input sanitization**

## 🧪 Testing Requirements

### Unit Testing

- ✅ Service layer testing
- ✅ Repository layer testing
- ✅ Utility function testing
- ✅ JWT token testing
- ✅ Validation testing

### Integration Testing

- ✅ API endpoint testing
- ✅ Database integration testing
- ✅ Authentication flow testing
- ✅ Authorization testing

### Testing Coverage

- **Target**: Minimal 80% code coverage
- **Focus**: Critical business logic
- **Tools**: Go native testing + testify/assert

## 📚 Documentation Requirements

### Code Documentation

- ✅ **Bahasa Indonesia** untuk komentar
- ✅ Package documentation
- ✅ Function documentation
- ✅ Complex logic explanation

### API Documentation

- ✅ OpenAPI/Swagger specification
- ✅ Request/response examples
- ✅ Error code documentation
- ✅ Authentication guide

### Project Documentation

- ✅ README.md with setup instructions
- ✅ Environment configuration guide
- ✅ Deployment guide
- ✅ Development workflow

## 🔧 Development Requirements

### Environment Configuration

```yaml
# config.yaml example
server:
  port: 8080
  host: localhost

database:
  host: localhost
  port: 3306
  name: task_manager
  user: username
  password: password

jwt:
  access_secret: your-access-secret
  refresh_secret: your-refresh-secret
  access_expire: 30m
  refresh_expire: 168h

oauth:
  google:
    client_id: your-google-client-id
    client_secret: your-google-client-secret
    redirect_url: http://localhost:8080/api/v1/auth/oauth/google/callback
```

### Development Tools

- ✅ **Makefile** untuk automation
- ✅ **Docker** untuk database (optional)
- ✅ **Air** untuk live reload (optional)
- ✅ **Golangci-lint** untuk code quality
- ✅ **Git hooks** untuk quality assurance

## 🎯 Learning Objectives

Setelah menyelesaikan project ini, Anda akan memahami:

1. **Go Best Practices**: Project structure, naming convention, error handling
2. **HTTP Server Development**: Routing, middleware, request handling
3. **Database Operations**: Raw SQL, connection management, transactions
4. **Authentication & Authorization**: JWT implementation, OAuth flow
5. **Security**: Password hashing, token management, input validation
6. **Testing**: Unit testing, integration testing, test-driven development
7. **API Design**: RESTful principles, OpenAPI specification
8. **Configuration Management**: Environment variables, config files
9. **Logging**: Structured logging, error tracking
10. **Documentation**: Code documentation, API documentation

## 📈 Success Criteria

Project dianggap berhasil jika:

- ✅ Semua endpoint API berfungsi sesuai specification
- ✅ Authentication dan authorization bekerja dengan benar
- ✅ Database operations berjalan tanpa error
- ✅ Unit tests passed dengan coverage > 80%
- ✅ Code mengikuti Go best practices
- ✅ API documentation lengkap dan akurat
- ✅ Error handling yang proper
- ✅ Security implementation yang aman

## 🚀 Implementation Phases

### Phase 1: Foundation

- Project setup dan structure
- Database schema dan migration
- Basic configuration

### Phase 2: Core Features

- User registration dan login
- Basic CRUD operations untuk tasks
- Input validation

### Phase 3: Security

- JWT implementation
- Authentication middleware
- Authorization logic

### Phase 4: Advanced Features

- OAuth integration
- Admin features
- Search dan filtering

### Phase 5: Quality Assurance

- Unit testing
- Integration testing
- Documentation

### Phase 6: Deployment

- Production configuration
- Deployment guide
- Performance optimization

---

**Note**: Requirements ini dapat disesuaikan berdasarkan kebutuhan pembelajaran dan waktu yang tersedia. Focus utama adalah memahami konsep-konsep backend development dengan Go.
