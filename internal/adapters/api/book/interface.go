package book

import (
	"context"

	"github.com/Beknur1003/clean-architecture.git/internal/domain/book"
)

// Лежит тут потому что используется в handler-controllers. В чистой архитектуре интерфейс сервиса(интерфейс use case) это
// доступ контролеру к приложению. По сути наше приложение это папка domain. Domain это бизнес логика нашего приложения
// Интерфейс, который тут лежит он служит мостиком/доступом в гексогональной архитектуре это как раз порт. Этот интерфейс явлется портом доступа в бизнес логику/приложению

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetAll(ctx context.Context, limit, offset int) []*book.Book
	Create(ctx context.Context, dto *book.CreateBookDTO) *book.Book
}
