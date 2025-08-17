<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Users Management</h1>
      <p class="text-gray-600 mt-2">Manage all users in the system</p>
    </div>

    <!-- Users Table -->
    <div class="card overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">
          Users ({{ total }} total)
        </h2>
      </div>

      <div v-if="loading" class="text-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
        <p class="text-gray-500 mt-2">Loading users...</p>
      </div>

      <div v-else-if="users.length === 0" class="text-center py-12">
        <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
        </svg>
        <p class="text-gray-500">No users found</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                User
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Email
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Role
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                OAuth Provider
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Joined
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-10 h-10 bg-blue-600 rounded-full flex items-center justify-center mr-4">
                    <span class="text-white font-medium">
                      {{ user.name.charAt(0).toUpperCase() }}
                    </span>
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">ID: {{ user.id }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ user.email }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="user.role === 'admin'
                  ? 'bg-red-100 text-red-800'
                  : 'bg-blue-100 text-blue-800'">
                  {{ user.role === 'admin' ? 'Administrator' : 'User' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span v-if="user.oauth_provider"
                  class="inline-flex items-center px-2 py-1 text-xs font-medium bg-green-100 text-green-800 rounded-full">
                  <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 24 24">
                    <path
                      d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" />
                    <path
                      d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" />
                    <path
                      d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" />
                    <path
                      d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" />
                  </svg>
                  {{ user.oauth_provider }}
                </span>
                <span v-else class="text-gray-400">Email/Password</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.created_at!) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-700">
            Showing {{ ((currentPage - 1) * limit) + 1 }} to {{ Math.min(currentPage * limit, total) }} of {{ total }}
            results
          </div>
          <div class="flex space-x-2">
            <button @click="prevPage" :disabled="!hasPrevPage" class="btn-secondary text-sm"
              :class="{ 'opacity-50 cursor-not-allowed': !hasPrevPage }">
              Previous
            </button>

            <div class="flex space-x-1">
              <button v-for="page in visiblePages" :key="page" @click="goToPage(page)"
                class="px-3 py-1 text-sm rounded-md" :class="page === currentPage
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'">
                {{ page }}
              </button>
            </div>

            <button @click="nextPage" :disabled="!hasNextPage" class="btn-secondary text-sm"
              :class="{ 'opacity-50 cursor-not-allowed': !hasNextPage }">
              Next
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useUsers } from '@/composables/useUsers';
import { formatDate } from '@/utils/helpers';

const {
  users,
  loading,
  total,
  currentPage,
  limit,
  totalPages,
  hasNextPage,
  hasPrevPage,
  fetchUsers,
  nextPage,
  prevPage,
  goToPage,
} = useUsers();

// Computed
const visiblePages = computed(() => {
  const pages = [];
  const start = Math.max(1, currentPage.value - 2);
  const end = Math.min(totalPages.value, currentPage.value + 2);

  for (let i = start; i <= end; i++) {
    pages.push(i);
  }

  return pages;
});

// Lifecycle
onMounted(() => {
  fetchUsers({ page: 1, limit: 10 });
});
</script>