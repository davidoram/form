package formdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Force loading of the postgres drivers
)

func Connect(host, user, dbname, password, sslmode string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	return sql.Open("postgres", psqlInfo)
}
