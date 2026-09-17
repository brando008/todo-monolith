# AI Agent Instructions: Todo App (Frontend)

## Project Context
This repository contains the frontend client for a Todo application. It is currently built as a standalone Single Page Application (SPA) but is architected with the intent to eventually merge into a monolith alongside a Go/PostgreSQL backend. 

All AI generations, refactoring, and terminal commands must align with a Windows 10 development environment.

## Technology Stack
*   **Framework:** Vue 3 (Composition API)
*   **Language:** TypeScript
*   **Build Tool:** Vite
*   **Package Manager & Runtime:** Bun
*   **Styling:** UnoCSS
*   **Linting/Formatting:** ESLint (v9 Flat Config) & Prettier

## Architectural Guidelines & Future-Proofing
1.  **API Isolation:** All external data fetching must be abstracted into a dedicated service layer (e.g., `src/services/api.ts`). Do not make raw `fetch` or `axios` calls directly inside Vue components. This ensures a seamless transition when the Go backend API is integrated.
2.  **Environment Variables:** Use Vite's `import.meta.env` for API base URLs to easily swap between local development mocks and the future Go backend server.
3.  **State Management:** Keep local component state minimal. If global state is required for todos, utilize Vue's reactivity system (composables) or Pinia.

## Coding Standards

### Vue & TypeScript
*   Strictly use the **Composition API** with `<script setup lang="ts">`.
*   Avoid the Options API entirely.
*   Define explicit TypeScript `interfaces` or `types` for all data models, particularly the `Todo` object (e.g., `id`, `title`, `completed`, `created_at`, `updated_at`) matching the future PostgreSQL schema.
*   Avoid `any` types; utilize strict typing for props, emits, and return values.

### Styling (UnoCSS)
*   Rely exclusively on UnoCSS utility classes for styling. 
*   Avoid writing custom CSS in `<style>` blocks unless absolutely necessary for complex animations or hyper-specific edge cases.
*   Utilize the configured presets (Forms, Typography, Icons) appropriately.

### Linting & Formatting
*   Adhere to the ESLint flat config rules defined in `eslint.config.js`. 
*   Note: `vue/multi-word-component-names` is explicitly turned off for this project.
*   Rely on Prettier for all formatting rules (2-space indent, no semicolons, single quotes).

## AI Workflow Directives
*   **Package Management:** Always use `bun add` or `bun add -D` for installing dependencies. Never use `npm`, `yarn`, or `pnpm`.
*   **Command Execution:** Provide terminal commands formatted for Windows 10 (PowerShell/CMD compatible).
*   **Component Generation:** When generating new Vue components, scaffold them with `<script setup lang="ts">`, a `<template>`, and omit the `<style>` block unless specifically requested.