package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucasmilhao/go-crud/models"
	"github.com/lucasmilhao/go-crud/service"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

const table = "tasks"

func (th *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	rows, err := service.FindAll(th.DB, table)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		tasks = append(tasks, task)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (th *TaskHandler) CreateTasks(writer http.ResponseWriter, request *http.Request) {
	var task models.Task

	err := json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = service.Save(th.DB, table, []string{"title", "description", "status"}, []any{task.Title, task.Description, task.Status})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (th *TaskHandler) UpdateTasks(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

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

	_, err = th.DB.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?", task.Title, task.Description, task.Status, id)

	rows, err := service.FindById(th.DB, table, id)

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		tasks = append(tasks, task)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (th *TaskHandler) DeleteTasks(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(writer, "ID de task inválido", http.StatusBadRequest)
		return
	}

	result, err := service.DeleteById(th.DB, table, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "No task found with this id", http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
