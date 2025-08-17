import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import { ROUTES } from '@/utils/constants'

const routes: RouteRecordRaw[] = [
  {
    path: ROUTES.HOME,
    name: 'Home',
    component: () => import('@/components/layouts/PublicLayout.vue'),
    children: [
      {
        path: '',
        name: 'Landing',
        component: () => import('@/components/pages/LandingPage.vue'),
      },
      {
        path: ROUTES.LOGIN,
        name: 'Login',
        component: () => import('@/components/auth/LoginPage.vue'),
        meta: { requiresGuest: true },
      },
      {
        path: ROUTES.REGISTER,
        name: 'Register',
        component: () => import('@/components/auth/RegisterPage.vue'),
        meta: { requiresGuest: true },
      },
      {
        path: '/auth/callback',
        name: 'OAuthCallback',
        component: () => import('@/components/pages/OAuthCallbackPage.vue'),
        meta: { requiresGuest: true },
      },
    ],
  },
  {
    path: ROUTES.DASHBOARD,
    component: () => import('@/components/layouts/DashboardLayout.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('@/components/pages/DashboardPage.vue'),
      },
      {
        path: ROUTES.TASKS,
        name: 'Tasks',
        component: () => import('@/components/pages/TasksPage.vue'),
      },
      {
        path: ROUTES.PROFILE,
        name: 'Profile',
        component: () => import('@/components/pages/ProfilePage.vue'),
      },
    ],
  },
  {
    path: ROUTES.ADMIN,
    component: () => import('@/components/layouts/AdminLayout.vue'),
    meta: { requiresAuth: true, requiresAdmin: true },
    children: [
      {
        path: '',
        name: 'AdminDashboard',
        component: () =>
          import('@/components/pages/admin/AdminDashboardPage.vue'),
      },
      {
        path: 'users',
        name: 'AdminUsers',
        component: () => import('@/components/pages/admin/AdminUsersPage.vue'),
      },
      {
        path: 'tasks',
        name: 'AdminTasks',
        component: () => import('@/components/pages/admin/AdminTasksPage.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/components/pages/NotFoundPage.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const { isAuthenticated, isAdmin } = useAuth()

  // Check if route requires authentication
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    next(ROUTES.LOGIN)
    return
  }

  // Check if route requires admin role
  if (to.meta.requiresAdmin && !isAdmin.value) {
    next(ROUTES.DASHBOARD)
    return
  }

  // Check if route requires guest (not authenticated)
  if (to.meta.requiresGuest && isAuthenticated.value) {
    next(ROUTES.DASHBOARD)
    return
  }

  next()
})

export default router
