# Task Management Frontend

Frontend aplikasi Task Management System yang dibangun dengan Vue 3, TypeScript, dan Tailwind CSS.

## 🚀 Features

- **Modern Tech Stack**: Vue 3 dengan Composition API, TypeScript, Tailwind CSS v4
- **Authentication**: Login/Register dengan JWT token management
- **Role-based Access**: User dan Admin dengan permission yang berbeda
- **Task Management**: CRUD operations untuk tasks dengan filtering dan pagination
- **Admin Panel**: Management users dan tasks untuk admin
- **Responsive Design**: Mobile-first design dengan Tailwind CSS
- **Type Safety**: Full TypeScript support untuk development yang aman

## 🛠️ Tech Stack

- **Frontend Framework**: Vue 3 dengan Composition API
- **Language**: TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS v4
- **State Management**: Pinia
- **Routing**: Vue Router 4
- **HTTP Client**: Fetch API
- **Icons**: Heroicons (optional)

## 📋 Prerequisites

- Node.js 18+ 
- npm atau yarn
- Backend API running di `http://localhost:8080`

## 🔧 Installation

1. **Install dependencies**
   ```bash
   npm install
   ```

2. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env sesuai konfigurasi backend Anda
   ```

3. **Start development server**
   ```bash
   npm run dev
   ```

4. **Build for production**
   ```bash
   npm run build
   ```

## 🏗️ Project Structure

```
src/
├── components/           # Vue components
│   ├── layouts/         # Layout components
│   ├── pages/           # Page components
│   ├── auth/            # Authentication components
│   └── common/          # Shared components
├── composables/         # Vue composables
├── router/              # Vue Router configuration
├── types/               # TypeScript type definitions
├── utils/               # Utility functions
└── style.css           # Global styles
```

## 🎯 Features Overview

### Authentication
- User registration dan login
- JWT token management
- Auto-redirect berdasarkan authentication status
- Logout functionality

### User Dashboard
- Overview tasks dengan statistics
- Quick actions untuk create task
- Recent tasks display

### Task Management
- Create, read, update, delete tasks
- Filter berdasarkan status
- Search functionality
- Pagination
- Status management (pending, in_progress, completed)

### Admin Panel
- User management dengan pagination
- All tasks overview
- System statistics
- Admin-only access control

### UI/UX Features
- Responsive design untuk semua device
- Loading states dan error handling
- Toast notifications
- Modal dialogs
- Form validation
- Consistent design system

## 🔐 Authentication Flow

1. **Register/Login**: User mendaftar atau login
2. **Token Storage**: JWT token disimpan di localStorage
3. **Auto-redirect**: Redirect ke dashboard setelah login
4. **Route Guards**: Protected routes memerlukan authentication
5. **Role-based Access**: Admin memiliki akses ke admin panel

## 📡 API Integration

Frontend berkomunikasi dengan backend Go melalui REST API:

- **Base URL**: `http://localhost:8080/api/v1`
- **Authentication**: Bearer token di header
- **Error Handling**: Centralized error handling dengan notifications
- **Auto-logout**: Otomatis logout jika token expired

## 🎨 Design System

### Colors
- **Primary**: Blue (600, 700)
- **Secondary**: Gray (200, 300)
- **Success**: Green (600, 700)
- **Warning**: Yellow (600, 700)
- **Danger**: Red (600, 700)

### Components
- **Buttons**: Primary, secondary, danger dengan consistent styling
- **Forms**: Input, select, textarea dengan validation states
- **Cards**: Consistent card design dengan shadow
- **Status Badges**: Color-coded status indicators

## 🧪 Development

### Available Scripts

```bash
# Development server
npm run dev

# Type checking
npm run type-check

# Build for production
npm run build

# Preview production build
npm run preview
```

### Code Style

- **TypeScript**: Strict mode enabled
- **Vue 3**: Composition API dengan `<script setup>`
- **Naming**: PascalCase untuk components, camelCase untuk functions
- **File Organization**: Feature-based organization

## 🚀 Deployment

1. **Build aplikasi**
   ```bash
   npm run build
   ```

2. **Deploy ke hosting**
   - Upload folder `dist/` ke web server
   - Configure web server untuk SPA routing
   - Set environment variables untuk production

### Environment Variables

```env
# Production
VITE_API_BASE_URL=https://api.yourdomain.com/api/v1
VITE_APP_NAME=Task Management System
VITE_APP_VERSION=1.0.0
```

## 🔧 Configuration

### Vite Configuration
- TypeScript support
- Path aliases (`@/` untuk `src/`)
- Proxy untuk development API calls
- Tailwind CSS integration

### Router Configuration
- Route guards untuk authentication
- Role-based route protection
- Lazy loading untuk code splitting

## 🤝 Contributing

1. Fork repository
2. Create feature branch
3. Commit changes dengan conventional commits
4. Push ke branch
5. Create Pull Request

## 📝 License

MIT License

## 🆘 Troubleshooting

### Common Issues

1. **API Connection Error**
   - Pastikan backend running di port 8080
   - Check CORS configuration di backend
   - Verify API base URL di .env

2. **Authentication Issues**
   - Clear localStorage dan cookies
   - Check JWT token expiration
   - Verify backend authentication endpoints

3. **Build Issues**
   - Clear node_modules dan reinstall
   - Check TypeScript errors
   - Verify all imports dan dependencies

### Development Tips

- Use Vue DevTools untuk debugging
- Check browser console untuk errors
- Use network tab untuk API debugging
- Verify responsive design di different screen sizes

## 📞 Support

Jika ada pertanyaan atau issue, silakan buat issue di repository ini.

---

**Happy Coding! 🚀**