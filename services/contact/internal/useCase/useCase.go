package useCase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"fmt"
	"math/rand"
)

type ContactUseCaseImpl struct {
	contactRepo repository.ContactRepository
}

func NewContactUseCase(contactRepo repository.ContactRepository) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		contactRepo: contactRepo,
	}
}

func (c ContactUseCaseImpl) CreateContact(firstName, lastName, middleName, phoneNumber string) error {
	const op = "internal.useCase.CreateContact"
	contact, err := domain.NewContact(rand.Intn(10000000), firstName, lastName, middleName, phoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.contactRepo.CreateContact(contact)
}

func (c ContactUseCaseImpl) GetContactByID(id int) (*domain.Contact, error) {
	return c.contactRepo.GetContactByID(id)
}

func (c ContactUseCaseImpl) UpdateContact(contact *domain.Contact) error {
	return c.contactRepo.UpdateContact(contact)
}

func (c ContactUseCaseImpl) DeleteContact(id int) error {
	return c.contactRepo.DeleteContact(id)
}

func (c ContactUseCaseImpl) CreateGroup(name string) error {
	const op = "internal.useCase.CreateGroup"
	group, err := domain.NewGroup(rand.Intn(10000000), name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.contactRepo.CreateGroup(group)
}

func (c ContactUseCaseImpl) GetGroupByID(id int) (*domain.Group, error) {
	return c.contactRepo.GetGroupByID(id)
}

func (c ContactUseCaseImpl) AddContactToGroup(contactID, groupID int) error {
	const op = "internal.useCase.AddContactToGroup"
	contact, err := c.contactRepo.GetContactByID(contactID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group, err := c.contactRepo.GetGroupByID(groupID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group.Contacts = append(group.Contacts, *contact)

	return c.contactRepo.AddContactToGroup(contact, group)
}
