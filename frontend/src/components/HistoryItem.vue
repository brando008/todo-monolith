<template>
    <div class="font-mono text-green-700/80 mb-2 flex">
      <!-- The text that is currently being typed -->
      <span>{{ displayedText }}</span>
      
      <!-- The Matrix cursor that only blinks while this specific row is typing -->
      <span 
        v-if="isTyping" 
        class="text-green-500 animate-pulse ml-[2px]"
      >_</span>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import type { Todo } from '../services/todo.service' 
  
  const props = defineProps<{
    task: Todo
    index: number
  }>()
  
  const displayedText = ref('')
  const isTyping = ref(false)
  
  onMounted(() => {
    // 1. Format the string to look like a system log
    const dateStr = new Date(props.task.deleted_at || '').toLocaleDateString()
    const fullText = `[SYS.DEL - ${dateStr}] ${props.task.title}`
  
    // 2. Calculate the cascade delay (e.g., 500ms per row)
    const cascadeDelay = props.index * 500
  
    // 3. Wait for our turn, then start typing
    setTimeout(() => {
      isTyping.value = true
      let charIndex = 0
  
      // 4. Fire every 30ms to reveal the next character
      const typingInterval = setInterval(() => {
        displayedText.value += fullText.charAt(charIndex)
        charIndex++
  
        // 5. Stop the timer when we run out of letters
        if (charIndex >= fullText.length) {
          clearInterval(typingInterval)
          isTyping.value = false
        }
      }, 30)
  
    }, cascadeDelay)
  })
  </script>