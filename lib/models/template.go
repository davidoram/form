package models

// Template represents the definition of a Form
type Template struct {
	ID         int64  `db:"id"`
	JSONSchema string `db:"json_schema"`
}

// IsUnsaved returns true if the Template has not yet been persisted to the db &  assigned an ID
func (t *Template) IsUnsaved() bool {
	return t.ID == 0
}
