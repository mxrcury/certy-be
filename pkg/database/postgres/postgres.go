package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mxrcury/dragonsage/internal/config"
)

var listOfTables = []string{"users", "channels"}

func Init(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DataSourceName)

	if err != nil {
		return nil, err
	}

	areAllMigrationsApplied, err := verifyIfMigrationsApplied(db, listOfTables)
	if err != nil {
		return nil, err
	}

	if !areAllMigrationsApplied {
		if err := runMigrations(db.DB, cfg.MigrationsDir, cfg.DatabaseDriver); err != nil {
			return nil, err
		}
	}

	return db, nil
}
