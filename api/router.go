package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetupRouter(handler *BookHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Api is healthy"))
	}).Methods("GET")

	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/books", handler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("PUT")

	router.Use(loggingMiddleware)

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
