import { ref, computed } from 'vue'
import type {
  User,
  UserLoginRequest,
  UserRegisterRequest,
  AuthResponse,
} from '@/types/api'
import { useApi } from './useApi'
import { STORAGE_KEYS } from '@/utils/constants'
import { notifications } from '@/utils/notifications'

const user = ref<User | null>(null)
const isAuthenticated = computed(() => !!user.value)
const isAdmin = computed(() => user.value?.role === 'admin')

export function useAuth() {
  const { post, get } = useApi()

  // Initialize user from localStorage
  const initAuth = async (): Promise<boolean> => {
    const storedUser = localStorage.getItem(STORAGE_KEYS.USER)
    const token = localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN)

    if (token) {
      if (storedUser) {
        try {
          user.value = JSON.parse(storedUser)
          return true
        } catch (error) {
          console.error('Error parsing stored user:', error)
          // If user data is corrupted, fetch fresh profile
          return (await getProfile()) !== null
        }
      } else {
        // If we have token but no user data, fetch profile
        return (await getProfile()) !== null
      }
    } else {
      // No token, clear everything
      logout()
      return false
    }
  }

  // Initialize auth with stored data
  const initAuthFromStorage = () => {
    const storedUser = localStorage.getItem(STORAGE_KEYS.USER)
    const token = localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN)

    if (storedUser && token) {
      try {
        user.value = JSON.parse(storedUser)
      } catch (error) {
        console.error('Error parsing stored user:', error)
        logout()
      }
    }
  }

  const login = async (credentials: UserLoginRequest): Promise<boolean> => {
    try {
      const result = await post<AuthResponse>('/auth/login', credentials)

      if (result.ok && result.data.access_token && result.data.user) {
        // Store token and user
        localStorage.setItem(
          STORAGE_KEYS.ACCESS_TOKEN,
          result.data.access_token
        )
        localStorage.setItem(
          STORAGE_KEYS.USER,
          JSON.stringify(result.data.user)
        )
        user.value = result.data.user

        notifications.success('Login berhasil!')
        return true
      }

      return false
    } catch (error) {
      console.error('Login error:', error)
      return false
    }
  }

  const register = async (userData: UserRegisterRequest): Promise<boolean> => {
    try {
      const result = await post<AuthResponse>('/auth/register', userData)

      if (result.ok && result.data.access_token && result.data.user) {
        // Store token and user
        localStorage.setItem(
          STORAGE_KEYS.ACCESS_TOKEN,
          result.data.access_token
        )
        localStorage.setItem(
          STORAGE_KEYS.USER,
          JSON.stringify(result.data.user)
        )
        user.value = result.data.user

        notifications.success('Registrasi berhasil!')
        return true
      }

      return false
    } catch (error) {
      console.error('Register error:', error)
      return false
    }
  }

  const logout = () => {
    // Clear storage
    localStorage.removeItem(STORAGE_KEYS.ACCESS_TOKEN)
    localStorage.removeItem(STORAGE_KEYS.USER)
    user.value = null

    // Call logout endpoint (optional, for server-side cleanup)
    post('/auth/logout').catch(() => {
      // Ignore errors on logout
    })

    notifications.info('Anda telah logout')
  }

  const getProfile = async (): Promise<User | null> => {
    try {
      const result = await get<User>('/auth/me')

      if (result.ok) {
        // The API returns the user directly, not wrapped in a data object
        const userData = result.data as User
        user.value = userData
        localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(userData))
        return userData
      }

      return null
    } catch (error) {
      console.error('Get profile error:', error)
      return null
    }
  }

  // Initialize on composable creation
  initAuthFromStorage()

  return {
    user: computed(() => user.value),
    isAuthenticated,
    isAdmin,
    login,
    register,
    logout,
    getProfile,
    initAuth,
  }
}
