<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">All Tasks Management</h1>
      <p class="text-gray-600 mt-2">Manage all tasks across the system</p>
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

    <!-- Tasks Table -->
    <div class="card overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">
          Tasks ({{ total }} total)
        </h2>
      </div>

      <div v-if="loading" class="text-center py-12">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
        <p class="text-gray-500 mt-2">Loading tasks...</p>
      </div>

      <div v-else-if="tasks.length === 0" class="text-center py-12">
        <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <p class="text-gray-500">No tasks found</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Task
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Owner
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Created
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="task in tasks" :key="task.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ task.title }}</div>
                  <div v-if="task.description" class="text-sm text-gray-500 mt-1">
                    {{ truncateText(task.description, 100) }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center mr-3">
                    <span class="text-gray-600 font-medium text-sm">
                      {{ getUserInitial(task.user_id) }}
                    </span>
                  </div>
                  <div class="text-sm text-gray-900">
                    User #{{ task.user_id }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTaskStatusClass(task.status)">
                  {{ getTaskStatusLabel(task.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(task.created_at!) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <div class="flex justify-end space-x-2">
                  <button @click="editTask(task)" class="text-blue-600 hover:text-blue-900" title="Edit Task">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button @click="confirmDelete(task)" class="text-red-600 hover:text-red-900" title="Delete Task">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
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
import type { Task } from '@/types/api';
import type { TaskStatus } from '@/types/api';
import { useTasks } from '@/composables/useTasks';
import { formatDate, getTaskStatusLabel, getTaskStatusClass, truncateText, debounce } from '@/utils/helpers';
import CreateTaskModal from '@/components/common/CreateTaskModal.vue';
import EditTaskModal from '@/components/common/EditTaskModal.vue';
import ConfirmModal from '@/components/common/ConfirmModal.vue';

const {
  tasks,
  loading,
  total,
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
const getUserInitial = (userId: number): string => {
  return `U${userId}`;
};

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
  fetchTasks({ page: 1, limit: 10 });
});
</script>