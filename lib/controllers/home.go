package controllers

import (
	"net/http"

	"github.com/davidoram/form/lib/context"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	fc := c.(*context.FormContext)
	return fc.Render(http.StatusOK, "views/home.gohtml", map[string]interface{}{})
}
