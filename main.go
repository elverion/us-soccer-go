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
)

func main() {
	ctx, cancel := setup()

	if err := chix.RunContext(
		ctx, httpServer(),
	); err != nil {
		cancel()
		log.WithError(err).Fatal("shutting down")
	}
}
