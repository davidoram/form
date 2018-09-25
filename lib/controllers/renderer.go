package controllers

import (
	"fmt"
	"html/template"
	"io"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	box        *rice.Box
	csrfConfig *middleware.CSRFConfig
}

func GetTemplateRenderer(box *rice.Box, csrfConfig *middleware.CSRFConfig) *TemplateRenderer {
	return &TemplateRenderer{
		box:        box,
		csrfConfig: csrfConfig,
	}
}

// Render renders a template document
// Inside the template refer to any data passed in as '{{ name }}' etc
// Also provides functions 'log', 'urlFor', 'csrfHeader', 'csrfToken' and 'assetUrl'
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

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
			c.Logger().Infof("urlFor('%s',%+v)", name, params)
			return c.Echo().Reverse(name, params...)
		},

		// The 'assetUrl' function returns the URL for a given file under
		// he /public directory
		// Usage:
		// 	{{ assetUrl "css/styles.css" }}
		"assetUrl": func(file string) string {
			return fmt.Sprintf("/public/%s", file)
		},

		// The 'csrfHeader' function returns the name of the Header that will contain
		// the CSRF token value. See also 'csrfToken'
		// Usage:
		// 	{{ csrfHeader }}
		"csrfHeader": func() string {
			return echo.HeaderXCSRFToken
		},

		// The 'csrfToken' function returns the CSRF token value. See also 'csrfHeader'
		// Usage:
		// 	{{ csrfToken }}
		"csrfToken": func() string {
			return c.Get(t.csrfConfig.ContextKey).(string)
		},
	}).Parse(templateString)

	if err != nil {
		c.Logger().Errorf("Can't parse template: '%s'", name)
		return err
	}
	c.Logger().Debugf("Template variables: %+v", data)
	err = tmplMessage.Execute(w, data)
	if err != nil {
		c.Logger().Errorf("Error executing template: '%s'. '%+v'", name, err)
	}
	return err

}
