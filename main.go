package main

import (
	"context"
	"us-soccer-go-test/internal/models"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/clix"
)

var (
	cli = &clix.CLI[models.Flags]{}

	logger log.Interface
)

func init() {
	cli.Parse()
	logger = cli.Logger
}

func main() {
	ctx := context.Background()

	if err := chix.RunContext(
		ctx, httpServer(),
	); err != nil {

		log.WithError(err).Fatal("shutting down")
	}
}
