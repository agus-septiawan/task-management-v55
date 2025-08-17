<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Admin Dashboard</h1>
      <p class="text-gray-600 mt-2">System overview and management</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Total Users</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalUsers }}</p>
          </div>
        </div>
      </div>

      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Total Tasks</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalTasks }}</p>
          </div>
        </div>
      </div>

      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Pending Tasks</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.pendingTasks }}</p>
          </div>
        </div>
      </div>

      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Completed Tasks</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.completedTasks }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="card p-6 mb-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
      <div class="flex flex-wrap gap-4">
        <router-link to="/admin/users" class="btn-primary">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
          </svg>
          Manage Users
        </router-link>
        
        <router-link to="/admin/tasks" class="btn-secondary">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
          View All Tasks
        </router-link>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Recent Users -->
      <div class="card p-6">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold text-gray-900">Recent Users</h2>
          <router-link to="/admin/users" class="text-blue-600 hover:text-blue-700 text-sm font-medium">
            View all →
          </router-link>
        </div>

        <div v-if="usersLoading" class="text-center py-8">
          <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 mx-auto"></div>
          <p class="text-gray-500 mt-2 text-sm">Loading users...</p>
        </div>

        <div v-else-if="recentUsers.length === 0" class="text-center py-8">
          <p class="text-gray-500">No users found</p>
        </div>

        <div v-else class="space-y-3">
          <div
            v-for="user in recentUsers"
            :key="user.id"
            class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <div class="flex items-center space-x-3">
              <div class="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center">
                <span class="text-white font-medium text-sm">
                  {{ user.name.charAt(0).toUpperCase() }}
                </span>
              </div>
              <div>
                <p class="font-medium text-gray-900 text-sm">{{ user.name }}</p>
                <p class="text-xs text-gray-500">{{ user.email }}</p>
              </div>
            </div>
            <span class="text-xs px-2 py-1 rounded-full" :class="user.role === 'admin' ? 'bg-red-100 text-red-800' : 'bg-blue-100 text-blue-800'">
              {{ user.role }}
            </span>
          </div>
        </div>
      </div>

      <!-- Recent Tasks -->
      <div class="card p-6">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold text-gray-900">Recent Tasks</h2>
          <router-link to="/admin/tasks" class="text-blue-600 hover:text-blue-700 text-sm font-medium">
            View all →
          </router-link>
        </div>

        <div v-if="tasksLoading" class="text-center py-8">
          <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 mx-auto"></div>
          <p class="text-gray-500 mt-2 text-sm">Loading tasks...</p>
        </div>

        <div v-else-if="recentTasks.length === 0" class="text-center py-8">
          <p class="text-gray-500">No tasks found</p>
        </div>

        <div v-else class="space-y-3">
          <div
            v-for="task in recentTasks"
            :key="task.id"
            class="p-3 bg-gray-50 rounded-lg"
          >
            <div class="flex justify-between items-start">
              <div class="flex-1">
                <p class="font-medium text-gray-900 text-sm">{{ task.title }}</p>
                <p v-if="task.description" class="text-xs text-gray-600 mt-1">
                  {{ truncateText(task.description, 60) }}
                </p>
              </div>
              <span :class="getTaskStatusClass(task.status)">
                {{ getTaskStatusLabel(task.status) }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useUsers } from '@/composables/useUsers';
import { useTasks } from '@/composables/useTasks';
import { formatDate, getTaskStatusLabel, getTaskStatusClass, truncateText } from '@/utils/helpers';

const { users, loading: usersLoading, fetchUsers } = useUsers();
const { tasks, loading: tasksLoading, fetchTasks } = useTasks();

// Recent data (limit to 5)
const recentUsers = computed(() => users.value.slice(0, 5));
const recentTasks = computed(() => tasks.value.slice(0, 5));

// Statistics
const stats = computed(() => {
  const totalUsers = users.value.length;
  const totalTasks = tasks.value.length;
  const pendingTasks = tasks.value.filter(task => task.status === 'pending').length;
  const completedTasks = tasks.value.filter(task => task.status === 'completed').length;
  
  return {
    totalUsers,
    totalTasks,
    pendingTasks,
    completedTasks,
  };
});

onMounted(async () => {
  // Fetch recent data
  await Promise.all([
    fetchUsers({ limit: 10 }),
    fetchTasks({ limit: 10 }),
  ]);
});
</script>