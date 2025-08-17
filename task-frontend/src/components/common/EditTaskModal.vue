<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50" @click="closeModal">
    <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white" @click.stop>
      <div class="mt-3">
        <!-- Header -->
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium text-gray-900">Edit Task</h3>
          <button
            @click="closeModal"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div>
              <label for="title" class="block text-sm font-medium text-gray-700">
                Title *
              </label>
              <input
                id="title"
                v-model="form.title"
                type="text"
                required
                class="form-input mt-1"
                :class="{ 'border-red-500': errors.title }"
                placeholder="Enter task title"
              />
              <p v-if="errors.title" class="mt-1 text-sm text-red-600">
                {{ errors.title }}
              </p>
            </div>

            <div>
              <label for="description" class="block text-sm font-medium text-gray-700">
                Description
              </label>
              <textarea
                id="description"
                v-model="form.description"
                rows="3"
                class="form-input mt-1"
                placeholder="Enter task description (optional)"
              ></textarea>
            </div>

            <div>
              <label for="status" class="block text-sm font-medium text-gray-700">
                Status
              </label>
              <select
                id="status"
                v-model="form.status"
                class="form-select mt-1"
              >
                <option value="pending">Pending</option>
                <option value="in_progress">In Progress</option>
                <option value="completed">Completed</option>
              </select>
            </div>
          </div>

          <div v-if="submitError" class="mt-4 bg-red-50 border border-red-200 rounded-md p-4">
            <p class="text-sm text-red-600">{{ submitError }}</p>
          </div>

          <!-- Actions -->
          <div class="flex justify-end space-x-3 mt-6">
            <button
              type="button"
              @click="closeModal"
              class="btn-secondary"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="btn-primary"
            >
              <span v-if="loading" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Updating...
              </span>
              <span v-else>Update Task</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import type { Task, TaskStatus } from '@/types/api';
import { useTasks } from '@/composables/useTasks';

interface Props {
  task: Task;
}

const props = defineProps<Props>();

const emit = defineEmits<{
  close: [];
  updated: [];
}>();

const { updateTask } = useTasks();

const loading = ref(false);
const submitError = ref('');

const form = reactive({
  title: '',
  description: '',
  status: 'pending' as TaskStatus,
});

const errors = reactive({
  title: '',
});

const validateForm = (): boolean => {
  errors.title = '';
  
  let isValid = true;

  if (!form.title.trim()) {
    errors.title = 'Title is required';
    isValid = false;
  } else if (form.title.length > 255) {
    errors.title = 'Title must be less than 255 characters';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = async () => {
  if (!validateForm()) return;

  loading.value = true;
  submitError.value = '';

  try {
    const taskData = {
      title: form.title.trim(),
      description: form.description.trim() || undefined,
      status: form.status,
    };

    const result = await updateTask(props.task.id, taskData);
    
    if (result) {
      emit('updated');
    } else {
      submitError.value = 'Failed to update task. Please try again.';
    }
  } catch (error) {
    submitError.value = 'An error occurred while updating the task.';
    console.error('Update task error:', error);
  } finally {
    loading.value = false;
  }
};

const closeModal = () => {
  emit('close');
};

onMounted(() => {
  // Initialize form with task data
  form.title = props.task.title;
  form.description = props.task.description || '';
  form.status = props.task.status;
});
</script>