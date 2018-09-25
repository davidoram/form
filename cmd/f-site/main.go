package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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

	Migrate string `long:"db-migrate" description:"Database migration operation" default:"none" choice:"none" choice:"up" choice:"down" required:"false" env:"F_DB_MIGRATE"`

	HttpPort int `long:"http-port" description:"HTTP port to run on" required:"true" env:"F_HTTP_PORT"`

	LogLevel int `long:"log-level" descripton:"Logging level, 1=DEBUG, 2=INFO, 3=WARN, 4=ERROR" default:"2" choice:"1" choice:"2" choice:"3" choice:"4"  required:"false" env:"F_LOGLEVEL"`

	PrintRoutes bool `long:"print-routes" descripton:"Print the routes to stdout and exit."  required:"false"`
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

	if strings.EqualFold(opts.Migrate, "none") {
		e.Logger.Infof("Skip database migrations")
	} else {
		e.Logger.Infof("Migrating the database: %s", opts.Migrate)
		n, err := formdb.Migrate(db.DB, opts.Migrate)
		if err != nil {
			log.Fatal("db migrations failed: ", err)
		}
		e.Logger.Infof("%d migrations applied/removed", n)
		os.Exit(0)
	}

	// -------------------------------
	//            Middleware
	//
	e.Use(context.FormContextMiddleware(db)) // MUST be registered first
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.CORS())
	csrfConfig := middleware.DefaultCSRFConfig
	e.Use(middleware.CSRFWithConfig(csrfConfig))
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	// Access templates via embedded file system
	tbox := rice.MustFindBox("templates")
	assets := controllers.StaticAssets{Root: "public", Box: rice.MustFindBox("public")}

	// Setup template renderer
	e.Renderer = controllers.GetTemplateRenderer(tbox, &csrfConfig)

	// -------------------------------
	//            Routes
	//
	e.GET("/", controllers.Home).Name = "home"
	e.GET("/public/*", assets.GetStaticAssets).Name = "static"
	e.GET("/templates", controllers.ListTemplates).Name = "list_templates"
	e.GET("/templates/new", controllers.NewTemplate).Name = "get_new_template"
	e.POST("/templates/new", controllers.CreateTemplate).Name = "save_new_template"
	e.GET("/templates/:id", controllers.OpenTemplate).Name = "open_template"
	e.POST("/templates/:id", controllers.UpdateTemplate).Name = "update_template"

	// Print routes to stdout & finish
	if opts.PrintRoutes {
		routes, err := json.MarshalIndent(e.Routes(), "", "  ")
		if err != nil {
			fmt.Println("Error: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", routes)
		os.Exit(0)
	}

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", opts.HttpPort)))
}
