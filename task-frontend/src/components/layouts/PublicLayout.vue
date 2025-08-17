<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow-sm border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <router-link to="/" class="flex items-center space-x-2">
              <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
                <span class="text-white font-bold text-sm">TM</span>
              </div>
              <span class="text-xl font-semibold text-gray-900">Task Manager</span>
            </router-link>
          </div>
          
          <div class="flex items-center space-x-4">
            <template v-if="!isAuthenticated">
              <router-link
                to="/login"
                class="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
              >
                Login
              </router-link>
              <router-link
                to="/register"
                class="btn-primary text-sm"
              >
                Register
              </router-link>
            </template>
            <template v-else>
              <router-link
                to="/dashboard"
                class="btn-primary text-sm"
              >
                Dashboard
              </router-link>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main>
      <router-view />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t mt-auto">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <div class="text-center text-gray-500 text-sm">
          © 2024 Task Management System. Built with Vue 3 & Go.
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth';

const { isAuthenticated } = useAuth();
</script>