package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	flags "github.com/jessevdk/go-flags"
)

// CmdOptions defines the command line options
type CmdOptions struct {
	DbHost string `long:"db-host" description:"Database host" required:"true" env:"F_DB_HOST"`
	DbPort int    `long:"db-port" description:"Database port" required:"true" env:"F_DB_PORT"`
	DbName string `long:"db-name" description:"Database name" required:"true" env:"F_DB_NAME"`
	DbUser string `long:"db-user" description:"Database username" required:"true" env:"F_DB_USER"`
	DbPass string `long:"db-password" description:"Database password" required:"true" env:"F_DB_PASSWORD"`

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
