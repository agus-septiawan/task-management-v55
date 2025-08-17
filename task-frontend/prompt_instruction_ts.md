# Vue 3 Project Structure - TypeScript Version

## Technology Stack
- **Frontend**: Vue 3 with Composition API
- **Language**: TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS v4 (installed via npm)
- **State Management**: Pinia
- **Routing**: Vue Router 4
- **HTTP Client**: Fetch API
- **Notifications**: SweetAlert2 (optional)
- **Icons**: Heroicons or Lucide (optional)
- **Fonts**: Inter or system fonts

## Project Structure

```
project-name/
├── .env                          # Environment variables
├── .gitignore                    # Git ignore rules
├── .prettierrc                   # Prettier configuration
├── .vscode/
│   └── extensions.json           # VS Code recommended extensions
├── README.md                     # Project documentation
├── index.html                    # Main HTML template
├── package.json                  # Dependencies and scripts
├── tailwind.config.js           # Tailwind CSS v4 configuration
├── tsconfig.json                 # TypeScript configuration
├── tsconfig.node.json           # Node-specific TypeScript config
├── vite.config.ts               # Vite configuration
├── public/                       # Static assets
└── src/
    ├── App.vue                  # Root component
    ├── main.ts                  # Application entry point
    ├── style.css                # Tailwind CSS imports
    ├── types/                   # TypeScript type definitions
    │   ├── index.ts             # Main type exports
    │   ├── api.ts               # API-related types
    │   └── [feature].ts         # Feature-specific types
    ├── components/              # Vue components
    │   ├── layouts/             # Layout components
    │   ├── auth/                # Authentication components
    │   ├── [feature]/           # Feature-specific components
    │   └── common/              # Shared components
    ├── composables/             # Vue composables
    │   ├── useApi.ts           # API request composable
    │   ├── useAuth.ts          # Authentication composable
    │   └── use[Feature].ts     # Feature-specific composables
    ├── router/                  # Vue Router configuration
    │   └── index.ts            # Router setup and guards
    ├── services/               # API service layer
    │   ├── apiService.ts       # Base API service
    │   └── [feature]Service.ts # Feature-specific services
    ├── stores/                 # Pinia stores (if using Pinia)
    │   └── [feature]Store.ts   # Feature-specific stores
    └── utils/                  # Utility functions
        ├── constants.ts        # Application constants
        ├── helpers.ts          # Helper functions
        └── notifications.ts    # Notification utilities
```

## Configuration Files

### Package.json Dependencies
```json
{
  "name": "vue-ts-project",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vue-tsc && vite build",
    "preview": "vite preview",
    "type-check": "vue-tsc --noEmit"
  },
  "dependencies": {
    "@vueuse/core": "^13.4.0",
    "pinia": "^3.0.3",
    "vue": "^3.5.13",
    "vue-router": "^4.5.1"
  },
  "devDependencies": {
    "@tailwindcss/vite": "^4.0.0-alpha.26",
    "@types/node": "^22.10.5",
    "@vitejs/plugin-vue": "^5.2.3",
    "tailwindcss": "^4.0.0-alpha.26",
    "typescript": "^5.7.3",
    "vite": "^6.3.5",
    "vue-tsc": "^2.2.0"
  },
  "optionalDependencies": {
    "sweetalert2": "^11.22.1",
    "@heroicons/vue": "^2.0.18",
    "lucide-vue-next": "^0.400.0"
  }
}
```

### TypeScript Configuration (tsconfig.json)
```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "preserve",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": ["src/**/*.ts", "src/**/*.tsx", "src/**/*.vue"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

### Vite Configuration (vite.config.ts)
```typescript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue(), tailwindcss()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
})
```

### Tailwind CSS v4 Setup

#### tailwind.config.js
```javascript
import { defineConfig } from '@tailwindcss/vite'

export default {
  content: ['./index.html', './src/**/*.{vue,ts,tsx}']
}
```

#### src/style.css
```css
@import "tailwindcss";

/* Custom styles here */
```

### HTML Template (index.html)
```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Your App Title</title>
</head>
<body>
<div id="app"></div>
<script type="module" src="/src/main.ts"></script>
</body>
</html>
```

### VS Code Extensions (.vscode/extensions.json)
```json
{
  "recommendations": ["Vue.volar", "Vue.vscode-typescript-vue-plugin"]
}
```

## TypeScript Type System

### Core Types Structure (src/types/index.ts)
```typescript
// API Response Types
export interface ApiResponse<T = any> {
  data: T;
  errors?: string;
  paging?: PagingInfo;
}

export interface PagingInfo {
  page: number;
  total_page: number;
  total_item: number;
}

// Generic CRUD Types
export interface BaseEntity {
  id: number;
  created_at?: string;
  updated_at?: string;
}

// Request/Response Types
export interface ApiRequestResult<T = any> {
  response: Response;
  data: ApiResponse<T>;
  ok: boolean;
  status: number;
}

// Form Types
export interface BaseForm {
  [key: string]: any;
}

// Component Props Types
export interface BaseProps {
  id?: string | number;
}
```

### Feature-Specific Types (src/types/[feature].ts)
```typescript
// Example: User types
export interface User extends BaseEntity {
  username: string;
  name: string;
  email?: string;
}

export interface UserCreateRequest {
  username: string;
  name: string;
  password: string;
}

