package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/winlock-server/internal/app/model"
	"github.com/ythosa/winlock-server/internal/app/store/sqlstore"
)

func TestCardRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("cards")

	s := sqlstore.New(db)
	c := model.TestCard(t)

	assert.NoError(t, s.Card().Create(c))
	assert.NotNil(t, c.ID)
}

func TestCardRepository_GetAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("cards")

	s := sqlstore.New(db)
	_, err := s.Card().GetAll()
	assert.NoError(t, err)
}
