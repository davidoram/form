package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidoram/form/lib/formdb"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	flags "github.com/jessevdk/go-flags"
)

// CmdOptions defines the command line options
type CmdOptions struct {
	DbHost    string `long:"db-host" description:"Database host" required:"true" env:"F_DB_HOST"`
	DbPort    int    `long:"db-port" description:"Database port" required:"true" env:"F_DB_PORT"`
	DbName    string `long:"db-name" description:"Database name" required:"true" env:"F_DB_NAME"`
	DbUser    string `long:"db-user" description:"Database username" required:"true" env:"F_DB_USER"`
	DbPass    string `long:"db-password" description:"Database password" required:"true" env:"F_DB_PASSWORD"`
	DBSSLMode string `long:"db-ssl-mode" description:"Database sslmode" required:"false" env:"F_DB_SSLMODE"`

	HttpPort int `long:"http-port" description:"HTTP port to run on" required:"true" env:"F_HTTP_PORT"`
}

func main() {
	// Echo instance
	e := echo.New()

	var opts CmdOptions
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	e.Logger.Info("Opening db connection...")
	db, err := formdb.Connect(opts.DbHost, opts.DbUser, opts.DbName, opts.DbPass, opts.DBSSLMode, opts.DbPort)
	if err != nil {
		panic(err)
	}

	e.Logger.Info("Migrating the database...")
	n, err := formdb.Migrate(db)
	if err != nil {
		log.Fatal("db migrations failed: ", err)
	}
	e.Logger.Infof("%d migrations applied", n)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", opts.HttpPort)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
