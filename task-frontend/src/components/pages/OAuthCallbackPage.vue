<template>
  <div class="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <div v-if="loading" class="mx-auto h-12 w-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
          <svg class="animate-spin h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
            </path>
          </svg>
        </div>

        <div v-else-if="error" class="mx-auto h-12 w-12 bg-red-100 rounded-full flex items-center justify-center mb-4">
          <svg class="h-6 w-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
          </svg>
        </div>

        <h2 class="mt-6 text-center text-2xl font-extrabold text-gray-900">
          {{ loading ? 'Processing...' : 'Authentication Failed' }}
        </h2>

        <p class="mt-2 text-center text-sm text-gray-600">
          {{ loading ? 'Please wait while we complete your authentication.' : error }}
        </p>

        <div v-if="error" class="mt-6">
          <router-link to="/login" class="btn-primary">
            Back to Login
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuth } from '@/composables/useAuth';
import { STORAGE_KEYS } from '@/utils/constants';

const route = useRoute();
const router = useRouter();
const { initAuth } = useAuth();

const loading = ref(true);
const error = ref('');

onMounted(async () => {
  try {
    const token = route.query.token as string;
    const successQuery = route.query.success as string;

    if (!token || successQuery !== 'true') {
      throw new Error('Authentication failed: Invalid callback parameters.');
    }

    // Store the token directly (it comes from backend redirect)
    localStorage.setItem(STORAGE_KEYS.ACCESS_TOKEN, token);

    // Initialize auth state (this will fetch user profile)
    const authSuccessful = await initAuth();

    if (authSuccessful) {
      router.push('/dashboard');
    } else {
      throw new Error('Authentication failed: Could not retrieve user profile.');
    }

  } catch (err) {
    console.error('OAuth callback error:', err);
    error.value = err instanceof Error ? err.message : 'An unknown error occurred during authentication.';
    loading.value = false;
  }
});
</script>