import { ref, computed } from 'vue';
import type { Task, TaskCreateRequest, TaskUpdateRequest, TasksResponse, TaskStatus } from '@/types/api';
import { useApi } from './useApi';
import { notifications } from '@/utils/notifications';

export function useTasks() {
  const { get, post, put, delete: del } = useApi();
  
  const tasks = ref<Task[]>([]);
  const loading = ref(false);
  const total = ref(0);
  const currentPage = ref(1);
  const limit = ref(10);

  // Computed
  const totalPages = computed(() => Math.ceil(total.value / limit.value));
  const hasNextPage = computed(() => currentPage.value < totalPages.value);
  const hasPrevPage = computed(() => currentPage.value > 1);

  // Fetch tasks with filters
  const fetchTasks = async (options: {
    page?: number;
    limit?: number;
    status?: TaskStatus | '';
    search?: string;
  } = {}) => {
    loading.value = true;
    
    try {
      const params = new URLSearchParams();
      
      if (options.page) params.append('page', options.page.toString());
      if (options.limit) params.append('limit', options.limit.toString());
      if (options.status) params.append('status', options.status);
      if (options.search) params.append('search', options.search);
      
      const queryString = params.toString();
      const endpoint = `/tasks${queryString ? `?${queryString}` : ''}`;
      
      const result = await get<TasksResponse>(endpoint);
      
      if (result.ok && result.data) {
        tasks.value = result.data.tasks || [];
        total.value = result.data.total || 0;
        currentPage.value = result.data.page || 1;
        limit.value = result.data.limit || 10;
      }
    } catch (error) {
      console.error('Fetch tasks error:', error);
      tasks.value = [];
    } finally {
      loading.value = false;
    }
  };

  // Get single task
  const getTask = async (id: number): Promise<Task | null> => {
    try {
      const result = await get<Task>(`/tasks/${id}`);
      return result.ok && result.data ? result.data : null;
    } catch (error) {
      console.error('Get task error:', error);
      return null;
    }
  };

  // Create task
  const createTask = async (taskData: TaskCreateRequest): Promise<Task | null> => {
    try {
      const result = await post<Task>('/tasks', taskData);
      
      if (result.ok && result.data) {
        notifications.success('Task berhasil dibuat!');
        // Refresh tasks list
        await fetchTasks({ page: currentPage.value, limit: limit.value });
        return result.data;
      }
      
      return null;
    } catch (error) {
      console.error('Create task error:', error);
      return null;
    }
  };

  // Update task
  const updateTask = async (id: number, taskData: TaskUpdateRequest): Promise<Task | null> => {
    try {
      const result = await put<Task>(`/tasks/${id}`, taskData);
      
      if (result.ok && result.data) {
        notifications.success('Task berhasil diupdate!');
        
        // Update task in local array
        const index = tasks.value.findIndex(task => task.id === id);
        if (index !== -1) {
          tasks.value[index] = result.data;
        }
        
        return result.data;
      }
      
      return null;
    } catch (error) {
      console.error('Update task error:', error);
      return null;
    }
  };

  // Delete task
  const deleteTask = async (id: number): Promise<boolean> => {
    try {
      const result = await del(`/tasks/${id}`);
      
      if (result.ok) {
        notifications.success('Task berhasil dihapus!');
        
        // Remove task from local array
        tasks.value = tasks.value.filter(task => task.id !== id);
        total.value = Math.max(0, total.value - 1);
        
        return true;
      }
      
      return false;
    } catch (error) {
      console.error('Delete task error:', error);
      return false;
    }
  };

  // Pagination helpers
  const nextPage = () => {
    if (hasNextPage.value) {
      fetchTasks({ page: currentPage.value + 1, limit: limit.value });
    }
  };

  const prevPage = () => {
    if (hasPrevPage.value) {
      fetchTasks({ page: currentPage.value - 1, limit: limit.value });
    }
  };

  const goToPage = (page: number) => {
    if (page >= 1 && page <= totalPages.value) {
      fetchTasks({ page, limit: limit.value });
    }
  };

  return {
    // State
    tasks,
    loading,
    total,
    currentPage,
    limit,
    
    // Computed
    totalPages,
    hasNextPage,
    hasPrevPage,
    
    // Methods
    fetchTasks,
    getTask,
    createTask,
    updateTask,
    deleteTask,
    
    // Pagination
    nextPage,
    prevPage,
    goToPage,
  };
}