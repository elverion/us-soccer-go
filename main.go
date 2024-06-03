package main

import (
	"us-soccer-go-test/internal/ent"
	"us-soccer-go-test/internal/models"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/clix"

	_ "us-soccer-go-test/internal/ent/runtime"
)

var (
	cli = &clix.CLI[models.Flags]{}

	logger log.Interface

	db *ent.Client

	cancelDaemon func()
)

func main() {
	ctx := setup()

	if err := chix.RunContext(
		ctx, httpServer(),
	); err != nil {

		if cli.Flags.RunDaemon {
			cancelDaemon()
		}

		logger.WithError(err).Fatal("shutting down")
	}
}
