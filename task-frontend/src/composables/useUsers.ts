import { ref, computed } from 'vue'
import type { User, UsersResponse } from '@/types/api'
import { useApi } from './useApi'

export function useUsers() {
  const { get } = useApi()

  const users = ref<User[]>([])
  const loading = ref(false)
  const total = ref(0)
  const currentPage = ref(1)
  const limit = ref(10)

  // Computed
  const totalPages = computed(() => Math.ceil(total.value / limit.value))
  const hasNextPage = computed(() => currentPage.value < totalPages.value)
  const hasPrevPage = computed(() => currentPage.value > 1)

  // Fetch users (admin only)
  const fetchUsers = async (
    options: {
      page?: number
      limit?: number
    } = {}
  ) => {
    loading.value = true

    try {
      const params = new URLSearchParams()

      if (options.page) params.append('page', options.page.toString())
      if (options.limit) params.append('limit', options.limit.toString())

      const queryString = params.toString()
      const endpoint = `/admin/users${queryString ? `?${queryString}` : ''}`

      const result = await get<UsersResponse>(endpoint)

      if (result.ok) {
        // The API returns the UsersResponse directly
        const usersData = result.data as UsersResponse
        users.value = usersData.users || []
        total.value = usersData.total || 0
        currentPage.value = usersData.page || 1
        limit.value = usersData.limit || 10
      }
    } catch (error) {
      console.error('Fetch users error:', error)
      users.value = []
    } finally {
      loading.value = false
    }
  }

  // Pagination helpers
  const nextPage = () => {
    if (hasNextPage.value) {
      fetchUsers({ page: currentPage.value + 1, limit: limit.value })
    }
  }

  const prevPage = () => {
    if (hasPrevPage.value) {
      fetchUsers({ page: currentPage.value - 1, limit: limit.value })
    }
  }

  const goToPage = (page: number) => {
    if (page >= 1 && page <= totalPages.value) {
      fetchUsers({ page, limit: limit.value })
    }
  }

  return {
    // State
    users,
    loading,
    total,
    currentPage,
    limit,

    // Computed
    totalPages,
    hasNextPage,
    hasPrevPage,

    // Methods
    fetchUsers,

    // Pagination
    nextPage,
    prevPage,
    goToPage,
  }
}
