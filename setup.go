package main

import (
	"context"
	"us-soccer-go-test/internal/database"
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/weather"

	"github.com/apex/log"
)

func setup() context.Context {
	cli.Parse()
	logger = cli.Logger

	// Initalize the database
	db = database.Open(logger, cli.Flags.DB)

	ctx := ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(ctx, logger)

	if cli.Flags.RunDaemon {
		cancelDaemon = weather.StartWeatherDaemon(db, logger)
	}

	return ctx
}
