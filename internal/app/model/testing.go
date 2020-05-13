package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "hacker@hacker.hack",
		Password: "anonym",
	}
}

// TestCard ...
func TestCard(t *testing.T) *Card {
	return &Card{
		Digits: "4111111111111111",
		Cvv:    123,
		Date:   "1/23",
		Owner:  "Sername Name",
	}
}
