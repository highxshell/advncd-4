package repository

import "architecture_go/services/contact/internal/domain"

type ContactRepository interface {
	CreateContact(contact *domain.Contact) error
	GetContactByID(id int) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) error
	GetGroupByID(id int) (*domain.Group, error)
	AddContactToGroup(contact *domain.Contact, group *domain.Group) error
}
