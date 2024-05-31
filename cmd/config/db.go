package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func psql(config *Config) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(buildConnectionString(config))
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func buildConnectionString(config *Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName, config.DbSslMode)
}
