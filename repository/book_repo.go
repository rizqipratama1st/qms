package repository

import (
	"qms/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

type BookRepoInterface interface {
	CreateBook(book models.SlotBooking) (models.SlotBooking, error)
	FindAllBank() ([]models.Bank, error)
}

func (r *BookRepository) CreateBook(book models.SlotBooking) (models.SlotBooking, error) {
	trx := r.db.Create(&book)

	return book, trx.Error
}

func (r *BookRepository) FindAllBank() ([]models.Bank, error) {
	var banks []models.Bank
	findResult := r.db.Find(&banks)

	return banks, findResult.Error
}
