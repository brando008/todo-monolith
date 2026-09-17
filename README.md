# Todo Monolith

A full-stack todo application with a Go backend and Vue 3 frontend.

## Structure

```
├── main.go          # Go backend entry point
├── frontend/        # Vue 3 frontend
└── go.mod           # Go module
```

## Backend

Go backend using pgx for PostgreSQL.

```bash
go run main.go
```

## Frontend

Vue 3 + TypeScript + Vite + UnoCSS

```bash
cd frontend
npm install
npm run dev
```
## What's Happening?
Following the responses from the start of the HTML, through the Headers, and finally to the database, we're able to understand this deployment.
The responses from the Vue app are captured and sent through the todo.service.ts, which sends out a header for the appropriate request. This could be
a delete, add, or completed request. Main.go initializes the connection between the database and http handlers. The HTTP handler sees the request, 
adds the information to a shared object, and lets the DAO know which operation to perform. The DAO then performs the operation on the database and
passes the results back to the handler, which in turn responds to the Vue frontend.