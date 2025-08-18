<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">My Tasks</h1>
      <p class="text-gray-600 mt-2">Manage your personal tasks</p>
    </div>

    <!-- Filters and Search -->
    <div class="card p-6 mb-6">
      <div class="flex flex-col md:flex-row md:items-center md:justify-between space-y-4 md:space-y-0 md:space-x-4">
        <div class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
          <!-- Search -->
          <div class="relative">
            <input v-model="searchQuery" type="text" placeholder="Search tasks..."
              class="form-input pl-10 w-full md:w-64" @input="debouncedSearch" />
            <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" fill="none"
              stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>

          <!-- Status Filter -->
          <select v-model="statusFilter" class="form-select" @change="handleFilterChange">
            <option value="">All Status</option>
            <option value="pending">Pending</option>
            <option value="in_progress">In Progress</option>
            <option value="completed">Completed</option>
          </select>
        </div>

        <!-- Create Task Button -->
        <button @click="showCreateModal = true" class="btn-primary">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          Create New Task
        </button>
      </div>
    </div>

    <!-- Tasks Grid -->
    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
      <p class="text-gray-500 mt-2">Loading tasks...</p>
    </div>

    <div v-else-if="tasks.length === 0" class="text-center py-12">
      <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
          d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
      </svg>
      <p class="text-gray-500 mb-4">No tasks found</p>
      <button @click="showCreateModal = true" class="btn-primary">
        Create Your First Task
      </button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
      <div v-for="task in tasks" :key="task.id" class="card p-6 hover:shadow-lg transition-shadow duration-200">
        <div class="flex justify-between items-start mb-4">
          <h3 class="text-lg font-semibold text-gray-900 flex-1 mr-2">
            {{ task.title }}
          </h3>
          <span :class="getTaskStatusClass(task.status)">
            {{ getTaskStatusLabel(task.status) }}
          </span>
        </div>

        <p v-if="task.description" class="text-gray-600 text-sm mb-4 line-clamp-3">
          {{ task.description }}
        </p>

        <div class="flex justify-between items-center text-xs text-gray-500 mb-4">
          <span>Created: {{ formatDate(task.created_at!) }}</span>
          <span v-if="task.updated_at">Updated: {{ formatDate(task.updated_at) }}</span>
        </div>

        <div class="flex justify-end space-x-2">
          <button @click="editTask(task)" class="text-blue-600 hover:text-blue-800 p-2 rounded-md hover:bg-blue-50"
            title="Edit Task">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
          </button>
          <button @click="confirmDelete(task)" class="text-red-600 hover:text-red-800 p-2 rounded-md hover:bg-red-50"
            title="Delete Task">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="flex justify-center">
      <div class="flex space-x-2">
        <button @click="prevPage" :disabled="!hasPrevPage" class="btn-secondary"
          :class="{ 'opacity-50 cursor-not-allowed': !hasPrevPage }">
          Previous
        </button>

        <div class="flex space-x-1">
          <button v-for="page in visiblePages" :key="page" @click="goToPage(page)" class="px-4 py-2 text-sm rounded-md"
            :class="page === currentPage
              ? 'bg-blue-600 text-white'
              : 'bg-gray-200 text-gray-700 hover:bg-gray-300'">
            {{ page }}
          </button>
        </div>

        <button @click="nextPage" :disabled="!hasNextPage" class="btn-secondary"
          :class="{ 'opacity-50 cursor-not-allowed': !hasNextPage }">
          Next
        </button>
      </div>
    </div>

    <!-- Create Task Modal -->
    <CreateTaskModal v-if="showCreateModal" @close="showCreateModal = false" @created="handleTaskCreated" />

    <!-- Edit Task Modal -->
    <EditTaskModal v-if="showEditModal && selectedTask" :task="selectedTask" @close="showEditModal = false"
      @updated="handleTaskUpdated" />

    <!-- Delete Confirmation Modal -->
    <ConfirmModal v-if="showDeleteModal && selectedTask" title="Delete Task"
      :message="`Are you sure you want to delete '${selectedTask.title}'? This action cannot be undone.`"
      confirm-text="Delete" @confirm="handleDeleteConfirm" @cancel="showDeleteModal = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import type { Task, TaskStatus } from '@/types/api';
import { useTasks } from '@/composables/useTasks';
import { formatDate, getTaskStatusLabel, getTaskStatusClass, debounce } from '@/utils/helpers';
import CreateTaskModal from '@/components/common/CreateTaskModal.vue';
import EditTaskModal from '@/components/common/EditTaskModal.vue';
import ConfirmModal from '@/components/common/ConfirmModal.vue';

const {
  tasks,
  loading,
  currentPage,
  limit,
  totalPages,
  hasNextPage,
  hasPrevPage,
  fetchTasks,
  deleteTask,
  nextPage,
  prevPage,
  goToPage,
} = useTasks();

// Local state
const searchQuery = ref('');
const statusFilter = ref<TaskStatus | ''>('');
const showCreateModal = ref(false);
const showEditModal = ref(false);
const showDeleteModal = ref(false);
const selectedTask = ref<Task | null>(null);

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

// Methods
const handleFilterChange = () => {
  fetchTasks({
    page: 1,
    limit: limit.value,
    status: statusFilter.value === '' ? undefined : statusFilter.value,
    search: searchQuery.value || undefined,
  });
};

const debouncedSearch = debounce(() => {
  handleFilterChange();
}, 500);

const editTask = (task: Task) => {
  selectedTask.value = task;
  showEditModal.value = true;
};

const confirmDelete = (task: Task) => {
  selectedTask.value = task;
  showDeleteModal.value = true;
};

const handleDeleteConfirm = async () => {
  if (!selectedTask.value) return;

  const success = await deleteTask(selectedTask.value.id);
  if (success) {
    showDeleteModal.value = false;
    selectedTask.value = null;
    // Refresh current page
    fetchTasks({
      page: currentPage.value,
      limit: limit.value,
      status: statusFilter.value === '' ? undefined : statusFilter.value,
      search: searchQuery.value || undefined,
    });
  }
};

const handleTaskCreated = () => {
  showCreateModal.value = false;
  // Refresh first page to show new task
  fetchTasks({
    page: 1,
    limit: limit.value,
    status: statusFilter.value === '' ? undefined : statusFilter.value,
    search: searchQuery.value || undefined,
  });
};

const handleTaskUpdated = () => {
  showEditModal.value = false;
  selectedTask.value = null;
  // Refresh current page
  fetchTasks({
    page: currentPage.value,
    limit: limit.value,
    status: statusFilter.value === '' ? undefined : statusFilter.value,
    search: searchQuery.value || undefined,
  });
};

// Lifecycle
onMounted(() => {
  fetchTasks({ page: 1, limit: 12 });
});
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>