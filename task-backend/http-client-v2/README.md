# Complete Testing Workflow - Task Management API

## Overview

Semua file HTTP client sudah diperbarui dengan variables dan auto-token extraction untuk testing yang lebih efisien.

## File Structure

```
http-client/
├── auth.http      - Authentication endpoints with auto-token saving
├── oauth.http     - OAuth endpoints with auto-token saving
├── task.http      - Task management endpoints using saved tokens
└── README.md      - Documentation
```

## Key Features

### 1. **Auto-Token Management**

- ✅ Tokens otomatis tersimpan setelah login
- ✅ Tidak perlu copy-paste manual
- ✅ Support untuk user, admin, dan OAuth tokens

### 2. **Variables untuk Flexibility**

- ✅ Base URL dapat diubah dengan mudah
- ✅ Email dan password tersimpan dalam variables
- ✅ Task IDs otomatis tersimpan dari response

### 3. **Structured Testing Flow**

- ✅ Urutan testing yang logis
- ✅ Error test cases included
- ✅ Multiple authentication methods

## Testing Workflow

### Step 1: Authentication Setup

1. **Run `auth.http`** - Start here first!
   ```http
   ### Login User (Auto-save token)
   # @name userLogin
   POST {{baseUrl}}/api/v1/auth/login
   ```
   - Jalankan user login → `@userToken` tersimpan otomatis
   - Jalankan admin login → `@adminToken` tersimpan otomatis

### Step 2: OAuth Setup (Optional)

2. **Run `oauth.http`** - If using Google OAuth
   ```http
   ### Google OAuth Callback
   # @name oauthCallback
   GET {{baseUrl}}/api/v1/auth/oauth/google/callback?code=...
   ```
   - OAuth callback → `@oauthToken` tersimpan otomatis

### Step 3: Task Management

3. **Run `task.http`** - Use saved tokens automatically
   ```http
   ### Create Task (uses auto-saved user token)
   POST {{baseUrl}}/api/v1/tasks
   Authorization: Bearer {{userToken}}
   ```
   - Semua endpoints langsung menggunakan saved tokens
   - Task IDs juga disimpan untuk update/delete

## Variables Reference

### Global Variables (Set in each file)

```http
@baseUrl = http://localhost:8080
@userEmail = john@example.com
@userPassword = password123
@adminEmail = admin@example.com
@adminPassword = admin123
```

### Auto-Generated Variables

```http
# From auth.http
@userToken = {{userLogin.response.body.data.access_token}}
@adminToken = {{adminLogin.response.body.data.access_token}}
@userRefreshToken = {{userLogin.response.body.data.refresh_token}}

# From oauth.http
@oauthToken = {{oauthCallback.response.body.data.access_token}}

# From task.http
@createdTaskId = {{createTask.response.body.data.id}}
@adminTaskId = {{adminCreateTask.response.body.data.id}}
```

## Usage Instructions

### 1. **First Time Setup**

```bash
# Install VS Code REST Client extension
# Or use any HTTP client that supports variables
```

### 2. **Environment Configuration**

```http
# Change these variables as needed:
@baseUrl = http://localhost:8080  # Change for different environments
@userEmail = john@example.com     # Change for different test users
```

### 3. **Testing Flow**

1. **Start with health check** in `auth.http`
2. **Register users** if needed
3. **Login users** - tokens auto-saved
4. **Run task operations** in `task.http` - uses saved tokens
5. **Test error cases** - included in each file

### 4. **Multi-Environment Testing**

```http
# Local
@baseUrl = http://localhost:8080

# Staging
@baseUrl = http://staging.yourdomain.com

# Production
@baseUrl = https://api.yourdomain.com
```

## Benefits

### ✅ **No Manual Token Copy-Paste**

- Tokens otomatis tersimpan setelah login
- Langsung digunakan di semua endpoint berikutnya

### ✅ **Environment Flexibility**

- Mudah switch antara local/staging/production
- Cukup ubah `@baseUrl` variable

### ✅ **Comprehensive Testing**

- User, admin, dan OAuth authentication
- CRUD operations dengan auto-saved IDs
- Error testing scenarios included

### ✅ **Developer Experience**

- Organized dan mudah dibaca
- Step-by-step workflow
- Clear variable naming

## Tips

1. **Always run auth.http first** untuk generate tokens
2. **Check response format** - pastikan path extraction sesuai dengan API response
3. **Use meaningful variable names** untuk clarity
4. **Save responses** untuk debugging jika diperlukan
5. **Test error cases** setelah happy path testing

## Troubleshooting

### Token Extraction Issues

```http
# Pastikan path response sesuai dengan actual API response
@userToken = {{userLogin.response.body.data.access_token}}

# Jika struktur response berbeda, adjust path:
@userToken = {{userLogin.response.body.access_token}}  # Direct access
@userToken = {{userLogin.response.body.token}}        # Different field name
```

### Variable Scope

- Variables tersimpan dalam scope file yang sama
- Untuk share across files, gunakan environment variables atau copy definitions

Dengan setup ini, testing workflow menjadi jauh lebih efisien dan tidak ada lagi manual copy-paste tokens! 🚀
