package formdb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Force loading of the postgres drivers
)

func Connect(host, user, dbname, password, sslmode string, port int) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	return sqlx.Open("postgres", psqlInfo)
}
