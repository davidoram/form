package formdb

import (
	"github.com/davidoram/form/lib/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// GetTemplate reads a template by id
func GetTemplate(db *sqlx.Tx, id int64) (*models.Template, error) {
	sql := `SELECT id
	               ,json_schema
					FROM   templates
					WHERE  id = $1`
	var tpl models.Template
	err := db.QueryRowx(sql, tpl.ID).StructScan(&tpl)
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// InsertTemplate adds a new template into the database
func InsertTemplate(db *sqlx.Tx, js string) (*models.Template, error) {
	tpl := models.Template{
		ID:         0,
		JSONSchema: js,
	}
	sql := `INSERT INTO templates (
						created_at
						,updated_at
						,id
						,json_schema
					)
			    VALUES (
						DEFAULT
						,DEFAULT
						,DEFAULT
						,$1
					) RETURNING id`
	return &tpl, db.QueryRow(sql, js).Scan(&tpl.ID)
}

// func (c *context.FormContext) UpdateTemplate(id int) (*models.Template, error) {

// }

// func (c *context.FormContext) DeleteTemplate(id int) (*models.Template, error) {

// }
