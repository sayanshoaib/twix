package app

import (
	"context"
	"fmt"
	"twix/config"
	"twix/ent"
)

func ConnectToPostgresql(cfg *config.DatabaseCfg) (*ent.Client, error) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password)

	if cfg.SSLEnabled {
		dbUrl += " sslmode=require"
	} else {
		dbUrl += " sslmode=disable"
	}

	client, err := ent.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, err
	}

	tx.Rollback()

	return client, nil
}
