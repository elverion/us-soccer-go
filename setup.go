package main

import (
	"context"
	"us-soccer-go-test/internal/database"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/weather"

	"github.com/apex/log"
)

func setup() (context.Context, func()) {
	cli.Parse()
	logger = cli.Logger

	if !cli.Flags.Configured {
		logger.Fatal("Not configured yet, please configure")
	}

	// Initalize the database
	db = database.Open(logger, cli.Flags.DB)

	ctx := ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(ctx, logger)

	cancel := weather.StartWeatherDaemon(db)

	return ctx, cancel
}
