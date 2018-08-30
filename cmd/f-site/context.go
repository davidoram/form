package main

import (
	"database/sql"

	"github.com/labstack/echo"
)

type FormContext struct {
	echo.Context
	*sql.DB
}
