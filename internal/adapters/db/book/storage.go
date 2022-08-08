package book

import (
	"github.com/Beknur1003/clean-architecture.git/internal/domain/book"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db *mongo.Database
}

// Конструктор NewStorage()

func NewStorage(db *mongo.Database) book.Storage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book.Book) error {
	return nil
}
