// Simple notification system without external dependencies
export interface NotificationOptions {
  title?: string;
  message: string;
  type?: 'success' | 'error' | 'warning' | 'info';
  duration?: number;
}

class NotificationManager {
  private container: HTMLElement | null = null;

  private createContainer() {
    if (this.container) return;
    
    this.container = document.createElement('div');
    this.container.className = 'fixed top-4 right-4 z-50 space-y-2';
    document.body.appendChild(this.container);
  }

  private createNotification(options: NotificationOptions) {
    this.createContainer();
    
    const notification = document.createElement('div');
    const typeClasses = {
      success: 'bg-green-500 text-white',
      error: 'bg-red-500 text-white',
      warning: 'bg-yellow-500 text-white',
      info: 'bg-blue-500 text-white',
    };
    
    notification.className = `
      ${typeClasses[options.type || 'info']} 
      px-4 py-3 rounded-lg shadow-lg max-w-sm transform transition-all duration-300 translate-x-full opacity-0
    `;
    
    notification.innerHTML = `
      <div class="flex items-center justify-between">
        <div>
          ${options.title ? `<div class="font-semibold">${options.title}</div>` : ''}
          <div class="text-sm">${options.message}</div>
        </div>
        <button class="ml-4 text-white hover:text-gray-200" onclick="this.parentElement.parentElement.remove()">
          ×
        </button>
      </div>
    `;
    
    this.container!.appendChild(notification);
    
    // Animate in
    setTimeout(() => {
      notification.classList.remove('translate-x-full', 'opacity-0');
    }, 10);
    
    // Auto remove
    setTimeout(() => {
      notification.classList.add('translate-x-full', 'opacity-0');
      setTimeout(() => notification.remove(), 300);
    }, options.duration || 5000);
  }

  success(message: string, title?: string) {
    this.createNotification({ message, title, type: 'success' });
  }

  error(message: string, title?: string) {
    this.createNotification({ message, title, type: 'error' });
  }

  warning(message: string, title?: string) {
    this.createNotification({ message, title, type: 'warning' });
  }

  info(message: string, title?: string) {
    this.createNotification({ message, title, type: 'info' });
  }
}

export const notifications = new NotificationManager();