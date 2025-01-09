package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func InitDB(addr string, maxOpenConn, maxIdleConn, maxIdleTime int) (*sqlx.DB, error) {
	// Open the connection.
	db, err := sqlx.Open("postgres", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxIdleTime) * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to make sure the connection is alive.
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
