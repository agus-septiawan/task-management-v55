# HTTP Client Testing

Folder ini berisi file-file untuk testing API menggunakan HTTP client seperti VS Code REST Client extension atau tools seperti Postman.

## Setup

1. Install VS Code REST Client extension untuk menjalankan file `.http`
2. Pastikan server berjalan di `http://localhost:8080`
3. Pastikan database MySQL sudah running dan terkonfigurasi

## Files

- `auth.http` - Testing endpoints untuk authentication dan user management
- `tasks.http` - Testing endpoints untuk task management
- `oauth.http` - Testing endpoints untuk OAuth authentication

## Usage

### 1. Testing Flow untuk User Registration & Login

1. **Health Check** - Pastikan server berjalan
2. **Register User** - Daftarkan user baru
3. **Login User** - Login dan dapatkan access token
4. **Get Profile** - Test endpoint yang memerlukan authentication

### 2. Testing Flow untuk Admin

1. **Register Admin** - Daftarkan admin (manual update role di database)
2. **Login Admin** - Login sebagai admin
3. **Get All Users** - Test admin endpoint

### 3. Testing Flow untuk Tasks

1. **Create Task** - Buat task baru
2. **Get Tasks** - Ambil semua tasks dengan berbagai filter
3. **Get Task by ID** - Ambil task spesifik
4. **Update Task** - Update task
5. **Delete Task** - Hapus task

### 4. Testing Flow untuk OAuth

1. **Google OAuth** - Redirect ke Google OAuth
2. **OAuth Callback** - Handle callback dari Google
3. **OAuth User Profile** - Test profile OAuth user

### 5. Testing Error Cases

- Invalid JSON format
- Missing required fields
- Invalid email format
- Password too short
- Wrong credentials
- Unauthorized access
- Invalid tokens
- Invalid task data
- Access control violations

## Notes

- Ganti `YOUR_ACCESS_TOKEN_HERE` dengan access token yang didapat dari response login
- Ganti `YOUR_ADMIN_ACCESS_TOKEN_HERE` dengan access token admin
- Untuk membuat admin user, setelah register, update role di database:
  ```sql
  UPDATE users SET role = 'admin' WHERE email = 'admin@example.com';
  ```

## OAuth Setup

Untuk testing OAuth:

1. **Setup Google OAuth**:
   - Buat project di Google Cloud Console
   - Enable Google+ API
   - Buat OAuth 2.0 credentials
   - Set redirect URI: `http://localhost:8080/api/v1/auth/oauth/google/callback`
   - Update `configs/config.yaml` dengan client ID dan secret

2. **Testing OAuth Flow**:
   - Buka URL OAuth di browser
   - Login dengan akun Google
   - Copy callback URL dengan code dan state
   - Test callback endpoint dengan parameter tersebut

## Database Setup

Pastikan database MySQL sudah running dan jalankan migration:

```bash
# Buat database
CREATE DATABASE task_manager;

# Jalankan migration files
mysql -u root -p task_manager < migrations/20250816103214_create_users_table.up.sql
```