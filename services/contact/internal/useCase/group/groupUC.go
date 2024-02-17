package group

import (
	groupType "architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"context"
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

func (c GroupUseCaseImpl) CreateGroup(ctx context.Context, name string) error {
	const op = "internal.useCase.CreateGroup"
	group, err := groupType.New(uuid.New(), name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.groupRepo.CreateGroup(ctx, group)
}

func (c GroupUseCaseImpl) GetGroupByID(ctx context.Context, id uuid.UUID) (*groupType.Group, error) {
	return c.groupRepo.GetGroupByID(ctx, id)
}

func (c GroupUseCaseImpl) AddContactToGroup(ctx context.Context, contactID, groupID uuid.UUID) error {
	const op = "internal.useCase.AddContactToGroup"
	contact, err := c.contactRepo.GetContactByID(ctx, contactID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group, err := c.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	group.Contacts = append(group.Contacts, *contact)

	return c.groupRepo.AddContactToGroup(ctx, contact, group)
}

func (c GroupUseCaseImpl) GetAllGroups(ctx context.Context) ([]*groupType.Group, error) {
	return c.groupRepo.GetAllGroups(ctx)
}
