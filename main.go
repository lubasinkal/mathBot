package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		userText := r.URL.Query().Get("msg")
		botResponse := getBotResponse(userText)
		fmt.Fprintf(w, botResponse)
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
