# Todo Monolith

A full-stack todo application with a Go backend and Vue 3 frontend.
Ability to togle between Todo and History by typing "/todo" | "/history"

## What's Happening?
Following the responses from the start of the HTML, through the Headers, and finally to the database, we're able to understand this deployment.
The responses from the Vue app are captured and sent through the todo.service.ts, which sends out a header for the appropriate request. This could be
a delete, add, or completed request. Main.go initializes the connection between the database and http handlers. The HTTP handler sees the request, 
adds the information to a shared object, and lets the DAO know which operation to perform. The DAO then performs the operation on the database and
passes the results back to the handler, which in turn responds to the Vue frontend.

## Installation
Setup SQL database
```bash
CREATE DATABASE todo_db;
/c todo_db
CREATE TABLE todos (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```
Move to folder | setup .env
```bash
cd todo-monolith
go mod tidy
cd frontend
bun run build
```
Compiles
```bash
cd ..
go build -a -o todo-app .
./todo-app
```