package user

import (
	"errors"
	"fmt"
	"strings"
)

/*
PATTERN: ENCAPSULATION + COMPOSITION

Encapsulation: Control access to internal data.
Composition: Build a larger model using smaller components.
Constructor: Create a valid object through one entry point.
*/

// COMPOSITION:
// Keep related contact details together in a separate component.
type contactInfo struct {
	email string
	phone string
}

// ENCAPSULATION:
// Unexported fields cannot be accessed directly from other packages.
type User struct {
	id      string
	name    string
	contact contactInfo
}

// CONSTRUCTOR:
// Validate required fields before creating a User.
func New(id, name, email, phone string) (*User, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	phone = strings.TrimSpace(phone)

	if id == "" || name == "" || email == "" {
		return nil, errors.New("id, name, and email are required")
	}

	// Return a pointer to the newly created User.
	return &User{
		id:   id,
		name: name,
		contact: contactInfo{
			email: email,
			phone: phone,
		},
	}, nil
}

// VALUE RECEIVER:
// Read the User's information without changing its state.
func (u User) Summary() string {
	return fmt.Sprintf(
		"ID: %s | Name: %s | Email: %s | Phone: %s",
		u.id,
		u.name,
		u.contact.email,
		u.contact.phone,
	)
}

// POINTER RECEIVER:
// Update the original User after validating the new email.
func (u *User) ChangeEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email cannot be empty")
	}

	u.contact.email = email
	return nil
}
