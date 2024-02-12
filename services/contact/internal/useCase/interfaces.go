package useCase

import "architecture_go/services/contact/internal/domain"

type ContactUseCase interface {
	CreateContact(firstName, lastName, middleName, phoneNumber string) error
	GetContactByID(id int) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(id int) error
	CreateGroup(name string) error
	GetGroupByID(id int) (*domain.Group, error)
	AddContactToGroup(contactID, groupID int) error
}
