package postgres

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"github.com/google/uuid"
)

type ContactRepositoryImpl struct {
	db *postgres.Storage
}

func NewContactRepository(db *postgres.Storage) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{db: db}
}

func (c ContactRepositoryImpl) CreateContact(contact *contact.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) GetContactByID(id uuid.UUID) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) UpdateContact(contact *contact.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) DeleteContact(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) CreateGroup(group *group.Group) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) GetGroupByID(id uuid.UUID) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) AddContactToGroup(contact *contact.Contact, group *group.Group) error {
	//TODO implement me
	panic("implement me")
}
