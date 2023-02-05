package entity

import (
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
