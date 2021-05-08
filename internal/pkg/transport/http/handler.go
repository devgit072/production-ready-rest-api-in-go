package http

import (
	"encoding/json"
	"fmt"
	"github.com/devgit072/production-ready-rest-api-in-go/internal/books"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Router *mux.Router
	Service *books.Service
}

// struct for rest api response.
type ApiResponse struct {
	Msg string
}

func NewHandler(service *books.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateRoutes() {
	log.Println("Creating Routes for various end points")
	h.Router = mux.NewRouter()
	// Books service crud api
	h.Router.HandleFunc("/books/{id}", h.GetBook).Methods("GET")
	h.Router.HandleFunc("/books", h.CreateBook).Methods("POST")
	h.Router.HandleFunc("/books", h.UpdateBook).Methods("PUT")
	h.Router.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")
	h.Router.HandleFunc("/ping", h.Ping).Methods("GET")
}

// Ping method to check if service is up or not.
func (h *Handler) Ping (w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	if err := json.NewEncoder(w).Encode(ApiResponse{Msg: "pong"}); err != nil {
		log.Fatalf("Service is down. Error: %s", err.Error())
	}
}
// Rest api implementaion for get book by its ID.
func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	v := mux.Vars(r)
	idStr := v["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Error while parsing id: %s", err.Error())
		return
	}
	book, err := h.Service.GetBoook(uint(id))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		fmt.Fprintf(w, "Error while parsing into json. Error:%s", err.Error())
	}
}

// Rest api implementation for creating new book entry.
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	book, err := h.Service.CreateBook(books.Book{
		Author: "Java",
		Title:  "Goetz",
		Price:  340,
	})
	if err != nil {
		fmt.Fprintf(w, "Error while creating new book: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "%+v", book)
}

// Rest api implementation for updating existing book with new values
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	err := h.Service.Updatebook(1, books.Book{
		Author: "Java-new",
		Title:  "Goetz-new",
		Price:  248,
	})
	if err != nil {
		fmt.Fprintf(w, "Error while update book: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "Book with id updated successfully ")
}

// Rest api implementation for deleting book by its given id.
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	v := mux.Vars(r)
	idStr := v["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "Error while parsing id:%s", idStr)
	}
	err = h.Service.DeleteBook(uint(id))
	if err != nil {
		fmt.Fprintf(w, "Error while book: %s", err.Error())
	}
	fmt.Fprintf(w, "Book with id: %d deleted successfully", id)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
