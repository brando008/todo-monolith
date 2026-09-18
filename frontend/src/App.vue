<script setup lang="ts">
import {ref, onMounted } from 'vue'
import { todoService, type Todo } from './services/todo.service'


import TodoForm from './components/TodoForm.vue'
import TodoItem from './components/TodoItem.vue'
import HistoryItem from './components/HistoryItem.vue'


const todos = ref<Todo[]>([])
const isLoading = ref(true)
const currentMode = ref<'list' | 'history'>('list')
const historyTodos = ref<Todo[]>([])

onMounted(async () => {
 try {
   todos.value = await todoService.getAll()
 } catch (error) {
   console.error("Failed to fetch: ", error)
 } finally {
   isLoading.value = false
 }
})


const handleAdd= async (title: string) => {
 try {
   const newTodo = await todoService.create(title)
   todos.value.push(newTodo)
 } catch (error) {
   console.error("Failed to add todo: ", error)
 }
}

const handleModeChange = async (mode: 'list' | 'history') => {
  currentMode.value = mode
  if (mode === 'history') {
    try {
      historyTodos.value = await todoService.getHistory()
    } catch (error) {
      console.error("Failed to fetch history: ", error)
    } finally {
      isLoading.value = false
    }
  }
}

const handleDelete = async (id:string) => {
 try {
   await todoService.delete(id)
   todos.value = todos.value.filter(task => task.id !== id)
 } catch (error) {
   console.error("Failed to delete todo: ", error)
 }
}


const handleToggle = async (task: Todo) => {
  task.completed = !task.completed

 try {
   await todoService.done(task.id)
 } catch (error) {
   task.completed = !task.completed
 }
}
</script>


<template>
  <!-- Full screen deep abyss background -->
  <div class="min-h-screen bg-terminal-bg p-4 md:p-8 flex flex-col items-center justify-center selection:bg-terminal-dim selection:text-terminal-text">
    
    <!-- The Window Wrapper (Uses your shortcut) -->
    <div class="term-window w-full max-w-2xl h-[80vh] flex flex-col">
      
      <!-- Window Header -->
      <div class="term-header">
        <span>TODO_TASKS_SYSTEM :: ROOT</span>
        <span class="text-terminal-glow animate-pulse">_</span>
      </div>

      <!-- Window Body -->
      <div class="term-body flex-1 flex flex-col overflow-y-auto">
        
        <TodoForm :current-mode="currentMode" @add="handleAdd" @changeMode="handleModeChange"/>

        <div v-if="isLoading" class="mt-8 text-terminal-glow animate-[terminal-blink_1s_steps(2,start)_infinite]">
          Loading datastore...
        </div>

        <ul v-else class="space-y-2 mt-6">
          <div v-if="currentMode === 'list'">
            <TodoItem 
              v-for="task in todos" 
              :key="task.id" 
              :task="task" 
              @toggle="handleToggle" 
              @delete="handleDelete" 
            />
          </div>

          <div v-else-if="currentMode === 'history'">         
            <HistoryItem 
              v-for="(task, index) in historyTodos" 
              :key="task.id" 
              :task="task"
              :index="index"
            />
          </div>
        </ul>
        
      </div>
    </div>
  </div>
</template>

