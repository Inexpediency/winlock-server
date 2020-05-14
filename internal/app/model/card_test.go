package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/winlock-server/internal/app/model"
)

func TestCard_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		c       func() *model.Card
		isValid bool
	}{
		{
			name: "valid",
			c: func() *model.Card {
				return model.TestCard(t)
			},
			isValid: true,
		},
		{
			name: "not valid card number",
			c: func() *model.Card {
				c := model.TestCard(t)
				c.Digits = "12312"
				return c
			},
			isValid: false,
		},
		{
			name: "not valid card date",
			c: func() *model.Card {
				c := model.TestCard(t)
				c.Date = "1111/12"
				return c
			},
			isValid: false,
		},
		{
			name: "not valid card owner",
			c: func() *model.Card {
				c := model.TestCard(t)
				c.Owner = "asdasdas"
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.c().Validate())
			} else {
				assert.Error(t, tc.c().Validate())
			}
		})
	}
}
