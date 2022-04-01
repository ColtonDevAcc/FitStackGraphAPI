package postgres

import (
	"context"
	"fmt"
	"log"
	"net"
	"path"
	"runtime"

	"cloud.google.com/go/cloudsqlconn"
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
		log.Fatalf("can't parse postgres config: %v", err)
	}

	// Create a new dialer with any options
	d, err := cloudsqlconn.NewDialer(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dialer: %v", err)
	}
	defer d.Close()

	dbConf.ConnConfig.DialFunc = func(ctx context.Context, _ string, instance string) (net.Conn, error) {
		return d.Dial(ctx, "fitstack-343223:us-central1:fitstackapi")
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	fmt.Printf(dbConf.ConnString())

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

func (db *DB) Drop() error {
	_, b, _, _ := runtime.Caller(0)

	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(b))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error create the migrate instance: %v", err)
	}

	if err := m.Drop(); err != nil {
		return fmt.Errorf("error drop: %v", err)
	}

	log.Println("migration drop")

	return nil
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

func (db *DB) Truncate(ctx context.Context) error {
	if _, err := db.Pool.Exec(ctx, `
		DELETE FROM users;
	`); err != nil {
		return fmt.Errorf("error truncate: %v", err)
	}

	return nil
}
