<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <section class="bg-gradient-to-br from-blue-600 to-blue-800 text-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
        <div class="text-center">
          <h1 class="text-4xl md:text-6xl font-bold mb-6">
            Manage Your Tasks
            <span class="text-blue-200">Efficiently</span>
          </h1>
          <p class="text-xl md:text-2xl mb-8 text-blue-100 max-w-3xl mx-auto">
            A powerful task management system built with modern technologies. 
            Organize, track, and complete your tasks with ease.
          </p>
          <div class="space-x-4">
            <router-link
              to="/register"
              class="inline-block bg-white text-blue-600 font-semibold px-8 py-3 rounded-lg hover:bg-gray-100 transition-colors duration-200"
            >
              Get Started
            </router-link>
            <router-link
              to="/login"
              class="inline-block border-2 border-white text-white font-semibold px-8 py-3 rounded-lg hover:bg-white hover:text-blue-600 transition-colors duration-200"
            >
              Sign In
            </router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Why Choose Our Task Manager?
          </h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto">
            Built with cutting-edge technology to provide you with the best task management experience.
          </p>
        </div>

        <div class="grid md:grid-cols-3 gap-8">
          <div class="text-center p-6">
            <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Easy Task Management</h3>
            <p class="text-gray-600">
              Create, update, and organize your tasks with an intuitive interface designed for productivity.
            </p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Secure & Reliable</h3>
            <p class="text-gray-600">
              Your data is protected with JWT authentication and secure API endpoints.
            </p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Fast & Modern</h3>
            <p class="text-gray-600">
              Built with Vue 3, TypeScript, and Go for optimal performance and developer experience.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Tech Stack Section -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Built with Modern Technology
          </h2>
          <p class="text-xl text-gray-600">
            Leveraging the best tools and frameworks for a superior experience.
          </p>
        </div>

        <div class="grid md:grid-cols-2 gap-12 items-center">
          <div>
            <h3 class="text-2xl font-bold text-gray-900 mb-6">Frontend</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-green-500 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">V</span>
                </div>
                <span class="text-lg text-gray-700">Vue 3 with Composition API</span>
              </div>
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-blue-500 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">TS</span>
                </div>
                <span class="text-lg text-gray-700">TypeScript for type safety</span>
              </div>
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-cyan-500 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">T</span>
                </div>
                <span class="text-lg text-gray-700">Tailwind CSS for styling</span>
              </div>
            </div>
          </div>

          <div>
            <h3 class="text-2xl font-bold text-gray-900 mb-6">Backend</h3>
            <div class="space-y-4">
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-blue-600 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">Go</span>
                </div>
                <span class="text-lg text-gray-700">Go (Golang) for performance</span>
              </div>
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-orange-500 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">M</span>
                </div>
                <span class="text-lg text-gray-700">MySQL for data storage</span>
              </div>
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 bg-gray-700 rounded flex items-center justify-center">
                  <span class="text-white font-bold text-sm">J</span>
                </div>
                <span class="text-lg text-gray-700">JWT for authentication</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-20 bg-blue-600">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-white mb-4">
          Ready to Get Started?
        </h2>
        <p class="text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
          Join thousands of users who are already managing their tasks more efficiently.
        </p>
        <router-link
          to="/register"
          class="inline-block bg-white text-blue-600 font-semibold px-8 py-3 rounded-lg hover:bg-gray-100 transition-colors duration-200"
        >
          Create Your Account
        </router-link>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// Landing page component
</script>