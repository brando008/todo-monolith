<script setup lang="ts">
import type { Todo } from '../services/todo.service'

defineProps<{
  task: Todo
}>()

const emit = defineEmits<{
    (e: 'toggle', task: Todo): void
    (e: 'delete', id: string): void
}>()
</script>

<template>
  <li class="flex justify-between items-center p-2 border border-terminal-bg hover:border-terminal-dim/50 transition-colors group">
    
    <label class="flex items-center gap-4 cursor-pointer flex-1">
      <!-- Hidden native checkbox, replaced with terminal brackets -->
      <input 
        type="checkbox" 
        :checked="task.completed"
        @change="emit('toggle', task)"
        class="hidden"
      >
      <!-- Custom Terminal Checkbox -->
      <span class="text-terminal-glow font-bold w-6">
        {{ task.completed ? '[X]' : '[]' }}
      </span>
      
      <!-- Text dynamically changes based on status -->
      <span :class="task.completed ? 'line-through text-terminal-dim' : 'text-terminal-text'">
        {{ task.title }}
      </span>
    </label>
    
    <!-- Delete button (Hidden until hover) -->
    <button 
      @click="emit('delete', task.id)"
      class="opacity-0 group-hover:opacity-100 text-red-500 hover:text-red-400 font-bold px-2 transition-opacity"
    >
      DEL
    </button>
  </li>
</template>