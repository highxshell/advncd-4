package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"github.com/google/uuid"
)

type Contact interface {
	CreateContact(ctx context.Context, firstName, lastName, middleName, phoneNumber string) error
	GetContactByID(ctx context.Context, id uuid.UUID) (*contact.Contact, error)
	UpdateContact(ctx context.Context, contact *contact.Contact) error
	DeleteContact(ctx context.Context, id uuid.UUID) error
	GetAllContacts(ctx context.Context) ([]*contact.Contact, error)
}

type Group interface {
	CreateGroup(ctx context.Context, name string) error
	GetGroupByID(ctx context.Context, id uuid.UUID) (*group.Group, error)
	AddContactToGroup(ctx context.Context, contactID, groupID uuid.UUID) error
	GetAllGroups(ctx context.Context) ([]*group.Group, error)
}
