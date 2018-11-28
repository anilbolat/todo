package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/heppu/todo/api"
	"github.com/heppu/todo/mem"
)

func main() {
	todoList := mem.NewList()

	server := http.Server{
		Addr:    ":8000",
		Handler: api.NewHandler(todoList),
	}

	unexpectedError := make(chan error)

	log.Println("Listening at :8000")
	go func() {
		unexpectedError <- server.ListenAndServe()
	}()

	go func() { log.Println(http.ListenAndServe(":6060", nil)) }()

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
