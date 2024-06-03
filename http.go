package main

import (
	"net/http"
	"strings"
	"time"

	"us-soccer-go-test/internal/stadium"
	"us-soccer-go-test/internal/weather"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/lrstanley/chix"
)

func httpServer() *http.Server {
	chix.DefaultAPIPrefix = "/api/"

	r := chi.NewRouter()

	if len(cli.Flags.HTTP.TrustedProxies) > 0 {
		r.Use(chix.UseRealIP(cli.Flags.HTTP.TrustedProxies, chix.OptUseXForwardedFor))
	}

	r.Use(
		chix.UseContextIP,
		middleware.RequestID,
		chix.UseStructuredLogger(logger),
		chix.UseDebug(cli.Debug),
		chix.UseRecoverer,
		middleware.StripSlashes,
		middleware.Compress(5),
		middleware.Maybe(middleware.StripSlashes, func(r *http.Request) bool {
			return !strings.HasPrefix(r.URL.Path, "/debug/")
		}),
		chix.UseNextURL,
		middleware.Timeout(30*time.Second),
	)

	r.Use(
		httprate.LimitByIP(200, 1*time.Minute),
	)

	if cli.Debug {
		r.Mount("/debug", middleware.Profiler())
	}

	/** Routes **/
	r.Route("/api/stadiums", stadium.NewController(logger, db).Route)
	r.Route("/api/weather", weather.NewController(logger, db).Route)

	// Setup our http server.
	return &http.Server{
		Addr:    cli.Flags.HTTP.Addr,
		Handler: r,

		// Some sane defaults.
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
