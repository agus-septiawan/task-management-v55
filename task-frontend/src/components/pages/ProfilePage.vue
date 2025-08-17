<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Profile</h1>
      <p class="text-gray-600 mt-2">Manage your account information</p>
    </div>

    <div class="max-w-2xl">
      <!-- Profile Information -->
      <div class="card p-6 mb-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Profile Information</h2>
        
        <div class="space-y-4">
          <div class="flex items-center space-x-4">
            <div class="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center">
              <span class="text-white font-bold text-xl">
                {{ user?.name?.charAt(0).toUpperCase() }}
              </span>
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900">{{ user?.name }}</h3>
              <p class="text-gray-600">{{ user?.email }}</p>
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800 mt-1">
                {{ user?.role === 'admin' ? 'Administrator' : 'User' }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Account Details -->
      <div class="card p-6 mb-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Account Details</h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Full Name</label>
            <div class="mt-1 text-sm text-gray-900">{{ user?.name }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">Email Address</label>
            <div class="mt-1 text-sm text-gray-900">{{ user?.email }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">Role</label>
            <div class="mt-1 text-sm text-gray-900">
              {{ user?.role === 'admin' ? 'Administrator' : 'User' }}
            </div>
          </div>
          
          <div v-if="user?.oauth_provider">
            <label class="block text-sm font-medium text-gray-700">OAuth Provider</label>
            <div class="mt-1 text-sm text-gray-900 capitalize">
              {{ user.oauth_provider }}
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700">Member Since</label>
            <div class="mt-1 text-sm text-gray-900">
              {{ user?.created_at ? formatDate(user.created_at) : 'N/A' }}
            </div>
          </div>
        </div>
      </div>

      <!-- Account Statistics -->
      <div class="card p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Your Statistics</h2>
        
        <div v-if="statsLoading" class="text-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
          <p class="text-gray-500 mt-2">Loading statistics...</p>
        </div>
        
        <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div class="text-center p-4 bg-yellow-50 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600">{{ taskStats.pending }}</div>
            <div class="text-sm text-yellow-700">Pending Tasks</div>
          </div>
          
          <div class="text-center p-4 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ taskStats.inProgress }}</div>
            <div class="text-sm text-blue-700">In Progress</div>
          </div>
          
          <div class="text-center p-4 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ taskStats.completed }}</div>
            <div class="text-sm text-green-700">Completed</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useAuth } from '@/composables/useAuth';
import { useTasks } from '@/composables/useTasks';
import { formatDate } from '@/utils/helpers';

const { user } = useAuth();
const { tasks, loading: tasksLoading, fetchTasks } = useTasks();

const statsLoading = ref(true);

// Task statistics
const taskStats = computed(() => {
  const pending = tasks.value.filter(task => task.status === 'pending').length;
  const inProgress = tasks.value.filter(task => task.status === 'in_progress').length;
  const completed = tasks.value.filter(task => task.status === 'completed').length;
  
  return {
    pending,
    inProgress,
    completed,
  };
});

onMounted(async () => {
  try {
    // Fetch all user tasks to calculate statistics
    await fetchTasks({ limit: 1000 }); // Large limit to get all tasks
  } finally {
    statsLoading.value = false;
  }
});
</script>