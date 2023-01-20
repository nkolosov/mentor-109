package entity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type CategoryId uuid.UUID

type Category struct {
	Id               CategoryId
	Name             string
	CreateDate       time.Time
	ModificationDate time.Time
	DeleteDate       time.Time
}

type CategoryRepository interface {
	Create(ctx context.Context, id CategoryId, name string) (*Category, error)
	Update(ctx context.Context, id CategoryId, name string) (*Category, error)
	Delete(ctx context.Context, id CategoryId) error
	Filter(ctx context.Context, ids []CategoryId) ([]*Category, error)
}
