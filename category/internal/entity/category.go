package entity

import "time"

type Category struct {
	Id               int
	Name             string
	CreateDate       time.Time
	ModificationDate time.Time
	DeleteDate       time.Time
	IsDelete         bool
}
