// cmd/main.go
package main

import (
	"fmt"
	"hw/configs"
	"hw/internal/todo"
	"hw/pkg/db"
	"net/http"
)

func main() {
	conf := configs.DefaultConfig()
	pgDb := db.NewDb(conf)

	fmt.Println(conf.Db.Dsn)

	router := http.NewServeMux()

	taskRepository := todo.NewTaskRepository(pgDb)

	todo.NewTaskHandler(router, taskRepository)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
