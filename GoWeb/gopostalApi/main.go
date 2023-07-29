package main

import (
	"flag"
	"fmt"
	"log"

	"com.loollab/postalapi/api"
	"com.loollab/postalapi/storage"
	// "net/http"
)

func main() {

	listenAddr := flag.String("listenAddr", "3000", "the Server Address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server runs on", *listenAddr)
	log.Fatal(server.Start())

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })

	// http.ListenAndServe(":8080", nil)
}
