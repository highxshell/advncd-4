package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Contact interface {
	CreateContact(firstName, lastName, middleName, phoneNumber string) error
	GetContactByID(id uuid.UUID) (*contact.Contact, error)
	UpdateContact(contact *contact.Contact) error
	DeleteContact(id uuid.UUID) error
}

type Group interface {
	CreateGroup(name string) error
	GetGroupByID(id uuid.UUID) (*group.Group, error)
	AddContactToGroup(contactID, groupID uuid.UUID) error
}
