package teststore

import "github.com/ythosa/winlock-server/internal/app/model"

// CardRepository ...
type CardRepository struct {
	store *Store
	cards []*model.Card
}

// Create ...
func (r *CardRepository) Create(c *model.Card) error {
	if err := c.Validate(); err != nil {
		return err
	}

	c.ID = len(r.cards) + 1
	r.cards[c.ID] = c

	return nil
}

// GetAll ...
func (r *CardRepository) GetAll() ([]*model.Card, error) {
	return r.cards, nil
}
