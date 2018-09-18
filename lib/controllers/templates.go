package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/davidoram/form/lib/context"
	"github.com/davidoram/form/lib/formdb"
	"github.com/davidoram/form/lib/models"
	"github.com/labstack/echo"
)

// ListTemplates will list the existsing template
func ListTemplates(c echo.Context) error {
	fc := c.(*context.FormContext)
	tmpls, err := formdb.ListTemplates(fc.DB)
	if err != nil {
		return err
	}

	// An anonymous struct gives you more type safety than using a Map
	data := struct {
		Templates []*models.Template
	}{
		tmpls,
	}
	return fc.Render(http.StatusOK, "views/templates/list.gohtml", data)
}

// NewTemplate returns a new template
func NewTemplate(c echo.Context) error {
	fc := c.(*context.FormContext)
	return fc.Render(http.StatusOK, "views/templates/edit.gohtml", map[string]interface{}{})
}

// CreateTemplate inserts a new Template into the db
func CreateTemplate(c echo.Context) error {
	fc := c.(*context.FormContext)
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}
	if !isJSON(b) {
		return errors.New("Body not valid JSON")
	}

	tx := fc.DB.MustBegin()
	tmpl, err := formdb.InsertTemplate(tx, string(b))
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return fc.Render(http.StatusOK, "views/templates/edit.gohtml", map[string]interface{}{"template": tmpl})
}

// UpdateTemplate will update an existing template
func UpdateTemplate(c echo.Context) error {
	fc := c.(*context.FormContext)
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}
	if !isJSON(b) {
		return errors.New("Body not valid JSON")
	}

	tx := fc.DB.MustBegin()
	tmpl, err := formdb.InsertTemplate(tx, string(b))
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return fc.Render(http.StatusOK, "views/templates/edit.gohtml", map[string]interface{}{"template": tmpl})
}

// isJSON tests if the bytes represents valid JSON
func isJSON(b []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(b, &js) == nil
}
