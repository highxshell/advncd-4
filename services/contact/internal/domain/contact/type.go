package contact

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

type Contact struct {
	ID                                           uuid.UUID
	firstName, lastName, middleName, phoneNumber string
}

func NewContact(id uuid.UUID, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {
	contact := &Contact{
		ID:         id,
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
	}
	if err := contact.SetPhoneNumber(phoneNumber); err != nil {
		return nil, err
	}

	return contact, nil
}

func (c *Contact) FullName() string {
	return c.lastName + " " + c.firstName + " " + c.middleName
}

func (c *Contact) SetPhoneNumber(phoneNumber string) error {
	const op = "internal.domain.SetPhoneNumber"
	_, err := regexp.MatchString("^[0-9]+$", phoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	c.phoneNumber = phoneNumber
	return nil
}
