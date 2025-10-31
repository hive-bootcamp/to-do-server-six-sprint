package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hive-bootcamp/go-rest-api-homework/internal/model"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "ошибка преобразования JSON", http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		http.Error(w, "id не может быть пустым", http.StatusBadRequest)
		return
	}

	if _, exists := model.Tasks[task.ID]; exists {
		http.Error(w, "задача с таким id уже существует", http.StatusBadRequest)
		return
	}

	model.Tasks[task.ID] = task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Tasks[task.ID])
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task model.Task
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id не может быть пустым", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "ошибка преобразования JSON", http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		http.Error(w, "id не может быть пустым", http.StatusBadRequest)
		return
	}

	if _, exists := model.Tasks[id]; !exists {
		http.Error(w, "задача с таким id не существует", http.StatusNotFound)
		return
	}

	model.Tasks[id] = task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Tasks[task.ID])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id не может быть пустым", http.StatusBadRequest)
		return
	}

	if _, exists := model.Tasks[id]; !exists {
		http.Error(w, "задача с таким id не найдена", http.StatusNotFound)
		return
	}

	delete(model.Tasks, id)

	w.WriteHeader(http.StatusNoContent)
}
