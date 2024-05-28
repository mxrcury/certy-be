package config

func NewDatabaseConfig() (*DatabaseConfig, error) {
	dataSourceName, err := getEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	migrationsDir, err := getEnv("MIGRATIONS_DIR")
	if err != nil {
		return nil, err
	}

	databaseDriver, err := getEnv("DATABASE_DRIVER")
	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{DataSourceName: dataSourceName, MigrationsDir: migrationsDir, DatabaseDriver: databaseDriver}, nil
}
