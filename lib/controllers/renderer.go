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

// RenderContext is passed to all templates
// URLFor is the method to call to generate a URL to a given handler eg: 'call .URLFor "home"'
// Data is the variable passed in by the controller
// TODO: Add other data like the current User
type RenderContext struct {
	LogInfo func(i ...interface{})
	URLFor  func(name string, params ...interface{}) string
	Data    interface{}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	ctx := RenderContext{LogInfo: c.Logger().Info, URLFor: c.Echo().Reverse, Data: data}
	// // Add global methods if data is a map
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	// 	c.Logger().Info("Added reverse")
	// 	viewContext["reverse"] =
	// }

	// get file contents as string
	templateString, err := t.box.String(name)
	if err != nil {
		c.Logger().Errorf("Can't get template: '%s' contents", name)
		return err
	}
	// parse and execute the template
	tmplMessage, err := template.New(name).Funcs(template.FuncMap{

		// The 'log' function can be used as follows within templates:
		// {{ log "This is a message" }}
		// It will print a message to the echo logs, and returns empty string
		"log": func(i interface{}) string {
			c.Logger().Infof("%+v", i)
			return ""
		},

		// The 'urlFor' function returns the URL for a given named route.
		// Usage:
		// 	{{ urlFor "home" }}
		"urlFor": func(name string, params ...interface{}) string {
			return c.Echo().Reverse(name, params)
		},
	}).Parse(templateString)

	if err != nil {
		c.Logger().Errorf("Can't parse template: '%s'", name)
		return err
	}
	err = tmplMessage.Execute(w, ctx)
	if err != nil {
		c.Logger().Errorf("Error executing template: '%s'. '%+v'", name, err)
	}
	return err

}
