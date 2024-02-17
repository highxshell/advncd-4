package http

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ContactHTTPDelivery struct {
	contactUseCase useCase.Contact
	groupUseCase   useCase.Group
}

func NewContactHTTP(contactUseCase useCase.Contact, groupUseCase useCase.Group) *ContactHTTPDelivery {
	return &ContactHTTPDelivery{contactUseCase: contactUseCase, groupUseCase: groupUseCase}
}

func (d *ContactHTTPDelivery) Start(port int) {
	const op = "delivery.http.Start"
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/contacts/create", d.HandleCreateContact)
	http.HandleFunc("/contacts/read", d.HandleReadContact)
	http.HandleFunc("/contacts/get", d.HandleGetAllContacts)
	http.HandleFunc("/contacts/update", d.HandleUpdateContact)
	http.HandleFunc("/contacts/delete", d.HandleDeleteContact)

	http.HandleFunc("/groups/create", d.HandleCreateGroup)
	http.HandleFunc("/groups/read", d.HandleReadGroup)
	http.HandleFunc("/groups/get", d.HandleGetAllGroups)
	http.HandleFunc("/groups/add", d.HandleAddContactToGroup)
	go func() {
		log.Println("Starting server on port", port)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("%s: %w", op, err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down gracefully...")
}

func (d *ContactHTTPDelivery) HandleCreateContact(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleCreateContact"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	var contactData contact.Contact
	if err := json.NewDecoder(r.Body).Decode(&contactData); err != nil {
		http.Error(w, op+" incorrect format of data: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := d.contactUseCase.CreateContact(ctx, contactData.FirstName, contactData.LastName, contactData.MiddleName, contactData.PhoneNumber); err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Contact created"))
}

func (d *ContactHTTPDelivery) HandleReadContact(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleReadContact"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	contactID := extractContactID(r)
	cont, err := d.contactUseCase.GetContactByID(ctx, contactID)
	if err != nil {
		if contactID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusNotFound)
		log.Printf("%s: %s\n", op, err.Error())
		return
	}
	fullName := fmt.Sprintf("%s %s %s", cont.LastName, cont.FirstName, cont.MiddleName)
	if cont.FullName() != fullName {
		http.Error(w, op+": "+err.Error(), http.StatusNotFound)
		log.Printf("%s: %s\n", op, err.Error())
		return
	}
	json.NewEncoder(w).Encode(cont)
}

func (d *ContactHTTPDelivery) HandleUpdateContact(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleUpdateContact"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	contactID := extractContactID(r)
	var updateData contact.Contact
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, op+" incorrect format of data: "+err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := d.contactUseCase.GetContactByID(ctx, contactID); err != nil {
		if contactID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	updateData.ID = contactID
	if err := updateData.SetPhoneNumber(updateData.PhoneNumber); err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := d.contactUseCase.UpdateContact(ctx, &updateData); err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact updated"))
}

func (d *ContactHTTPDelivery) HandleDeleteContact(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleDeleteContact"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	contactID := extractContactID(r)
	if _, err := d.contactUseCase.GetContactByID(ctx, contactID); err != nil {
		if contactID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	err := d.contactUseCase.DeleteContact(ctx, contactID)
	if err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact deleted"))
}

func (d *ContactHTTPDelivery) HandleGetAllContacts(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleGetAllContacts"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	contacts, err := d.contactUseCase.GetAllContacts(ctx)
	if err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contacts)
}

func (d *ContactHTTPDelivery) HandleCreateGroup(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleCreateGroup"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	var groupData group.Group
	if err := json.NewDecoder(r.Body).Decode(&groupData); err != nil {
		http.Error(w, op+" incorrect format of data: "+err.Error(), http.StatusBadRequest)
		return
	}
	err := d.groupUseCase.CreateGroup(ctx, groupData.Name)
	if err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Group created"))
}

func (d *ContactHTTPDelivery) HandleReadGroup(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleReadGroup"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	groupID := extractGroupID(r)
	gr, err := d.groupUseCase.GetGroupByID(ctx, groupID)
	if err != nil {
		if groupID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(gr)
}

func (d *ContactHTTPDelivery) HandleGetAllGroups(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleGetAllGroups"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	groups, err := d.groupUseCase.GetAllGroups(ctx)
	if err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(groups)
}

func (d *ContactHTTPDelivery) HandleAddContactToGroup(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.http.HandleAddContactToGroup"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", uuid.New())
	contactID := extractContactID(r)
	if _, err := d.contactUseCase.GetContactByID(ctx, contactID); err != nil {
		if contactID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	groupID := extractGroupID(r)
	if _, err := d.groupUseCase.GetGroupByID(ctx, groupID); err != nil {
		if groupID == uuid.Nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	err := d.groupUseCase.AddContactToGroup(ctx, contactID, groupID)
	if err != nil {
		http.Error(w, op+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact added to group"))
}

func extractContactID(r *http.Request) uuid.UUID {
	id := r.URL.Query().Get("id")
	contactID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return contactID
}

func extractGroupID(r *http.Request) uuid.UUID {
	groupID := r.URL.Query().Get("group_id")
	parsedGroupID, err := uuid.Parse(groupID)
	if err != nil {
		return uuid.Nil
	}
	return parsedGroupID
}
