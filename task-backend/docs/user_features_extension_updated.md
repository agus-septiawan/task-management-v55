# Extended User Features - Requirements (Updated)

## 📋 Overview

- Profile photo management (dengan folder `temp/`, `profiles/`, `thumbnails/`) misalnya
- Forgot/Reset password menggunakan SMTP
- Tracking IP address dan User-Agent untuk aktivitas sensitif
- **OAuth Photo Support** → jika user login via Google OAuth, foto profil default diambil dari Google

---

## 👤 Profile Photo Management

### Folder Structure

mungkin contohnya seperti ini(hanya contoh)

```
uploads/
├── temp/         # File sementara saat upload
├── profiles/     # File hasil resize (ukuran penuh)
└── thumbnails/   # File hasil thumbnail (preview kecil)
```

### Flow Upload (Email/Password User)

1. User upload file → backend simpan di `temp/`
2. Validasi file (ukuran, format, MIME type)
3. Proses resize/compress → simpan ke `profiles/`
4. Generate thumbnail (150x150) → simpan ke `thumbnails/`
5. Hapus file di `temp/`

### Flow OAuth (Google Login)

- Jika user login via **Google OAuth**, backend akan menyimpan `oauth_photo_url` dari Google.
- Frontend secara default menampilkan `oauth_photo_url`.
- User tetap bisa mengganti photo profil manual

---

## 🔐 Forgot & Reset Password

### Flow

1. User input email → backend generate token
2. Token + IP + User-Agent disimpan di tabel `password_reset_tokens`
3. Kirim email via **SMTP server**
4. User klik link reset → frontend tampilkan form reset
5. Backend validasi token (expired, used, IP/userAgent tracking)
6. User set password baru → token ditandai used

### Email Example

```
Subject: Reset Your Password

Hi John,
Someone requested a password reset for your account.
If this was you, click the link below:

https://yourapp.com/reset-password?token=abcdef123456

This link will expire in 15 minutes.
```

### Database Table

```sql
CREATE TABLE password_reset_tokens (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

---

## 📧 SMTP Integration

- Menggunakan SMTP
- Jangan bocorkan apakah email terdaftar atau tidak (respon selalu sama)

## 🛡️ Security Considerations

- Token reset password expire cepat (misalnya 15 menit)
- Token hanya bisa dipakai sekali
- Rate limiting forgot password (misalnya max 3x/jam)
- Upload file dibatasi maksimal tertentu yang direkomendasikan
- Validasi MIME type dan extension
- Cleanup otomatis file di `temp/`
- Mendukung load konfigurasi dari sebuh file, jadi smtp dll itu setupnya/loadnya dari folder configs/
- **Untuk OAuth user**: jika tidak pernah set password manual, forgot password tidak tersedia
