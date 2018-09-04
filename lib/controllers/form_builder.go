package controllers

import (
	"net/http"

	"github.com/davidoram/form/lib/context"
	"github.com/labstack/echo"
)

// NewFormBuilder displays a new form builder
func NewFormBuilder(c echo.Context) error {
	fc := c.(*context.FormContext)
	//_, err := fc.DB.Exec("select 1")
	return fc.Render(http.StatusOK, "views/new_form.html", map[string]interface{}{
		"name": "Dolly!",
	})
}
