package checkout

import "errors"

/*
PATTERN: DEPEND ON BEHAVIOR, NOT CONCRETE TYPES

Checkout should work with any payment implementation
that satisfies the PaymentMethod interface.
*/

// INTERFACE: Define only the behavior checkout needs.
type PaymentMethod interface {
	Pay(amount int) error
}

// SINGLE RESPONSIBILITY:
// Checkout manages order-payment coordination.
type Checkout struct {
	payment PaymentMethod
}

// DEPENDENCY INJECTION:
// Supply the payment implementation from outside.
func New(payment PaymentMethod) *Checkout {
	return &Checkout{
		payment: payment,
	}
}

// Validate the order before delegating payment processing.
func (c *Checkout) PayOrder(amount int) error {
	if amount <= 0 {
		return errors.New("order amount must be positive")
	}

	if c.payment == nil {
		return errors.New("payment method is required")
	}

	return c.payment.Pay(amount)
}
