package main

import (
	"time"

	"github.com/Gononet-LLC/go-contact-service/schema"
	"github.com/uptrace/bun"
)

type GenericResponse struct {
	Success bool   `json:"success"`
	ID      int    `json:"id,omitempty"`
	UID     string `json:"uid,omitempty"`
}

type GenericErrorResponse struct {
	Success      bool   `json:"sucess"`
	ErrorMessage string `json:"error_message"` // <-- This will be removed in production. Currently in use in Dev phase
	Message      string `json:"message"`
}

type CreateContactRequest struct {
	Contact   schema.CreateContactParams          `json:"contact"`
	Addresses []schema.CreateContactAddressParams `json:"addresses"`
}

type ContactResponse struct {
	Pages    int64                       `json:"pages"`
	Contacts *[]schema.GetAllContactsRow `json:"contacts"`
}

/*type ContactFilter struct {
	TotalElement int       `json:"totalElement"`
	Record       int       `json:"record"`
	Contacts     []Contact `json:"content"`
	TotalPages   int       `json:"totalPages"`
	CurrentPage  int       `json:"currentPage"`
}
*/

/*
type Contact struct {
	ContactID       *string `json:"contact_id,omitempty"`
	FirstName       *string `json:"first_name,omitempty"`
	LastName        *string `json:"last_name,omitempty"`
	BusinessName    *string `json:"business_name,omitempty"`
	JobTitle        *string `json:"job_title,omitempty"`
	LocationGroupID *string `json:"location_group_id,omitempty"`
	image           *string `json:"image,omitempty`
	Description     *string `json:"description,omitempty"`
	ContactGroupID  int     `json:"contact_group_id,omitempty"`
	Address         *string `json:"address,omitempty"`
	Zip             *string `json:"zip,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	Mobile          *string `json:"mobile,omitempty"`
	Email           *string `json:"email,omitempty"`
}

type User1 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
*/

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Emails string `json:"emails"`
}

type Status struct {
	bun.BaseModel `bun:"public.status"`
	ID            string    `json:"id" bun:"type:uuid,pk"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	OrderNumber   int       `json:"order_number"`
	Active        bool      `json:"active"`
	UpdatedBy     string    `json:"updated_by"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     string    `json:"updated_at"`
	CreatedAt     string    `json:"created_at"`
	DeletedAt     time.Time `bun:",soft_delete"`
}
