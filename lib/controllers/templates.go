package controllers

import (
	"net/http"

	"github.com/davidoram/form/lib/context"
	"github.com/labstack/echo"
)

func GetNewTemplate(c echo.Context) error {
	fc := c.(*context.FormContext)
	return fc.Render(http.StatusOK, "views/templates/new.gohtml", map[string]interface{}{})
}
