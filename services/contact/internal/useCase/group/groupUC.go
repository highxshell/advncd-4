package group

import (
	groupType "architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"fmt"
	"github.com/google/uuid"
)

type GroupUseCaseImpl struct {
	groupRepo   storage.Group
	contactRepo storage.Contact
}

func NewGroupUseCase(groupRepo storage.Group, contactRepo storage.Contact) *GroupUseCaseImpl {
	return &GroupUseCaseImpl{
		groupRepo:   groupRepo,
		contactRepo: contactRepo,
	}
}

func (c GroupUseCaseImpl) CreateGroup(name string) error {
	const op = "internal.useCase.CreateGroup"
	group, err := groupType.New(uuid.New(), name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.groupRepo.CreateGroup(group)
}

func (c GroupUseCaseImpl) GetGroupByID(id uuid.UUID) (*groupType.Group, error) {
	return c.groupRepo.GetGroupByID(id)
}

func (c GroupUseCaseImpl) AddContactToGroup(contactID, groupID uuid.UUID) error {
	const op = "internal.useCase.AddContactToGroup"
	contact, err := c.contactRepo.GetContactByID(contactID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group, err := c.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group.Contacts = append(group.Contacts, *contact)

	return c.groupRepo.AddContactToGroup(contact, group)
}
