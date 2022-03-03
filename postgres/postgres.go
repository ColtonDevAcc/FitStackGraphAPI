package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context) *DB {
	dbConf, err := pgxpool.ParseConfig()
	if err != nil {
		log.Fatalf("cant parse postgres config: %v", err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("cant connect to postgres: %v", err)
	}

	db := &DB{Pool: pool}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("cant ping postgres: %v", err)
	}

	log.Println("postgres connected")
}

func (db *DB) Close() {
	db.Pool.Close()
}
