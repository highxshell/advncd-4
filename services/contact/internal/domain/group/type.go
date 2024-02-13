package group

import (
	"architecture_go/services/contact/internal/domain/contact"
	"fmt"
	"github.com/google/uuid"
)

type Group struct {
	ID       uuid.UUID
	Name     string
	Contacts []contact.Contact
}

func NewGroup(id uuid.UUID, name string) (*Group, error) {
	group := &Group{
		ID:       id,
		Contacts: []contact.Contact{},
	}
	if err := group.SetName(name); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *Group) SetName(name string) error {
	const op = "internal.domain.SetName"
	if len(name) > 250 {
		return fmt.Errorf("%s: %s", op, "name should be less than 250 symbols")
	}
	g.Name = name
	return nil
}
