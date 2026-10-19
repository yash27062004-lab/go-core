package payment

import (
	"errors"
	"fmt"
)

// IMPLEMENTATION 1:
// Card payment follows the payment behavior contract.
type Card struct {
	LastFour string
}

func (c Card) Pay(amount int) error {
	if c.LastFour == "" {
		return errors.New("card details are required")
	}

	fmt.Printf("Card ending %s charged ₹%d\n", c.LastFour, amount)
	return nil
}

// IMPLEMENTATION 2:
// Wallet payment follows the same contract.
type Wallet struct {
	UserID string
}

func (w Wallet) Pay(amount int) error {
	if w.UserID == "" {
		return errors.New("wallet user ID is required")
	}

	fmt.Printf("Wallet %s charged ₹%d\n", w.UserID, amount)
	return nil
}
