package formdb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
)

// Migrate runs migrations "Up" or "Down" as required
func Migrate(db *sql.DB, direction string) (int, error) {
	var dir migrate.MigrationDirection
	if strings.EqualFold(direction, "up") {
		dir = migrate.Up
	} else if strings.EqualFold(direction, "down") {
		dir = migrate.Down
	} else {
		return 0, fmt.Errorf("Invalid migration direction '%s', expect 'up' or 'down'", direction)
	}
	return migrate.Exec(db, "postgres", migrations(), dir)
}

func migrations() migrate.MigrationSource {
	return &migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations"),
	}
}
