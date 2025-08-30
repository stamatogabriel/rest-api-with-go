package handlers

import (
	"backend/middlewares"
	"encoding/json"
	"net/http"
)

// Exemplo de como usar o user_id em um handler protegido
func createTodoProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Extrair o user_id do context
	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	// Agora você pode usar o userID para criar todos específicos do usuário
	var todo struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
		UserID      string `json:"user_id"`
	}

	// Decodificar o corpo da requisição
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Definir o user_id do token JWT
	todo.UserID = userID

	// Aqui você salvaria o todo no banco com o user_id
	// err = todoService.CreateTodo(todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo created successfully",
		"user_id": userID,
		"todo":    todo,
	})
}
