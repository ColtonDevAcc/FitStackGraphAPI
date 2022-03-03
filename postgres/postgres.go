package postgres

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/voodoostack/fitstackapi/config"
)

type DB struct {
	Pool *pgxpool.Pool
	conf *config.Config
}

func New(ctx context.Context, conf *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatalf("cant parse postgres config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("cant connect to postgres: %v", err)
	}

	db := &DB{Pool: pool, conf: conf}

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

func (db *DB) Migrate() error {
	_, b, _, _ := runtime.Caller(0)
	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(b))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrating up: %v", err)
	}

	log.Println("Migration done")

	return nil
}
