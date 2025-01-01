// cmd/main.go
package main

import (
	"fmt"
	"hw/internal/todo"
	"hw/pkg/db"
	"net/http"
)

func main() {
	db := db.NewDb()

	router := http.NewServeMux()

	taskRepository := todo.NewTaskRepository(db)

	todo.NewTaskHandler(router, taskRepository)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
