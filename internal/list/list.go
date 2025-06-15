package list

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/faizanabbas/godo/internal/db"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

type List struct {
	queries *db.Queries
	db      *sql.DB
}

func New(dbPath string) (*List, error) {
	sqlite, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	migrationsDir := "db/migrations"
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		migrationsDir = "../db/migrations"
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			migrationsDir = "../../db/migrations"
			if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
				sqlite.Close()
				return nil, fmt.Errorf("could not find migrations directory")
			}
		}
	}
	migrationsDir, err = filepath.Abs(migrationsDir)
	if err != nil {
		sqlite.Close()
		return nil, fmt.Errorf("error getting absolute path: %w", err)
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}
	_, err = migrate.Exec(sqlite, "sqlite3", migrations, migrate.Up)
	if err != nil {
		sqlite.Close()
		return nil, fmt.Errorf("error running migrations: %w", err)
	}
	queries := db.New(sqlite)
	return &List{
		queries: queries,
		db:      sqlite,
	}, nil
}

func (l *List) Close() error {
	return l.db.Close()
}

func (l *List) Add(text string) (db.Godo, error) {
	if strings.TrimSpace(text) == "" {
		return db.Godo{}, errors.New("text is empty")
	}
	ctx := context.Background()
	godo, err := l.queries.CreateGodo(ctx, db.CreateGodoParams{
		Text: text,
		Done: false,
	})
	if err != nil {
		err = fmt.Errorf("error creating godo: %w", err)
	}
	return godo, err
}

func (l *List) Complete(id int64) error {
	ctx := context.Background()
	_, err := l.queries.GetGodo(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("godo with id %d not found", id)
		}
		return fmt.Errorf("error checking godo existence: %w", err)
	}
	err = l.queries.UpdateGodoDone(ctx, db.UpdateGodoDoneParams{
		ID:   id,
		Done: true,
	})
	if err != nil {
		return fmt.Errorf("error completing godo: %w", err)
	}
	return nil
}

func (l *List) String() string {
	ctx := context.Background()
	godos, err := l.queries.ListGodos(ctx)
	if err != nil {
		return fmt.Sprintf("Error listing godos: %v", err)
	}
	if len(godos) == 0 {
		return "No godos in list"
	}
	result := "Godo list:\n"
	for _, godo := range godos {
		status := " "
		if godo.Done {
			status = "✓"
		}
		result += fmt.Sprintf("%d. [%s] %s\n", godo.ID, status, godo.Text)
	}
	return result
}
