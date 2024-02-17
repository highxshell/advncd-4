package contact

import (
	contactType "architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ContactUseCaseImpl struct {
	contactRepo storage.Contact
}

func NewContactUseCase(contactRepo storage.Contact) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		contactRepo: contactRepo,
	}
}

func (c ContactUseCaseImpl) CreateContact(ctx context.Context, firstName, lastName, middleName, phoneNumber string) error {
	const op = "internal.useCase.CreateContact"
	contact, err := contactType.New(uuid.New(), firstName, lastName, middleName, phoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.contactRepo.CreateContact(ctx, contact)
}

func (c ContactUseCaseImpl) GetContactByID(ctx context.Context, id uuid.UUID) (*contactType.Contact, error) {
	return c.contactRepo.GetContactByID(ctx, id)
}

func (c ContactUseCaseImpl) UpdateContact(ctx context.Context, contact *contactType.Contact) error {
	return c.contactRepo.UpdateContact(ctx, contact)
}

func (c ContactUseCaseImpl) DeleteContact(ctx context.Context, id uuid.UUID) error {
	return c.contactRepo.DeleteContact(ctx, id)
}

func (c ContactUseCaseImpl) GetAllContacts(ctx context.Context) ([]*contactType.Contact, error) {
	return c.contactRepo.GetAllContacts(ctx)
}
