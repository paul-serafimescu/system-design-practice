package database

import (
	"context"
	"fmt"
	"http-server/config"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func ConnectToDB(ctx context.Context, cfg *config.Config) (*postgres, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, "5432", cfg.DbName)

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Println("cannot connect")
		}

		pgInstance = &postgres{db}
	})

	return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
