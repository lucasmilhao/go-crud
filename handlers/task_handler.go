package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucasmilhao/go-crud/models"
	"github.com/lucasmilhao/go-crud/repository"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (th *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	tasks, err := th.repo.FindAll()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (th *TaskHandler) GetTask(writer http.ResponseWriter, request *http.Request) {
	id, err := parseId(request)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	task, err := th.repo.FindById(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task)
}

func (th *TaskHandler) CreateTasks(writer http.ResponseWriter, request *http.Request) {
	var task models.Task

	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.repo.Save(task); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (th *TaskHandler) UpdateTasks(writer http.ResponseWriter, request *http.Request) {
	id, err := parseId(request)

	if err != nil {
		http.Error(writer, "ID de task inválido", http.StatusBadRequest)
		return
	}

	var task models.Task

	err = json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = th.repo.Update(id, task); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (th *TaskHandler) DeleteTasks(writer http.ResponseWriter, request *http.Request) {
	id, err := parseId(request)

	if err != nil {
		http.Error(writer, "ID de task inválido", http.StatusBadRequest)
		return
	}

	isMoreThanZeroRows, err := th.repo.Delete(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if isMoreThanZeroRows {
		http.Error(writer, "No task found with this id", http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func parseId(request *http.Request) (int, error) {
	vars := mux.Vars(request)
	return strconv.Atoi(vars["id"])
}
