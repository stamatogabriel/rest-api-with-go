package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/services"

	"github.com/go-chi/chi/v5"
)

func createTodoHandler(w http.ResponseWriter, req *http.Request) {
	err := json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Error decoding request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = todo.InsertTodo(todo)

	if err != nil {
		errResp := Response{
			Msg:  "Failed to create todo",
			Code: 304,
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	res := Response{
		Msg:  "Successfully created todo",
		Code: 201,
	}

	jsonString, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonString)
}

func getTodosHandler(w http.ResponseWriter, req *http.Request) {
	todos, err := todo.GetTodos()
	if err != nil {
		log.Printf("Error getting todos: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if len(todos) == 0 {
		res := Response{
			Msg:  "No todos found",
			Code: 404,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodoByIDHandler(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	todo, err := todo.GetTodoByID(id)
	if err != nil {
		log.Printf("Error getting todo by ID: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if todo == nil {
		res := Response{
			Msg:  "Todo not found",
			Code: 404,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func updateTodoHandler(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var updatedTodo services.Todo
	err := json.NewDecoder(req.Body).Decode(&updatedTodo)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = todo.UpdateTodo(id, updatedTodo)
	if err != nil {
		log.Printf("Error updating todo: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := Response{
		Msg:  "Successfully updated todo",
		Code: 200,
	}
	json.NewEncoder(w).Encode(res)
}

func deleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := todo.DeleteTodo(id)
	if err != nil {
		log.Printf("Error deleting todo: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := Response{
		Msg:  "Successfully deleted todo",
		Code: 200,
	}
	json.NewEncoder(w).Encode(res)
}
