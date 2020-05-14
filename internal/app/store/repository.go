package store

import "github.com/ythosa/winlock-server/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}

// CardRepository ...
type CardRepository interface {
	Create(*model.Card) error
	GetAll() ([]*model.Card, error)
}
