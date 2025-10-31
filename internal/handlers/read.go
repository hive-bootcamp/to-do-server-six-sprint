package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hive-bootcamp/go-rest-api-homework/internal/model"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(model.Tasks)
	if err != nil {
		http.Error(w, "ошибка в получении задач", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id не может быть пустым", http.StatusBadRequest)
		return
	}

	task, ok := model.Tasks[id]
	if !ok {
		http.Error(w, "задача не найдена", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "ошибка в получении задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
