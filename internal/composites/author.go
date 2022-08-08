package composites

import (
	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api"
	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api/author"
	author3 "github.com/Beknur1003/clean-architecture.git/internal/adapters/db/author"
	author2 "github.com/Beknur1003/clean-architecture.git/internal/domain/author"
)

// Композиты это вынос всех созданий обьектов, адаптеров, всех внутренних слоев, внешниъ слоев в отдельную сущность. Так называемые композиты

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongoComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
