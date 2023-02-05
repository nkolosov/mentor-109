package entity

import (
	"github.com/google/uuid"
	"time"
)

type CategoryId uuid.UUID

type Category struct {
	id               CategoryId `db:"id"`
	name             string     `db:"name"`
	createDate       time.Time  `db:"created_at"`
	modificationDate time.Time  `db:"deleted_at"`
	deleteDate       time.Time  `db:"updated_at"`
}

func NewCategory(
	id CategoryId,
	name string,
	createDate time.Time,
	modificationDate time.Time,
	deleteDate time.Time,
) *Category {
	return &Category{
		id:               id,
		name:             name,
		createDate:       createDate,
		modificationDate: modificationDate,
		deleteDate:       deleteDate,
	}
}

func (c *Category) Id() CategoryId {
	return c.id
}

func (c *Category) SetId(id CategoryId) {
	c.id = id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) SetName(name string) {
	c.name = name
}

func (c *Category) CreateDate() time.Time {
	return c.createDate
}

func (c *Category) SetCreateDate(createDate time.Time) {
	c.createDate = createDate
}

func (c *Category) ModificationDate() time.Time {
	return c.modificationDate
}

func (c *Category) SetModificationDate(modificationDate time.Time) {
	c.modificationDate = modificationDate
}

func (c *Category) DeleteDate() time.Time {
	return c.deleteDate
}

func (c *Category) SetDeleteDate(deleteDate time.Time) {
	c.deleteDate = deleteDate
}
