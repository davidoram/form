package models

// Template represents the definition of a Form
type Template struct {
	// The internal ID used at the db level
	ID int64

	// The JSON form definition
	JSONSchema string
}
