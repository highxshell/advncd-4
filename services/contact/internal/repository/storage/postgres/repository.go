package postgres

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type ContactRepositoryImpl struct {
	db *postgres.Storage
}

func NewContactRepository(db *postgres.Storage) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{db: db}
}

func (c *ContactRepositoryImpl) CreateContact(ctx context.Context, contact *contact.Contact) error {
	const op = "repository.storage.postgres.CreateContact"
	query := `
        INSERT INTO contacts (id, first_name, last_name, middle_name, phone_number)
        VALUES ($1, $2, $3, $4, $5)
    `
	if err := c.db.DB.QueryRow(query, contact.ID, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return nil
}

func (c *ContactRepositoryImpl) GetContactByID(ctx context.Context, id uuid.UUID) (*contact.Contact, error) {
	const op = "repository.storage.postgres.GetContactByID"
	var (
		contactID                                    uuid.UUID
		firstName, lastName, middleName, phoneNumber string
	)
	query := `
        SELECT id, first_name, last_name, middle_name, phone_number
        FROM contacts
        WHERE id = $1
    `
	row := c.db.DB.QueryRow(query, id)
	if err := row.Scan(&contactID, &firstName, &lastName, &middleName, &phoneNumber); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("%s: contact with id (%s) not found\n", op, id.String())
			return nil, fmt.Errorf("contact with id (%s) not found", id.String())
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	con, _ := contact.New(id, firstName, lastName, middleName, phoneNumber)
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return con, nil
}

func (c *ContactRepositoryImpl) UpdateContact(ctx context.Context, contact *contact.Contact) error {
	const op = "repository.storage.postgres.UpdateContact"
	query := `
        UPDATE contacts
        SET first_name = $2, last_name = $3, middle_name = $4, phone_number = $5
        WHERE id = $1
    `
	_, err := c.db.DB.Exec(query, contact.ID, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return nil
}

func (c *ContactRepositoryImpl) DeleteContact(ctx context.Context, id uuid.UUID) error {
	const op = "repository.storage.postgres.DeleteContact"
	query := `
        DELETE FROM contacts 
        WHERE id = $1
    `
	_, err := c.db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s:request with context id = %s\n", op, ctx.Value("ID"))

	return nil
}

func (c *ContactRepositoryImpl) CreateGroup(ctx context.Context, group *group.Group) error {
	const op = "repository.storage.postgres.CreateGroup"
	query := `
        INSERT INTO groups (id, name)
        VALUES ($1, $2)
    `
	if err := c.db.DB.QueryRow(query, group.ID, group.Name).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return nil
}

func (c *ContactRepositoryImpl) GetGroupByID(ctx context.Context, id uuid.UUID) (*group.Group, error) {
	const op = "repository.storage.postgres.GetGroupByID"
	query := `
        SELECT name
        FROM groups
        WHERE id = $1
    `
	row := c.db.DB.QueryRow(query, id)
	var groupName string
	if err := row.Scan(&groupName); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("group with id (%s) not found", id.String())
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	gr, _ := group.New(id, groupName)
	contacts, err := c.getGroupContacts(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for _, uid := range contacts {
		con, _ := c.GetContactByID(ctx, uid)
		gr.Contacts = append(gr.Contacts, *con)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))
	return gr, nil
}

func (c *ContactRepositoryImpl) AddContactToGroup(ctx context.Context, contact *contact.Contact, group *group.Group) error {
	const op = "repository.storage.postgres.AddContactToGroup"
	if _, err := c.GetContactByID(ctx, contact.ID); err != nil {
		return fmt.Errorf("contact with id (%s) not found", contact.ID.String())
	}
	if _, err := c.GetGroupByID(ctx, group.ID); err != nil {
		return fmt.Errorf("group with id (%s) not found", group.ID.String())
	}
	query := `
        INSERT INTO group_contacts (group_id, contact_id)
        VALUES ($1, $2)
    `
	_, err := c.db.DB.Exec(query, group.ID, contact.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return nil
}

func (c *ContactRepositoryImpl) GetAllContacts(ctx context.Context) ([]*contact.Contact, error) {
	const op = "repository.storage.postgres.GetAllContacts"
	query := `
        SELECT id, first_name, last_name, middle_name, phone_number
        FROM contacts
    `
	rows, err := c.db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var contacts []*contact.Contact
	for rows.Next() {
		var contact contact.Contact
		err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contacts = append(contacts, &contact)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s: request with context id = %s\n", op, ctx.Value("ID"))

	return contacts, nil
}

func (c *ContactRepositoryImpl) GetAllGroups(ctx context.Context) ([]*group.Group, error) {
	const op = "repository.storage.postgres.GetAllGroups"
	query := `
        SELECT id, name
        FROM groups
    `
	rows, err := c.db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var groups []*group.Group
	for rows.Next() {
		var gr group.Group
		err := rows.Scan(&gr.ID, &gr.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contacts, err := c.getGroupContacts(gr.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		for _, uid := range contacts {
			con, _ := c.GetContactByID(ctx, uid)
			gr.Contacts = append(gr.Contacts, *con)
		}
		groups = append(groups, &gr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	fmt.Printf("%s:request with context id = %s\n", op, ctx.Value("ID"))

	return groups, nil
}

func (c *ContactRepositoryImpl) getGroupContacts(groupID uuid.UUID) ([]uuid.UUID, error) {
	const op = "repository.storage.postgres.getGroupContacts"
	query := `
        SELECT contact_id
        FROM group_contacts
        WHERE group_id = $1
    `
	rows, err := c.db.DB.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var contacts []uuid.UUID
	for rows.Next() {
		var contactID uuid.UUID
		err := rows.Scan(&contactID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contacts = append(contacts, contactID)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contacts, nil
}
