package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func runMigrations(db *sql.DB, migrationsDir string, databaseDriver string) error {
	n, err := migrate.Exec(db, databaseDriver, migrate.FileMigrationSource{Dir: migrationsDir}, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations\n", n)

	return nil
}

func verifyIfMigrationsApplied(db *sqlx.DB, listOfTables []string) (areAllMigrationsRan bool, err error) {
	for _, tableName := range listOfTables {
		err := db.Get(&areAllMigrationsRan, `SELECT EXISTS (
		SELECT FROM 
			pg_tables
		WHERE 
			schemaname = 'public' AND 
			tablename  = $1
		);`, tableName)

		if err != nil {
			return false, err
		}
	}

	return areAllMigrationsRan, nil
}
