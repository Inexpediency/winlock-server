package model

import (
	"errors"
	"regexp"
	"strconv"
)

// Card ...
type Card struct {
	ID     int    `json:"id"`
	Digits string `json:"digits"`
	Cvv    int    `json:"cvv"`
	Date   string `json:"date"`
	Owner  string `json:"owner"`
}

func luhnAlgoritm(digits string) bool {
	sum := 0
	for i := 0; i < len(digits); i++ {
		cardNum, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			return false
		}

		if (len(digits)-i)%2 == 0 {
			cardNum *= 2

			if cardNum > 9 {
				cardNum -= 9
			}
		}

		sum += int(cardNum)
	}

	return sum%10 == 0
}

// Validate ...
func (c *Card) Validate() error {
	ownerValidate := regexp.MustCompile(`^\w+ \w+$`)
	if !ownerValidate.Match([]byte(c.Owner)) {
		return errors.New("not valid card owner")
	}

	cvvValidate := regexp.MustCompile(`^\d{3}$`)
	if !cvvValidate.Match([]byte(strconv.Itoa(c.Cvv))) {
		return errors.New("not valid card cvv code")
	}

	dateValidate := regexp.MustCompile(`^\d{1,2}/\d{2}$`)
	if !dateValidate.Match([]byte(c.Date)) {
		return errors.New("not valid card expiration date")
	}

	if !luhnAlgoritm(c.Digits) {
		return errors.New("not valid card digits")
	}

	return nil
}
