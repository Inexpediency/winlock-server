package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ythosa/go-rest-api-server/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "not valid all props",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "not valid password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "11"

				return u
			},
			isValid: false,
		},
		{
			name: "not valid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "11"

				return u
			},
			isValid: false,
		},
		{
			name: "not valid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "asda@gg"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
