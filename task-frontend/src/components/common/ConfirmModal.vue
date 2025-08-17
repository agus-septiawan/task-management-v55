<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50" @click="handleCancel">
    <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white" @click.stop>
      <div class="mt-3 text-center">
        <!-- Icon -->
        <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100">
          <svg class="h-6 w-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
        </div>
        
        <!-- Title -->
        <h3 class="text-lg leading-6 font-medium text-gray-900 mt-4">
          {{ title }}
        </h3>
        
        <!-- Message -->
        <div class="mt-2 px-7 py-3">
          <p class="text-sm text-gray-500">
            {{ message }}
          </p>
        </div>
        
        <!-- Actions -->
        <div class="flex justify-center space-x-3 mt-6">
          <button
            @click="handleCancel"
            class="btn-secondary"
          >
            Cancel
          </button>
          <button
            @click="handleConfirm"
            class="btn-danger"
          >
            {{ confirmText }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string;
  message: string;
  confirmText?: string;
}

const props = withDefaults(defineProps<Props>(), {
  confirmText: 'Confirm',
});

const emit = defineEmits<{
  confirm: [];
  cancel: [];
}>();

const handleConfirm = () => {
  emit('confirm');
};

const handleCancel = () => {
  emit('cancel');
};
</script>