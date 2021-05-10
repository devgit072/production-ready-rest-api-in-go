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
	Router  *mux.Router
	Service *books.Service
}

// struct for rest api response.
type ApiResponse struct {
	Msg   string
	Error string
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
	h.Router.HandleFunc("/books/{id}", h.UpdateBook).Methods("PUT")
	h.Router.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")
	h.Router.HandleFunc("/ping", h.Ping).Methods("GET")
}

// Ping method to check if service is up or not.
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
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
		displayError(w, fmt.Sprintf("Error while parsing id: %s", idStr), err)
		return
	}
	book, err := h.Service.GetBoook(uint(id))
	if err != nil {
		displayError(w, fmt.Sprintf("Error while fetching book for id: %d", id), err)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		displayError(w, fmt.Sprintf("Error while fetching book for id: %d", id), err)
	}
}

// Rest api implementation for creating new book entry.
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var bookParam books.Book
	if err := json.NewDecoder(r.Body).Decode(&bookParam); err != nil {
		displayError(w, fmt.Sprintf("Error while parsing post params for book"), err)
		return
	}
	book, err := h.Service.CreateBook(bookParam)
	if err != nil {
		displayError(w, fmt.Sprintf("Error while creating book: %+v", bookParam), err)
		return
	}
	fmt.Fprintf(w, "%+v", book)
}

// Rest api implementation for updating existing book with new values
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var bookParam books.Book
	v := mux.Vars(r)
	idStr := v["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		displayError(w, fmt.Sprintf("Error while parsing id: %s", idStr), err)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&bookParam); err != nil {
		displayError(w, fmt.Sprintf("Error while parsing json params: %+v", bookParam), err)
		return
	}
	if err := h.Service.Updatebook(uint(id), bookParam); err != nil {
		displayError(w, fmt.Sprintf("Error while updating book: %+v", bookParam), err)
	}
	fmt.Fprintf(w, "Book with id: %d updated successfully.", id)
}

// Rest api implementation for deleting book by its given id.
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	v := mux.Vars(r)
	idStr := v["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		displayError(w, fmt.Sprintf("Error while parsing id: %s", idStr), err)
		return
	}
	err = h.Service.DeleteBook(uint(id))
	if err != nil {
		displayError(w, fmt.Sprintf("Error while deleting book with id: %d", id), err)
		return
	}
	fmt.Fprintf(w, "Book with id: %d deleted successfully", id)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func displayError(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(ApiResponse{Msg: msg, Error: err.Error()}); err != nil {
		panic(err)
	}
}
