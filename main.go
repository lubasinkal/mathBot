package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		userText := r.URL.Query().Get("msg")
		botResponse := getBotResponse(userText)
		fmt.Fprint(w, botResponse)
	})

	fmt.Println("Starting server on http://localhost:8080")
	router := http.DefaultServeMux

	server := http.Server{
		Addr:    ":8080",
		Handler: Logging(router),
	}
	log.Printf("Server running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
