package formdb

import (
	"database/sql"

	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
)

// Migrate runs all the migrations to get the form db up to the latest version
func Migrate(db *sql.DB) (int, error) {
	return migrate.Exec(db, "postgres", migrations(), migrate.Up)
}

func migrations() migrate.MigrationSource {
	return &migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations"),
	}
}
