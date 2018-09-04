package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/davidoram/form/lib/context"
	"github.com/davidoram/form/lib/controllers"
	"github.com/davidoram/form/lib/formdb"

	"github.com/gorilla/sessions"
	flags "github.com/jessevdk/go-flags"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	lslog "github.com/labstack/gommon/log"
)

// CmdOptions defines the command line options
type CmdOptions struct {
	DbHost    string `long:"db-host" description:"Database host" required:"true" env:"F_DB_HOST"`
	DbPort    int    `long:"db-port" description:"Database port" required:"true" env:"F_DB_PORT"`
	DbName    string `long:"db-name" description:"Database name" required:"true" env:"F_DB_NAME"`
	DbUser    string `long:"db-user" description:"Database username" required:"true" env:"F_DB_USER"`
	DbPass    string `long:"db-password" description:"Database password" required:"true" env:"F_DB_PASSWORD"`
	DBSSLMode string `long:"db-ssl-mode" description:"Database sslmode" default:"prefer" choice:"disable" choice:"allow" choice:"prefer" choice:"require" choice:"verify-ca" choice:"verify-full" required:"false" env:"F_DB_SSLMODE"`

	HttpPort int `long:"http-port" description:"HTTP port to run on" required:"true" env:"F_HTTP_PORT"`

	LogLevel int `long:"log-level" descripton:"Logging level, 1=DEBUG, 2=INFO, 3=WARN, 4=ERROR" default:"2" choice:"1" choice:"2" choice:"3" choice:"4"  required:"false" env:"F_LOGLEVEL"`
}

func main() {
	// Echo instance
	e := echo.New()

	var opts CmdOptions
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}
	e.Logger.SetLevel(lslog.Lvl(opts.LogLevel))
	e.Logger.Info("f-server booting")

	e.Logger.Info("Opening db connection")
	db, err := formdb.Connect(opts.DbHost, opts.DbUser, opts.DbName, opts.DbPass, opts.DBSSLMode, opts.DbPort)
	if err != nil {
		panic(err)
	}

	e.Logger.Info("Migrating the database")
	n, err := formdb.Migrate(db)
	if err != nil {
		log.Fatal("db migrations failed: ", err)
	}
	e.Logger.Infof("%d migrations applied", n)

	// Middleware
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.FormContext{c, db}
			return h(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	// Access templates via embedded file system
	box := rice.MustFindBox("templates")

	// Setup templates
	e.Renderer = controllers.GetTemplateRenderer(box)

	// Routes
	e.GET("/", hello)
	e.GET("/builder/new", controllers.NewFormBuilder).Name = "new_form_builder"

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", opts.HttpPort)))
}

// Handler
func hello(c echo.Context) error {
	fc := c.(*context.FormContext)
	_, err := fc.DB.Exec("select 1")
	return c.String(http.StatusOK, fmt.Sprintf("Hello, World! %v", err))
}
