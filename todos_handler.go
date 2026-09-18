// todos_handler.go
package main

import (
    "encoding/json"
    "net/http"
)

// TodoHandler holds our database interface
type TodoHandler struct {
    dao TodoDao
}

func NewTodoHandler(dao TodoDao) *TodoHandler {
    return &TodoHandler{dao: dao}
}

// GET /todos
func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    todos, err := h.dao.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
    todos, err := h.dao.GetHistory()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}
// POST /todos
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    // Read the incoming JSON from Vue and map it to our struct
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := h.dao.Create(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

// PUT /todos/{id}
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := h.dao.Update(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// DELETE /todos/{id}
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
    // Go 1.22+ lets us grab the ID directly from the URL path!
    id := r.PathValue("id")
    
    if err := h.dao.Delete(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

// GET /todos/{id}
func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    
    todo, err := h.dao.Get(id)
    if err != nil {
        // If the database can't find it, return a 404 Not Found
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// PATCH /todos/{id}/done
func (h *TodoHandler) Done(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    
    if err := h.dao.Done(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // StatusNoContent (204) is standard for successful updates that don't return data
    w.WriteHeader(http.StatusNoContent)
}