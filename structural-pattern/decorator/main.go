package main

import (
	"fmt"
	"net/http"
)

type Handler func(r *http.Request) (int, string)

type Middleware func(h Handler) Handler

func LoggingMiddleware(h Handler) Handler {
	return func(r *http.Request) (int, string) {
		fmt.Println("Handling request")
		status, body := h(r)
		fmt.Printf("Request handled with status %d\n", status)
		return status, body
	}
}

func main() {
	handler := func(r *http.Request) (int, string) {
		return http.StatusOK, "Hello, World!"
	}

	logginHandler := LoggingMiddleware(handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status, body := logginHandler(r)
		w.WriteHeader(status)
		w.Write([]byte(body))
	})

	http.ListenAndServe(":8080", nil)
}
