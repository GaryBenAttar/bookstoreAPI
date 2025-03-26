package main

import (
	"log"
	"net/http"
	"os"

	"github.com/garyben/booktore/api"
)

func main() {
	store := api.NewBookStore()

	store.CreateBook(api.Book{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan & Brian W. Kernighan", Year: 2015})
	store.CreateBook(api.Book{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008})
	store.CreateBook(api.Book{ID: "3", Title: "The Pragmatic Programmer", Author: "Andrew Hunt & David Thomas", Year: 1999})

	handler := api.NewBookHandler(store)
	router := api.SetupRouter(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default port is %s", port)
	}

	log.Printf("Server Starting on port: %s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
