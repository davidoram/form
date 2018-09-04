package controllers

import (
	"html/template"
	"io"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	box *rice.Box
}

func GetTemplateRenderer(box *rice.Box) *TemplateRenderer {
	return &TemplateRenderer{
		box: box,
	}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	// get file contents as string
	templateString, err := t.box.String(name)
	if err != nil {
		c.Logger().Errorf("Can't get template: '%s' contents", name)
		return err
	}
	// parse and execute the template
	tmplMessage, err := template.New("message").Parse(templateString)
	if err != nil {
		c.Logger().Errorf("Can't parse template: '%s'", name)
		return err
	}
	err = tmplMessage.Execute(w, data)
	if err != nil {
		c.Logger().Errorf("Error executing template: '%s'. '%+v'", name, err)
	}
	return err

}
