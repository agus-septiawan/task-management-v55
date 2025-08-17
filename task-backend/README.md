# Task Management API

Backend API untuk sistem manajemen task dengan fitur authentication, authorization, dan CRUD operations.

## 🚀 Features

- **User Authentication**: Register, Login, Logout dengan JWT
- **Authorization**: Role-based access control (User & Admin)
- **User Management**: CRUD operations untuk users
- **Security**: Password hashing, JWT tokens, CORS
- **Validation**: Input validation dengan error messages
- **Documentation**: OpenAPI/Swagger specification
- **Testing**: Unit dan integration tests

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Router**: Gorilla Mux
- **Database**: MySQL
- **Authentication**: JWT + bcrypt
- **Validation**: go-playground/validator
- **Configuration**: Viper
- **Testing**: Go native testing

## 📋 Prerequisites

- Go 1.24 atau lebih baru
- MySQL 8.0+
- Git

## 🔧 Installation

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd task-management-backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup database**
   ```bash
   # Menggunakan Docker (recommended)
   make docker-up
   
   # Atau install MySQL manual dan buat database
   CREATE DATABASE task_manager;
   ```

4. **Run database migrations**
   ```bash
   # Manual migration
   mysql -u root -p task_manager < migrations/20250816103214_create_users_table.up.sql
   mysql -u root -p task_manager < migrations/20250816103219_create_tasks_table.up.sql
   ```

5. **Setup configuration**
   ```bash
   cp .env.example .env
   # Edit .env sesuai dengan konfigurasi database Anda
   ```

## 🏃‍♂️ Running the Application

### Development Mode

```bash
# Dengan live reload (requires air)
make dev

# Atau manual
make run
```

### Production Mode

```bash
make build
./bin/server
```

## 🧪 Testing

```bash
# Run semua tests
make test

# Run unit tests saja
make test-unit

# Run integration tests saja
make test-int
```

## 📡 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/logout` - User logout
- `GET /api/v1/auth/me` - Get user profile

### Admin (Admin only)
- `GET /api/v1/admin/users` - Get all users

### Health Check
- `GET /health` - Server health status

## 📖 API Documentation

API documentation tersedia dalam format OpenAPI/Swagger di `docs/swagger.yaml`.

Untuk melihat dokumentasi interaktif, buka file tersebut di Swagger Editor atau tools serupa.

## 🔐 Authentication Flow

1. **Register**: User mendaftar dengan email, name, dan password
2. **Login**: User login dengan email dan password, mendapat access token
3. **Access**: Gunakan access token di header `Authorization: Bearer <token>`
4. **Refresh**: Gunakan refresh token (dari cookie) untuk mendapat access token baru

## 🧪 Testing dengan HTTP Client

Gunakan file di folder `http-client/` untuk testing API:

```bash
# Install VS Code REST Client extension
# Buka file http-client/auth.http
# Klik "Send Request" untuk testing endpoints
```

## 🏗️ Project Structure

```
task-management-backend/
├── cmd/server/           # Application entry point
├── internal/             # Private application code
│   ├── config/          # Configuration management
│   ├── database/        # Database connection
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   ├── router/          # Route definitions
│   └── service/         # Business logic
├── pkg/                 # Public library code
│   ├── jwt/            # JWT utilities
│   ├── response/       # Response utilities
│   └── validator/      # Validation utilities
├── migrations/         # Database migrations
├── configs/           # Configuration files
├── tests/             # Test files
├── http-client/       # HTTP client testing files
└── docs/              # Documentation
```

## 🔧 Configuration

Konfigurasi aplikasi menggunakan file `configs/config.yaml`:

```yaml
server:
  port: "8080"
  host: "localhost"

database:
  host: "localhost"
  port: "3306"
  name: "task_manager"
  user: "root"
  password: "password"

jwt:
  access_secret: "your-secret-key"
  refresh_secret: "your-refresh-secret"
  access_expire: "30m"
  refresh_expire: "168h"
```

## 🚀 Deployment

1. **Build aplikasi**
   ```bash
   make build
   ```

2. **Setup production database**
   - Buat database MySQL
   - Jalankan migrations
   - Update konfigurasi

3. **Run aplikasi**
   ```bash
   ./bin/server
   ```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch
3. Commit changes
4. Push ke branch
5. Buat Pull Request

## 📝 License

MIT License - lihat file LICENSE untuk detail.

## 🆘 Troubleshooting

### Database Connection Error
- Pastikan MySQL running
- Cek konfigurasi database di `configs/config.yaml`
- Pastikan database `task_manager` sudah dibuat

### JWT Token Error
- Pastikan JWT secrets dikonfigurasi
- Cek format Authorization header: `Bearer <token>`

### Migration Error
- Pastikan database sudah dibuat
- Jalankan migrations secara berurutan
- Cek permission database user

## 📞 Support

Jika ada pertanyaan atau issue, silakan buat issue di repository ini.