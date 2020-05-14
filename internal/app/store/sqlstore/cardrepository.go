package sqlstore

import (
	"github.com/ythosa/winlock-server/internal/app/model"
)

// CardRepository ...
type CardRepository struct {
	store *Store
}

// Create ...
func (r *CardRepository) Create(c *model.Card) error {
	if err := c.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO cards (digits, cvv, card_date, card_owner) VALUES ($1, $2, $3, $4) RETURNING id",
		c.Digits,
		c.Cvv,
		c.Date,
		c.Owner,
	).Scan(&c.ID)
}

// GetAll ...
func (r *CardRepository) GetAll() ([]*model.Card, error) {
	rows, err := r.store.db.Query("SELECT * FROM cards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cds := make([]*model.Card, 0)
	for rows.Next() {
		cd := new(model.Card)
		err := rows.Scan(&cd.ID, &cd.Digits, &cd.Cvv, &cd.Date, &cd.Owner)
		if err != nil {
			return nil, err
		}
		cds = append(cds, cd)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cds, nil
}
