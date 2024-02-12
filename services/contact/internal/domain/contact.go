package domain

import (
	"fmt"
	"regexp"
)

type Contact struct {
	ID                                           int
	firstName, lastName, middleName, phoneNumber string
}

func NewContact(id int, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {
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

type Group struct {
	ID       int
	Name     string
	Contacts []Contact
}

func NewGroup(id int, name string) (*Group, error) {
	group := &Group{
		ID:       id,
		Contacts: []Contact{},
	}
	if err := group.SetName(name); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *Group) SetName(name string) error {
	const op = "internal.domain.SetName"
	if len(name) > 250 {
		return fmt.Errorf("%s: %s", op, "name should be less than 250 symbols")
	}
	g.Name = name
	return nil
}
