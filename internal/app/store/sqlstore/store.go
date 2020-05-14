package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/ythosa/winlock-server/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	cardRepository *CardRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Card ...
func (s *Store) Card() store.CardRepository {
	if s.cardRepository != nil {
		return s.cardRepository
	}

	s.cardRepository = &CardRepository{
		store: s,
	}

	return s.cardRepository
}
