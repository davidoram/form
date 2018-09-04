package models

import (
	"time"

	"github.com/satori/go.uuid"
)

// Template represents the definition of a Form
type Template struct {
	// The internal ID used at the db level
	ID int

	// External facing identifier is combination of ExternalID & Version
	// So that a form can be versioned
	ExternalID uuid.UUID
	Version    int

	// The JSON form definition
	JSONSchema []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}
