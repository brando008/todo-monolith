<script setup lang="ts">
import {ref} from 'vue'

const emit = defineEmits<{
    (e: 'add', title:string): void
    (e: 'changeMode', mode: 'list' | 'history'): void
}>()

const props = defineProps<{
  currentMode: 'list' | 'history'
}>()

const newTodoTitle = ref('')
const currentMode = ref('list')
const inputRef = ref<HTMLInputElement | null>(null)

// Handles the input from the HTML and then verifies the input before emitting it to the higher component (App.vue)
const handleSubmit = () => {
  const text = newTodoTitle.value.trim()
    if (!text) return

    if (text == "/history") {
      currentMode.value = 'history'
      emit('changeMode', 'history')
      newTodoTitle.value = ''
      return
    }
   
    if (newTodoTitle.value == "/todo") {
      currentMode.value = 'list'
      emit('changeMode', 'list')
      newTodoTitle.value = ''
      return
    }

    if (text.startsWith('/')) {
      newTodoTitle.value = ''
      return
    }

    if (props.currentMode === 'history') {
      newTodoTitle.value = ''
      return
    }
    
    emit('add', newTodoTitle.value)
    newTodoTitle.value = ''
}

const focusInput = () => {
  inputRef.value?.focus()
}

</script>

<template>
    <form @submit.prevent="handleSubmit" @click="focusInput" class="flex items-center border-b border-terminal-dim/50 pb-4 cursor-text group">
    
    <span class="term-prompt">{{ props.currentMode === 'history' ? 'History ' : 'Todo ' }}></span>
    
    <!-- The Display Wrapper -->
    <div class="flex-1 font-mono relative flex items-center">    
          
      <!-- Placeholder (Only shows if input is empty) -->
      <span v-if="!newTodoTitle" class="text-terminal-dim/50 absolute pointer-events-none">
        enter command or task
      </span>
      
      <!-- What the user is actually typing -->
      <!-- 'whitespace-pre' is crucial here so spaces render correctly -->
      <span class="text-terminal-glow whitespace-pre pointer-events-none">{{ newTodoTitle }}</span>
      
      <!-- Blinking Terminal Cursor -->
      <span class="text-terminal-glow font-bold ml-[2px] pointer-events-none animate-[terminal-blink_1s_steps(2,start)_infinite]">
        _
      </span>
      
      <!-- The REAL Input (Completely invisible but doing all the work) -->
      <input 
        ref="inputRef"
        v-model="newTodoTitle"
        type="text" 
        class="absolute inset-0 w-full h-full bg-transparent text-transparent caret-transparent border-none outline-none focus:outline-none focus:ring-0 z-0 selection:bg-transparent selection:text-transparent"
        autocomplete="off"
        spellcheck="false"
        data-form-type="other"
    >
    </div>
    
    <button 
      type="submit" 
      class="ml-4 px-3 py-1 border border-terminal-dim text-terminal-glow group-hover:bg-terminal-glow group-hover:text-terminal-bg transition-colors font-bold text-sm z-10"
    >
      EXECUTE
    </button>
  </form>
  </template>