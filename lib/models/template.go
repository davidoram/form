package models

// Template represents the definition of a Form
type Template struct {
	ID         int64  `db:"id"`
	JSONSchema string `db:"json_schema"`
}
