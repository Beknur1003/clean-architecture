package book

import (
	"context"
	"fmt"

	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api/author"
	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api/book"
)

// Это реализация бизнес логики, ее интерфейс лежит там, где он используется (api/book/)
// Тут наша логика, наше приложение. Мы импортируем из адаптеров только интерфейс, которые адаптеры просят реализовать
// То есть связь только через интерфейс

type service struct {
	storage       Storage
	authorService author.Service // Связь между book service и author service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateBookDTO) *Book {
	author := s.authorService.GetByUUID(ctx, dto.AuthorUUID) // Проверка есть ли вообще такой автор
	fmt.Println(author)
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
