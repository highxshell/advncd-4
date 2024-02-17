package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	DB *sql.DB
}

func New(host, user, password, dbname string, port int) (*Storage, error) {
	const op = "pkg.storage.postgres.New"
	storagePath := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := sql.Open("pgx", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{DB: db}, nil
}
