package books

import (
	"gorm.io/gorm"
	"log"
)

// struct for books service.
type Service struct {
	DB *gorm.DB
}

//
type Book struct {
	gorm.Model
	Author string
	Title  string
	Price  float64
}

type BooksService interface {
	GetBook(ID uint) (Book, error)
	GetBooksByAuthor(author string) ([]Book, error)
	CreateBook(book Book) (Book, error)
	UpdateBook(ID uint, book Book) error
	DeleteBook(ID uint) error
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// Get book by given ID. ID will be primary key of Book table.
func (s *Service) GetBoook(ID uint) (*Book, error) {
	log.Printf("Fetching book with id: %d\n", ID)
	var book Book
	// ID is primary key.
	// First parameter is variable where you want to assign the results and second parameter is ID which will match with primary key.
	if r := s.DB.First(&book, ID); r.Error != nil {
		return nil, r.Error
	}
	return &book, nil
}

// Get List of Books by author name.
func (s *Service) GetBooksByAuthor(author string) ([]Book, error) {
	log.Println("Fetching books by author: %s", author)
	var books []Book
	if r := s.DB.Find(&books).Where("author = ?", author); r.Error != nil {
		return nil, r.Error
	}
	return books, nil
}

// Create a new book entry in the database.
func (s *Service) CreateBook(book Book) (*Book, error) {
	log.Println("Creating Book: %+v", book)
	if r := s.DB.Save(&book); r.Error != nil {
		return nil, r.Error
	}
	return &book, nil
}

// Update the existing book given its ID.
func (s *Service) Updatebook(ID uint, book Book) error {
	log.Println("Updating book: %+v", book)
	b, err := s.GetBoook(ID)
	if err != nil {
		return err
	}
	if r := s.DB.Model(&b).Updates(book); r.Error != nil {
		return r.Error
	}

	return nil
}

// Delete the Book given its ID.
func (s *Service) DeleteBook(ID uint) error {
	log.Println("Deleting book having id: %d", ID)
	b, err := s.GetBoook(ID)
	if err != nil {
		return err
	}
	log.Println("Deleting book: %+v", b)
	if r := s.DB.Delete(&Book{}, ID); r.Error != nil {
		return r.Error
	}
	return nil
}
