package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type Contact interface {
	CreateContact(contact *contact.Contact) error
	GetContactByID(id uuid.UUID) (*contact.Contact, error)
	UpdateContact(contact *contact.Contact) error
	DeleteContact(id uuid.UUID) error
}

type Group interface {
	CreateGroup(group *group.Group) error
	GetGroupByID(id uuid.UUID) (*group.Group, error)
	AddContactToGroup(contact *contact.Contact, group *group.Group) error
}
