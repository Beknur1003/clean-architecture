package author

import (
	"github.com/Beknur1003/clean-architecture.git/internal/domain/author"
	"go.mongodb.org/mongo-driver/mongo"
)

// db это не наше приложение, это просто слой доступа к данным. Это паттерн ДАО, паттерн репозиторий. Они импортируют данные из domain

type authorStorage struct {
	db *mongo.Database
}

// Конструктор NewStorage()

func NewStorage(db *mongo.Database) author.Storage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (bs *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}
func (bs *authorStorage) Create(book *author.Author) *author.Author {
	return nil
}
func (bs *authorStorage) Delete(book *author.Author) error {
	return nil
}
