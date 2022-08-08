package composites

import (
	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api"
	book2 "github.com/Beknur1003/clean-architecture.git/internal/adapters/api/book"
	book3 "github.com/Beknur1003/clean-architecture.git/internal/adapters/db/book"
	"github.com/Beknur1003/clean-architecture.git/internal/domain/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongoComposite.db)
	service := book.NewService(storage)
	handler := book2.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
