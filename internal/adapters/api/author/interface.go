package author

import (
	"context"

	"github.com/Beknur1003/clean-architecture.git/internal/domain/author"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *author.Author
	GetAll(ctx context.Context, limit, offset int) []*author.Author
	Create(ctx context.Context, dto *author.CreateAuthorDTO) *author.Author
}
