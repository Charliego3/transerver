package commons

import "context"

type Repository[T any] interface {
	FindById(ctx context.Context, uid int64) (*T, error)
	FindList(ctx context.Context, t *T) ([]T, error)
	Save(ctx context.Context, t *T) error
	Delete(ctx context.Context, uid int64) error
}
