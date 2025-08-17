import { ref } from 'vue'
import type { ApiRequestResult, ApiResponse } from '@/types'
import { API_BASE_URL, STORAGE_KEYS } from '@/utils/constants'
import { notifications } from '@/utils/notifications'

export function useApi() {
  const loading = ref(false)
  const error = ref<string | null>(null)

  const getAuthHeaders = (): HeadersInit => {
    const token = localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN)
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
    }

    if (token) {
      headers.Authorization = `Bearer ${token}`
    }

    return headers
  }

  const apiRequest = async <T = any>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiRequestResult<T>> => {
    loading.value = true
    error.value = null

    try {
      const url = endpoint.startsWith('http')
        ? endpoint
        : `${API_BASE_URL}${endpoint}`

      console.log(`Making API request to: ${url}`) // Debug log

      const response = await fetch(url, {
        ...options,
        headers: {
          ...getAuthHeaders(),
          ...options.headers,
        },
      })

      let data: ApiResponse<T>
      const contentType = response.headers.get('content-type')

      if (contentType && contentType.includes('application/json')) {
        data = await response.json()
      } else {
        // Handle non-JSON responses
        const text = await response.text()
        data = { error: text } as ApiResponse<T>
      }

      const result: ApiRequestResult<T> = {
        response,
        data,
        ok: response.ok,
        status: response.status,
      }

      if (!response.ok) {
        const errorMessage =
          data.message ||
          data.error ||
          `HTTP ${response.status}: ${response.statusText}`
        error.value = errorMessage

        // Handle 401 Unauthorized
        if (response.status === 401) {
          localStorage.removeItem(STORAGE_KEYS.ACCESS_TOKEN)
          localStorage.removeItem(STORAGE_KEYS.USER)
          window.location.href = '/login'
          return result
        }

        console.error('API Error:', {
          url,
          status: response.status,
          error: errorMessage,
          data,
        })

        throw new Error(errorMessage)
      }

      return result
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Network error'
      error.value = errorMessage

      console.error('API Request failed:', {
        endpoint,
        error: errorMessage,
        originalError: err,
      })

      // Don't show notification for auth errors to avoid spam
      if (!endpoint.includes('/auth/')) {
        notifications.error(errorMessage)
      }

      throw err
    } finally {
      loading.value = false
    }
  }

  const get = <T = any>(endpoint: string): Promise<ApiRequestResult<T>> => {
    return apiRequest<T>(endpoint, { method: 'GET' })
  }

  const post = <T = any>(
    endpoint: string,
    data?: any
  ): Promise<ApiRequestResult<T>> => {
    return apiRequest<T>(endpoint, {
      method: 'POST',
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  const put = <T = any>(
    endpoint: string,
    data?: any
  ): Promise<ApiRequestResult<T>> => {
    return apiRequest<T>(endpoint, {
      method: 'PUT',
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  const del = <T = any>(endpoint: string): Promise<ApiRequestResult<T>> => {
    return apiRequest<T>(endpoint, { method: 'DELETE' })
  }

  return {
    loading,
    error,
    apiRequest,
    get,
    post,
    put,
    delete: del,
  }
}
