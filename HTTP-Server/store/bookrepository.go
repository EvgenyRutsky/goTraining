package store

import "server/model"

type Repository interface {
	Create(b *model.Book) (interface{}, error)
}

