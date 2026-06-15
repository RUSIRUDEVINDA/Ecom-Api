package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/RUSIRUDEVINDA/Ecom-Api/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx) // meaning of this line is to close the connection to the database when the main function exits.
	// This is important to free up resources and prevent potential memory leaks. By using defer, we ensure that the connection will be closed regardless of how the function exits, whether it completes successfully or encounters an error.

	logger.Info("connected to the database", "dsn", cfg.db.dsn)
	api := application{
		config: cfg,
		db:conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}

}
