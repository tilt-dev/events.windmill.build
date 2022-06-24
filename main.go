package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var okHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
})

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(okHandler)
	return &Router{
		Router: r,
	}
}

func main() {
	http.Handle("/", NewRouter())

	log.Println("Serving on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
