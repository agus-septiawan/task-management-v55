// API Response Types
export interface ApiResponse<T = any> {
  data?: T;
  error?: string;
  message?: string;
  access_token?: string;
  user?: User;
}

export interface PagingInfo {
  page: number;
  limit: number;
  total: number;
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