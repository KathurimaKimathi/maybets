package postgres

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"

	// responsible for methods to connect to db for migrations to run
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	assets "github.com/KathurimaKimathi/maybets/db"
)

// RunMigrations applies all pending migrations
func RunMigrations() error {
	driver, err := iofs.New(assets.DBMigrations, "migrations")
	if err != nil {
		return err
	}

	path := os.Getenv("SQLITE_URL")
	if path == "" {
		path = "../../../../../.."
	}

	dbPath := fmt.Sprintf("%s/bets.db", path)

	m, err := migrate.NewWithSourceInstance("iofs", driver, "sqlite3://"+dbPath)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
