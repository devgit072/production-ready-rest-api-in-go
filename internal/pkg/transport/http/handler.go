package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateRoutes() {
	log.Println("Creating Routes for various end points")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, e := fmt.Fprintf(w, "Pong")
		if e != nil {
			log.Fatalf("Error: %s", e.Error())
		}
	})
}
