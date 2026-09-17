import {
    defineConfig,
    presetIcons,
    presetTypography,
    presetUno,
    transformerDirectives,
  } from 'unocss'
  import { presetForms } from '@julr/unocss-preset-forms'
  
  export default defineConfig({
    presets: [
      presetUno(),
      presetIcons(),
      presetTypography(),
      presetForms(),
    ],
    transformers: [
      transformerDirectives(),
    ],
    theme: {
      colors: {
        // Classic matrix green or amber CRT phosphor shades
        terminal: {
          bg: '#0c0f12',     // Deep abyss dark background
          glow: '#4af626',   // Bright phosphor green
          dim: '#1b6b10',    // Muted green for borders/metadata
          text: '#d8f3dc'    // Soft green-white for high readability
        }
      },
      fontFamily: {
        mono: 'Fira Code, JetBrains Mono, Courier New, monospace'
      },
      animation: {
        keyframes: {
          'terminal-blink': '{ 0%, 100% { opacity: 1 } 50% { opacity: 0 } }'
        },
        durations: {
          'terminal-blink': '1s'
        },
        timingFns: {
          'terminal-blink': 'steps(2, start)'
        }
      }
    },
    shortcuts: {
      // Reusable UI components for your terminal look
      'term-window': 'bg-terminal-bg border border-terminal-dim rounded-md shadow-[0_0_20px_rgba(74,246,38,0.15)] font-mono text-terminal-text overflow-hidden',
      'term-header': 'bg-terminal-dim/20 border-b border-terminal-dim/40 px-4 py-2 flex items-center justify-between text-xs select-none',
      'term-body': 'p-4 space-y-2 text-sm leading-relaxed md:text-base',
      'term-prompt': 'text-terminal-glow font-bold mr-2 select-none'
    }
  })