<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>
      <p class="text-gray-600 mt-2">Welcome back, {{ user?.name }}!</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Pending Tasks</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.pending }}</p>
          </div>
        </div>
      </div>

      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">In Progress</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.inProgress }}</p>
          </div>
        </div>
      </div>

      <div class="card p-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">Completed</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.completed }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="card p-6 mb-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
      <div class="flex flex-wrap gap-4">
        <button
          @click="showCreateModal = true"
          class="btn-primary"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Create New Task
        </button>
        
        <router-link to="/tasks" class="btn-secondary">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
          View All Tasks
        </router-link>
      </div>
    </div>

    <!-- Recent Tasks -->
    <div class="card p-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl font-semibold text-gray-900">Recent Tasks</h2>
        <router-link to="/tasks" class="text-blue-600 hover:text-blue-700 text-sm font-medium">
          View all →
        </router-link>
      </div>

      <div v-if="loading" class="text-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
        <p class="text-gray-500 mt-2">Loading tasks...</p>
      </div>

      <div v-else-if="recentTasks.length === 0" class="text-center py-8">
        <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <p class="text-gray-500">No tasks yet. Create your first task to get started!</p>
      </div>

      <div v-else class="space-y-3">
        <div
          v-for="task in recentTasks"
          :key="task.id"
          class="flex items-center justify-between p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors duration-200"
        >
          <div class="flex-1">
            <h3 class="font-medium text-gray-900">{{ task.title }}</h3>
            <p v-if="task.description" class="text-sm text-gray-600 mt-1">
              {{ truncateText(task.description, 80) }}
            </p>
            <p class="text-xs text-gray-500 mt-2">
              {{ formatDate(task.created_at!) }}
            </p>
          </div>
          <div class="ml-4">
            <span :class="getTaskStatusClass(task.status)">
              {{ getTaskStatusLabel(task.status) }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Task Modal -->
    <CreateTaskModal
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @created="handleTaskCreated"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useAuth } from '@/composables/useAuth';
import { useTasks } from '@/composables/useTasks';
import { formatDate, getTaskStatusLabel, getTaskStatusClass, truncateText } from '@/utils/helpers';
import CreateTaskModal from '@/components/common/CreateTaskModal.vue';

const { user } = useAuth();
const { tasks, loading, fetchTasks } = useTasks();

const showCreateModal = ref(false);

// Recent tasks (limit to 5)
const recentTasks = computed(() => tasks.value.slice(0, 5));

// Stats
const stats = computed(() => {
  const pending = tasks.value.filter(task => task.status === 'pending').length;
  const inProgress = tasks.value.filter(task => task.status === 'in_progress').length;
  const completed = tasks.value.filter(task => task.status === 'completed').length;
  
  return {
    pending,
    inProgress,
    completed,
  };
});

const handleTaskCreated = () => {
  showCreateModal.value = false;
  fetchTasks({ limit: 20 }); // Refresh tasks
};

onMounted(() => {
  fetchTasks({ limit: 20 }); // Load recent tasks
});
</script>