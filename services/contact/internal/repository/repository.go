package repository

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain"
)

type ContactRepositoryImpl struct {
	db *postgres.Storage
}

func NewContactRepository(db *postgres.Storage) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{db: db}
}

func (c ContactRepositoryImpl) CreateContact(contact *domain.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) GetContactByID(id int) (*domain.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) UpdateContact(contact *domain.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) DeleteContact(id int) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) CreateGroup(group *domain.Group) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) GetGroupByID(id int) (*domain.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (c ContactRepositoryImpl) AddContactToGroup(contact *domain.Contact, group *domain.Group) error {
	//TODO implement me
	panic("implement me")
}