export interface UserUpdateRequest {
  name?: string;
  password?: string;
}
```

## Core Architecture with TypeScript

### 1. Main Entry Point (src/main.ts)
```typescript
import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import "./style.css";

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.mount("#app");
```

### 2. Composables with TypeScript
```typescript
// Example: useApi composable
import type { ApiRequestResult } from '@/types';

export function useApi() {
  const apiRequest = async <T = any>(
    url: string, 
    options: RequestInit = {}
  ): Promise<ApiRequestResult<T>> => {
    // Implementation
  };
  
  return { apiRequest };
}
```

### 3. Service Layer with TypeScript
```typescript
// Example: Base service
import { useApi } from '@/composables/useApi';
import type { ApiRequestResult, BaseEntity } from '@/types';

export abstract class BaseService<T extends BaseEntity> {
  protected abstract endpoint: string;
  protected api = useApi();

  async getAll(): Promise<ApiRequestResult<T[]>> {
    return await this.api.apiRequest<T[]>(this.endpoint);
  }

  async getById(id: number): Promise<ApiRequestResult<T>> {
    return await this.api.apiRequest<T>(`${this.endpoint}/${id}`);
  }

  async create(data: Omit<T, 'id'>): Promise<ApiRequestResult<T>> {
    return await this.api.apiRequest<T>(this.endpoint, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async update(id: number, data: Partial<T>): Promise<ApiRequestResult<T>> {
    return await this.api.apiRequest<T>(`${this.endpoint}/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async delete(id: number): Promise<ApiRequestResult<void>> {
    return await this.api.apiRequest<void>(`${this.endpoint}/${id}`, {
      method: 'DELETE',
    });
  }
}
```

### 4. Router with TypeScript
```typescript
import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import { useAuth } from "@/composables/useAuth";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: () => import("@/components/layouts/PublicLayout.vue"),
    children: [
      {
        path: "login",
        name: "Login",
        component: () => import("@/components/auth/Login.vue"),
        meta: { requiresGuest: true },
      },
    ],
  },
  {
    path: "/dashboard",
    component: () => import("@/components/layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true },
    children: [
      // Dashboard routes
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guards with proper typing
router.beforeEach((to, from, next) => {
  const { isAuthenticated } = useAuth();
  
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    next("/login");
  } else if (to.meta.requiresGuest && isAuthenticated.value) {
    next("/dashboard");
  } else {
    next();
  }
});

export default router;
```

## Component Patterns with TypeScript

### 1. Component Props Definition
```vue
<script setup lang="ts">
interface Props {
  id: string | number;
  title?: string;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Default Title',
  disabled: false,
});

interface Emits {
  (e: 'update', value: string): void;
  (e: 'delete', id: number): void;
}

const emit = defineEmits<Emits>();
</script>
```

### 2. Reactive Data with Types
```vue
<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { User, UserForm } from '@/types';

const user = ref<User | null>(null);
const form = reactive<UserForm>({
  name: '',
  email: '',
});

const users = ref<User[]>([]);
</script>
```

## Development Guidelines

### 1. Type Safety Best Practices
- Use strict TypeScript configuration
- Define interfaces for all data structures
- Use generic types for reusable components
- Implement proper error handling with typed responses

### 2. File Naming Conventions
- Components: PascalCase (UserProfile.vue)
- Types: PascalCase interfaces (User, ApiResponse)
- Composables: camelCase with 'use' prefix (useAuth.ts)
- Services: PascalCase classes (UserService.ts)

### 3. Import/Export Patterns
```typescript
// Centralized type exports
export type { User, UserForm } from './user';
export type { ApiResponse, PagingInfo } from './api';

// Re-export pattern
export * from './user';
export * from './api';
```

### 4. Environment Variables with Types
```typescript
// src/types/env.d.ts
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string;
  readonly VITE_APP_NAME: string;
  readonly VITE_APP_VERSION: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
```

### 5. Optional Features Setup

#### SweetAlert2 (Optional)
```typescript
// src/utils/notifications.ts
import Swal from 'sweetalert2';

export const showSuccess = async (message: string): Promise<void> => {
  await Swal.fire({
    icon: 'success',
    title: 'Success',
    text: message,
    timer: 3000,
    showConfirmButton: false,
  });
};
```

#### Icons (Optional)
```typescript
// Using Heroicons
import { UserIcon } from '@heroicons/vue/24/outline'

// Using Lucide
import { User } from 'lucide-vue-next'
```

## Build and Deployment

### Scripts with Type Checking
```json
{
  "scripts": {
    "dev": "vite",
    "build": "vue-tsc && vite build",
    "preview": "vite preview",
    "type-check": "vue-tsc --noEmit",
    "lint": "eslint . --ext .vue,.ts,.tsx --fix --ignore-path .gitignore",
    "format": "prettier --write src/"
  }
}
```

### Environment Configuration
```env
# .env.local
VITE_API_BASE_URL=http://localhost:3000/api
VITE_APP_NAME=Your App Name
VITE_APP_VERSION=1.0.0
```

## Optional Features
- **Notifications**: SweetAlert2 for user feedback
- **Icons**: Heroicons or Lucide for UI icons
- **Form Validation**: VeeValidate or custom validation
- **Testing**: Vitest + Vue Test Utils
- **Linting**: ESLint + Prettier
- **State Management**: Advanced Pinia patterns
- **Internationalization**: Vue I18n

This TypeScript structure ensures type safety, better IDE support, and maintainable code while following Vue 3 and modern development best practices.
