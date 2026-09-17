// src/services/todo.service.ts

// 1. The blueprint matching your future Go/PostgreSQL table
export interface Todo {
    id: string
    title: string
    completed: boolean
    created_at?: string
    updated_at?: string
  }

  const API_URL = '/todos';
  
  // 4. The API Service
  export const todoService = {
    
    async getAll(): Promise<Todo[]> {
        
        const response = await fetch(API_URL);
        if (!response.ok) {
          throw new Error('Failed to fetch todos');
        }
        return response.json();
    },

    async get(id: string): Promise<Todo> {
      const response = await fetch(`${API_URL}/${id}`);
      if (!response.ok) throw new Error('Failed to fetch task');
      return response.json();
    },
  
    async create(title: string): Promise<Todo> {
      const newTodo: Todo = {
        id: crypto.randomUUID(), // Creates a random string ID
        title,
        completed: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }

      const response = await fetch(API_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newTodo)
      });

      if (!response.ok) {
        throw new Error('Failed to create todo');
      }
      return response.json();
    },
  
    async update(id: string, updates: Partial<Todo>): Promise<void> {
      if (updates.completed !== undefined) {
        const response = await fetch(`${API_URL}/${id}/done`, {
          method: 'PATCH'
        })
        if (!response.ok) {
          throw new Error('Failed to update todo');
        }
      }
    },

    async done(id: string): Promise<void> {
      const response = await fetch(`${API_URL}/${id}/done`, {
        method: 'PATCH'
      });
      if (!response.ok) throw new Error('Failed to toggle task status');
    },
  
    async delete(id: string): Promise<void> {
      const response = await fetch(`${API_URL}/${id}`, {
        method: 'DELETE'
      });
      if (!response.ok) {
        throw new Error('Failed to delete todo');
      }
    }
  }