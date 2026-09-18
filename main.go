// main.go
package main

import (
    "database/sql"
    "embed"
    "fmt"
    "io/fs"
    "log"
    "net/http"
	"os"

	"github.com/joho/godotenv"
    _ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed frontend/dist
var frontendFS embed.FS

// enableCORS intercepts every request to tell the browser it is safe
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        // Browsers send a pre-flight OPTIONS request before writing/deleting data
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    // 1. Connect to the PostgreSQL database
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: No .env file found")
    }

    // Grab the connection string from the environment
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL environment variable is missing!")
    }
    
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("Failed to open database: %v\n", err)
    }
    defer db.Close()

    // 2. Ping the database to ensure the connection actually works
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v\n", err)
    }
    fmt.Println("Successfully connected to the database!")

    // 3. Initialize our DAO
    dao := NewPGTodoDao(db)

	// 2. Pass the database to our HTTP handler
    handler := NewTodoHandler(dao)

    // 3. Set up our routes using Go's modern standard library syntax
    mux := http.NewServeMux()
    mux.HandleFunc("GET /todos", handler.GetAll)
    mux.HandleFunc("GET /todos/history", handler.GetHistory)
    mux.HandleFunc("POST /todos", handler.Create)
    mux.HandleFunc("PUT /todos/{id}", handler.Update)
    mux.HandleFunc("DELETE /todos/{id}", handler.Delete)
	mux.HandleFunc("GET /todos/{id}", handler.Get)
    mux.HandleFunc("PATCH /todos/{id}/done", handler.Done)

    distFS, err := fs.Sub(frontendFS, "frontend/dist")
    if err != nil {
        log.Fatalf("Failed to load frontend static files: %v", err)
    }

    fmt.Println("--- EMBEDDED FILES ---")
    fs.WalkDir(distFS, ".", func(path string, d fs.DirEntry, err error) error {
        fmt.Println("-", path)
        return nil
    })
    fmt.Println("----------------------")

    mux.Handle("/", http.FileServer(http.FS(distFS)))

    port := ":8080"
    fmt.Printf("Monolith running on http://localhost%s\n", port)
    
    // 4. Wrap the router in our CORS middleware and start the server
    err = http.ListenAndServe(port, enableCORS(mux))
    if err != nil {
        log.Fatalf("Server failed to start: %v\n", err)
    }
}