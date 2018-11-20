package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/heppu/todo"
)

func main() {
	todoList := todo.NewList()

	server := http.Server{
		Addr:    ":8000",
		Handler: todo.NewHandler(todoList),
	}

	unexpectedError := make(chan error)

	log.Println("Listening at :8000")
	go func() {
		unexpectedError <- server.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-interrupt:
		log.Println("Shutting down server")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Failed to shutdown server", err)
		}

	case err := <-unexpectedError:
		log.Fatal("Server closed unexpectedly:", err)
	}

	log.Println("Server closed successfully, exiting now")
}
