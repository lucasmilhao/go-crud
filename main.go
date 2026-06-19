package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmilhao/go-crud/config"
	"github.com/lucasmilhao/go-crud/handlers"
	"github.com/lucasmilhao/go-crud/models"
	"github.com/lucasmilhao/go-crud/repository"
)

func main() {
	dbConn := config.SetupDB()

	criar, err := dbConn.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(criar)

	defer dbConn.Close()

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(repository.NewTaskRepository(dbConn))

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")

	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")

	router.HandleFunc("/tasks", taskHandler.CreateTasks).Methods("POST")

	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTasks).Methods("DELETE")

	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTasks).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
