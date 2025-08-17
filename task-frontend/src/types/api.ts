import type { BaseEntity } from './index';

// User Types
export interface User extends BaseEntity {
  email: string;
  name: string;
  role: 'user' | 'admin';
  oauth_provider?: string;
  oauth_id?: string;
}

export interface UserRegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface UserLoginRequest {
  email: string;
  password: string;
}

export interface AuthResponse {
  access_token: string;
  user: User;
}

// Task Types
export interface Task extends BaseEntity {
  user_id: number;
  title: string;
  description?: string;
  status: TaskStatus;
}

export type TaskStatus = 'pending' | 'in_progress' | 'completed';

export interface TaskCreateRequest {
  title: string;
  description?: string;
  status?: TaskStatus;
}

export interface TaskUpdateRequest {
  title?: string;
  description?: string;
  status?: TaskStatus;
}

export interface TasksResponse {
  tasks: Task[];
  total: number;
  page: number;
  limit: number;
}

export interface UsersResponse {
  users: User[];
  total: number;
  page: number;
  limit: number;
}

// Error Types
export interface ValidationError {
  field: string;
  message: string;
}

export interface ValidationErrorResponse {
  error: string;
  details: ValidationError[];
}