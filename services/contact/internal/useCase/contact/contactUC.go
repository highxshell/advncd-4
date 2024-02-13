package contact

import (
	contactType "architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"fmt"
	"github.com/google/uuid"
)

type ContactUseCaseImpl struct {
	contactRepo storage.ContactRepository
}

func NewContactUseCase(contactRepo storage.ContactRepository) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		contactRepo: contactRepo,
	}
}

func (c ContactUseCaseImpl) CreateContact(firstName, lastName, middleName, phoneNumber string) error {
	const op = "internal.useCase.CreateContact"
	contact, err := contactType.NewContact(uuid.New(), firstName, lastName, middleName, phoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.contactRepo.CreateContact(contact)
}

func (c ContactUseCaseImpl) GetContactByID(id uuid.UUID) (*contactType.Contact, error) {
	return c.contactRepo.GetContactByID(id)
}

func (c ContactUseCaseImpl) UpdateContact(contact *contactType.Contact) error {
	return c.contactRepo.UpdateContact(contact)
}

func (c ContactUseCaseImpl) DeleteContact(id uuid.UUID) error {
	return c.contactRepo.DeleteContact(id)
}
