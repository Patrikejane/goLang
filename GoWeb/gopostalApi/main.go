package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"com.loollab/postalapi/api"
	"com.loollab/postalapi/storage"
	// "net/http"
)

func main() {

	// Create a channel to receive OS signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	listenAddr := flag.String("listenAddr", "3000", "the Server Address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server runs on", *listenAddr)
	log.Fatal(server.Start())

	fmt.Println("Server is running. Press Ctrl + C to stop.")

	// Block the main goroutine until a signal is received.
	<-sigChan
	// Perform any cleanup or graceful shutdown here before exiting.
	fmt.Println("\nServer is shutting down...")
	// For example, if you have a HTTP server:
	// stopHTTPServer()

	fmt.Println("Goodbye!")
}
