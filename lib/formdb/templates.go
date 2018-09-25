package formdb

import (
	"github.com/davidoram/form/lib/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ListTemplates reads the list of all templates
func ListTemplates(db sqlx.Queryer) ([]*models.Template, error) {
	sql := `SELECT id
	               ,json_schema
					FROM   templates`
	var tpls []*models.Template
	rows, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tpl models.Template
		err = rows.StructScan(&tpl)
		if err != nil {
			return nil, err
		}
		tpls = append(tpls, &tpl)
	}
	return tpls, nil
}

// GetNewTemplate constructs a new models.Template
func GetNewTemplate(db sqlx.Queryer) (*models.Template, error) {
	tpl := models.Template{
		ID:         0,
		JSONSchema: "{}",
	}
	return &tpl, nil
}

// GetTemplate reads a template by id
func GetTemplate(db sqlx.Queryer, id int64) (*models.Template, error) {
	sql := `SELECT id
	               ,json_schema
					FROM   templates
					WHERE  id = $1`
	tpl := models.Template{ID: id}
	err := db.QueryRowx(sql, tpl.ID).StructScan(&tpl)
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// InsertTemplate adds a new template into the database
func InsertTemplate(db sqlx.Queryer, js string) (*models.Template, error) {
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
	return &tpl, db.QueryRowx(sql, js).Scan(&tpl.ID)
}

// func (c *context.FormContext) UpdateTemplate(id int) (*models.Template, error) {

// }

// func (c *context.FormContext) DeleteTemplate(id int) (*models.Template, error) {

// }
