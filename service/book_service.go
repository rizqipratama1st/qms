package service

import (
	"fmt"

	"qms/models"
	"qms/repository"
)

type BookService struct {
	bookRepo repository.BookRepoInterface
}

func NewBookService(bookRepo repository.BookRepoInterface) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

type BookServiceInterface interface {
	CreateBook(book models.SlotBooking) (models.SlotBooking, error)
	GetBank() ([]models.Bank, error)
}

func (s *BookService) CreateBook(book models.SlotBooking) (models.SlotBooking, error) {

	response, err := s.bookRepo.CreateBook(book)
	fmt.Println(response)

	return response, err
}

func (s *BookService) GetBank() ([]models.Bank, error) {
	banks, err := s.bookRepo.FindAllBank()

	return banks, err
}
