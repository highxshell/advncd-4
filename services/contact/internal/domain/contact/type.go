package contact

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

type Contact struct {
	ID                                           uuid.UUID
	FirstName, LastName, MiddleName, PhoneNumber string
}

func New(id uuid.UUID, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {
	contact := &Contact{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
	}
	if err := contact.SetPhoneNumber(phoneNumber); err != nil {
		return nil, err
	}

	return contact, nil
}

func (c *Contact) FullName() string {
	return c.LastName + " " + c.FirstName + " " + c.MiddleName
}

func (c *Contact) SetPhoneNumber(phoneNumber string) error {
	const op = "internal.domain.SetPhoneNumber"
	matched, err := regexp.MatchString("^[0-9]+$", phoneNumber)
	if !matched {
		return fmt.Errorf("%s: phone numbers should contain only numbers", op)
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	c.PhoneNumber = phoneNumber
	return nil
}
